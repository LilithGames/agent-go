package agent

import (
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
	"sync"
	"time"
)

type manager struct {
	engine *Engine
	stream *proxyStream
}

func newManager(engine *Engine, stream *proxyStream) *manager {
	return &manager{
		engine: engine,
		stream: stream,
	}
}

func (m *manager) startService() {
	var mail *transfer.Mail
	var err error
	for {
		select {
		case <-m.stream.ctx.Done():
			return
		default:
			mail, err = m.stream.Recv()
		}
		if errors.Is(err, io.EOF) {
			log.Info("receive stop command...")
			return
		}
		if err != nil {
			log.Error("receive error from master", zap.Error(err))
			return
		}
		err = m.reviewMail(mail)
		if err != nil {
			log.Error("internal error in agent", zap.Error(err))
			return
		}
	}
}

func (m *manager) setEngineEnvs(content []byte) error {
	var envs map[string]string
	err := json.Unmarshal(content, &envs)
	if err != nil {
		return err
	}
	if len(envs) != 0 {
		m.engine.setMetadata(envs)
	}
	return nil
}

func (m *manager) reviewMail(mail *transfer.Mail) error {
	switch mail.Action {
	case transfer.ACTION_START_AGENT:
		return m.startAgentEngine(mail.Content)
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
	executors := m.buildExecutors()
	for  {
		for _, exec := range executors {
			m.startExecutor(exec)
		}
	}
}

func (m *manager) startAgentEngine(content []byte) error {
	err := m.setEngineEnvs(content)
	if err != nil {
		return fmt.Errorf("set envs error: %w", err)
	}
	m.startReadyService()
	return nil
}

func (m *manager) buildExecutors() []*executor {
	executors := make([]*executor, len(m.engine.plans))
	for index, plan := range m.engine.plans {
		treeCfg := m.engine.trees[plan.TreeName]
		executor := &executor{
			plan:     plan,
			treeCreator: func() *core.BehaviorTree {
				return loader.CreateBevTreeFromConfig(&treeCfg, m.engine.registerMap)
			},
			metadata: m.engine.metadata,
		}
		executors[index] = executor
	}
	return executors
}

func (m *manager) startReadyService() {
	m.stream.setPlanCount(len(m.engine.plans), false)
	executors := m.buildExecutors()
	for _, exec := range executors{
		m.startExecutor(exec)
	}
}

func (m *manager) startExecutor(executor *executor) {
	system := actor.NewActorSystem()
	props := actor.PropsFromProducer(reporterFactory(executor.plan.TreeName, m.stream))
	actuaryID := system.Root.Spawn(props)
	defer system.Root.Poison(actuaryID)

	robotNum := int(executor.plan.RobotNum)
	wg := &sync.WaitGroup{}
	wg.Add(robotNum)
	defer wg.Wait()

	for i := 0; i < int(executor.plan.RobotNum); i++ {
		job := newJob().
			withMetadata(executor.metadata).
			withCancelCtx(m.stream.ctx).
			withWaitGroup(wg).
			withStatPID(actuaryID).
			withBeTree(executor.treeCreator())
		con := int(executor.plan.Parallel)
		interval := time.Second * time.Duration(executor.plan.Interval)
		if con != 0 && i%con == 0 && i/con > 0 {
			time.Sleep(interval)
		}
		props := actor.PropsFromProducer(newRobot)
		robotID := system.Root.Spawn(props)
		system.Root.Send(robotID, job)
		system.Root.Poison(robotID)
	}
}
