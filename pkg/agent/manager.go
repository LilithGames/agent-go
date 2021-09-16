package agent

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/LilithGames/agent-go/pkg/transfer"
	"github.com/LilithGames/agent-go/tools/log"
	"github.com/magicsea/behavior3go/core"
	"github.com/magicsea/behavior3go/loader"
	"go.uber.org/zap"
	"io"
	"math"
	"os"
	"strconv"
	"sync"
	"time"
)

type manager struct {
	engine *Engine
	stream *proxyStream
	ctx    context.Context
	cancel context.CancelFunc
}

func newManager(engine *Engine, stream *proxyStream) *manager {
	ctx, cancel := context.WithCancel(context.Background())
	stream.withContext(ctx)
	return &manager{
		engine: engine,
		stream: stream,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (m *manager) startClusterService() {
	mailbox := make(chan *transfer.Mail)
	go m.receiveMail(mailbox)
	for {
		mail := <-mailbox
		if mail == nil {
			return
		}
		err := m.reviewMail(mail)
		if err != nil {
			log.Error("internal error in agent", zap.Error(err))
			return
		}
	}
}

func (m *manager) receiveMail(mailbox chan *transfer.Mail) {
	for {
		select {
		case <-m.ctx.Done():
			close(mailbox)
			return
		default:
			mail, err := m.stream.Recv()
			if errors.Is(err, io.EOF) {
				m.cancel()
				break
			}
			mailbox <- mail
		}
	}
}

func (m *manager) setEngineEnvs(content []byte) error {
	var envs map[string]string
	err := json.Unmarshal(content, &envs)
	if err != nil {
		return err
	}
	for k, v := range envs {
		err = os.Setenv(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *manager) reviewMail(mail *transfer.Mail) error {
	switch mail.Action {
	case transfer.ACTION_START_AGENT:
		return m.startAgentOnceEngine(mail.Content)
	case transfer.ACTION_START_CIRCLE:
		return m.startAgentCircleEngine(mail.Content)
	default:
		return nil
	}
}

func (m *manager) startAgentCircleEngine(content []byte) error {
	err := m.setEngineEnvs(content)
	if err != nil {
		return err
	}
	m.stream.setPlanCount(len(m.engine.plans), true)
	executors, market := m.buildExecutors()
	for {
		for _, exec := range executors {
			select {
			case <-m.ctx.Done():
				return nil
			default:
				m.startExecutor(exec, market)
			}
		}
	}
}

func (m *manager) startAgentOnceEngine(content []byte) error {
	err := m.setEngineEnvs(content)
	if err != nil {
		return fmt.Errorf("set envs error: %w", err)
	}
	m.stream.setPlanCount(len(m.engine.plans), false)
	m.startAgentOnceExecutors()
	return nil
}

func (m *manager) startLocalService() {
	m.stream.setPlanCount(len(m.engine.plans), false)
	for k, v := range m.engine.envs {
		_ = os.Setenv(k, v)
	}
	m.startAgentOnceExecutors()
}

func (m *manager) startAgentOnceExecutors() {
	executors, market := m.buildExecutors()
	for _, exec := range executors {
		select {
		case <-m.ctx.Done():
			return
		default:
			m.startExecutor(exec, market)
		}
	}
}

func (m *manager) buildExecutors() ([]*executor, *Market) {
	executors := make([]*executor, len(m.engine.plans))
	var mc int32
	for index, plan := range m.engine.plans {
		treeCfg := m.engine.trees[plan.TreeName]
		executor := &executor{
			plan: plan,
			treeCreator: func() *core.BehaviorTree {
				return loader.CreateBevTreeFromConfig(&treeCfg, m.engine.registerMap)
			},
		}
		executors[index] = executor
		if plan.RobotNum > mc {
			mc = plan.RobotNum
		}
	}
	market := newMarket(int(mc))
	return executors, market
}

func (m *manager) startExecutor(executor *executor, market *Market) {
	system := actor.NewActorSystem()
	props := actor.PropsFromProducer(reporterFactory(executor.plan.TreeName, m.stream))
	actuaryID := system.Root.Spawn(props)

	robotNum := int(executor.plan.RobotNum)
	wg := &sync.WaitGroup{}
	wg.Add(robotNum)

	parallel := m.getParallel()
	ticker := time.NewTicker(time.Second)
	for i := 0; i < int(executor.plan.RobotNum); i++ {
		job := newJob().
			withCancelCtx(m.ctx).
			withWaitGroup(wg).
			withStatPID(actuaryID).
			withBeTree(executor.treeCreator()).
			withMarket(market)
		if i%parallel == 0 && i/parallel > 0 {
			<-ticker.C
		}
		props := actor.PropsFromProducer(newRobot)
		robotID := system.Root.Spawn(props)
		system.Root.Send(robotID, job)
		system.Root.Poison(robotID)
	}
	wg.Wait()
	err := system.Root.PoisonFuture(actuaryID).Wait()
	if err != nil {
		log.Error("stop actuary error", zap.Error(err))
	}
	market.reset()
}

func (m *manager) getParallel() int {
	var rs int
	pe := os.Getenv("parallel")
	if pe != "" {
		rs, _ = strconv.Atoi(pe)
		return rs
	}
	for _, plan := range m.engine.plans {
		if int(plan.Parallel) > rs {
			rs = int(plan.Parallel)
		}
	}
	return math.MaxInt16
}
