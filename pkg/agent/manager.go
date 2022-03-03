package agent

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
	"sync"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/LilithGames/agent-go/pkg/transfer"
	"github.com/LilithGames/agent-go/tools/log"
	"github.com/magicsea/behavior3go/core"
	"github.com/magicsea/behavior3go/loader"
	"go.uber.org/zap"
)

type manager struct {
	*Agent
}

func newManagerFromAgent(agent *Agent) *manager {
	return &manager{agent}
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

func (m *manager) reviewMail(mail *transfer.Mail) error {
	switch mail.Action {
	case transfer.ACTION_START_AGENT:
		return m.startAgentEngine(mail.Content, false)
	case transfer.ACTION_START_CIRCLE:
		return m.startAgentEngine(mail.Content, true)
	}
	return nil
}

func (m *manager) startAgentEngine(content []byte, circle bool) error {
	var envs map[string]string
	err := json.Unmarshal(content, &envs)
	if err != nil {
		return fmt.Errorf("unmarshal envs error: %w", err)
	}
	m.setEnv(envs)
	m.stream.setPlanCount(len(m.engine.plans), circle)
	m.startAgentExecutors(circle)
	return nil
}

func (m *manager) startLocalService() {
	var circle bool
	if m.cfg.GetString("mode") == "circle" {
		circle = true
	}
	m.stream.setPlanCount(len(m.engine.plans), circle)
	m.startAgentExecutors(circle)
}

func (m *manager) setEnv(envs map[string]string) {
	for k, v := range envs {
		m.cfg.Set(k, v)
		if k == "logLevel" && v != "" {
			log.ResetLogLevel(v)
		}
	}
}

func (m *manager) startAgentExecutors(circle bool) {
	executors, market := m.buildExecutors()
	for {
		for _, exec := range executors {
			select {
			case <-m.ctx.Done():
				return
			default:
				m.startExecutor(exec, market)
			}
		}
		if !circle {
			return
		}
	}
}

func (m *manager) buildExecutors() ([]*executor, *Market) {
	executors := make([]*executor, len(m.engine.plans))
	var mc int32
	for index, plan := range m.engine.plans {
		treeCfg := m.engine.behavior.trees[plan.TreeName]
		executor := &executor{
			plan: plan,
			treeCreator: func() *core.BehaviorTree {
				return loader.CreateBevTreeFromConfig(&treeCfg, m.engine.behavior.registerMap)
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
			withMarket(market).
			withAlert(m.alert)
		if i%parallel == 0 && i/parallel > 0 {
			<-ticker.C
		}
		go newRobot(job).execute(system.Root)
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
	pe := m.cfg.GetString("parallel")
	if pe != "" {
		rs, _ = strconv.Atoi(pe)
		return rs
	}
	for _, plan := range m.engine.plans {
		if int(plan.Parallel) > rs {
			rs = int(plan.Parallel)
		}
	}
	return math.MaxInt8
}
