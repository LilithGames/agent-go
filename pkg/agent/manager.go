package agent

import (
	"encoding/json"
	"errors"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/LilithGames/agent-go/pkg/transfer"
	"github.com/LilithGames/agent-go/tools/log"
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

func (m *manager) reviewMail(mail *transfer.Mail) error {
	switch mail.Action {
	case transfer.ACTION_START_AGENT:
		return m.startAgentEngine(mail.Content)
	default:
		return nil
	}
}

func (m *manager) startAgentEngine(content []byte) error {
	var conf Config
	err := json.Unmarshal(content, &conf)
	if err != nil {
		return err
	}
	if len(conf.Trees) != 0 {
		m.engine.setBehaviorTrees(conf.Trees)
	}
	if conf.Plans != nil {
		m.engine.setExecPlans(conf.Plans)
	}
	if len(conf.Metadata) != 0 {
		m.engine.metadata = conf.Metadata
	}
	m.startReadyService()
	return nil
}

func (m *manager) startReadyService() {
	m.stream.setPlanCount(len(m.engine.plans))
	for _, plan := range m.engine.plans {
		treeCfg := m.engine.trees[plan.TreeName]
		tree := loader.CreateBevTreeFromConfig(treeCfg, m.engine.registerMap)
		executor := &executor{
			plan:     plan,
			handlers: m.engine.handlers,
			tree:     tree,
			metadata: m.engine.metadata,
		}
		m.startExecutor(executor)
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

	job := newJob().
		withMetadata(executor.metadata).
		withHandlers(executor.handlers).
		withCancelCtx(m.stream.ctx).
		withBeTree(executor.tree).
		withWaitGroup(wg).
		withStatPID(actuaryID)

	for i := 0; i < int(executor.plan.RobotNum); i++ {
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
