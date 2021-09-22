package agent

import (
	"encoding/json"
	"fmt"
	"github.com/LilithGames/agent-go/pkg/transfer"
	"github.com/LilithGames/agent-go/tools/log"
	"github.com/ghodss/yaml"
	"github.com/magicsea/behavior3go/config"
	"github.com/magicsea/behavior3go/core"
	"github.com/magicsea/behavior3go/loader"
	"go.uber.org/zap"
)

type Nodes map[string]core.NodeCreator

type Config struct {
	Plans        []*transfer.Plan
	Environments map[string]string
}

func (c *Config) MarshalBinary() (data []byte, err error) {
	return json.Marshal(c)
}

func (c *Config) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

func (c *Config) UnmarshalRawConfig(rawCfg []byte) error {
	var err error
	if len(rawCfg) != 0 {
		err = yaml.Unmarshal(rawCfg, c)
		if err != nil {
			return fmt.Errorf("unmarshal task config: %w", err)
		}
	}
	return nil
}

type executor struct {
	plan        *transfer.Plan
	treeCreator func() *core.BehaviorTree
}

type Behavior struct {
	trees       map[string]config.BTTreeCfg
	registerMap *core.RegisterStructMaps
}

func NewBehavior() *Behavior {
	return &Behavior{
		trees:       map[string]config.BTTreeCfg{},
		registerMap: core.NewRegisterStructMaps(),
	}
}

func (b *Behavior) RegisterHandler(name string, handler Handler) {
	creator := func() core.IBaseNode {
		return &Action{handler: handler}
	}
	b.registerMap.Register(name, creator)
}

func (b *Behavior) RegisterHandlers(handlers Handlers) {
	for name, handler := range handlers {
		b.RegisterHandler(name, handler)
	}
}

func (b *Behavior) RegisterNode(name string, creator core.NodeCreator) {
	b.registerMap.Register(name, creator)
}

func (b *Behavior) RegisterNodes(nodes Nodes) {
	for name, node := range nodes {
		b.RegisterNode(name, node)
	}
}

func (b *Behavior) RegisterTreeConfig(conf []byte) {
	var rawCfg config.RawProjectCfg
	err := json.Unmarshal(conf, &rawCfg)
	if err != nil {
		log.Panic("unmarshal behavior tree: ", zap.Error(err))
	}
	err = loader.CheckTreeComplete(rawCfg.Data.Trees, b.registerMap)
	if err != nil {
		log.Panic("behavior tree check failed: ", zap.Error(err))
	}
	for _, tree := range rawCfg.Data.Trees {
		b.trees[tree.Title] = tree
	}
}

func (b *Behavior) BuildEngineFromConfig(conf []byte) *Engine {
	var cfg Config
	err := yaml.Unmarshal(conf, &cfg)
	if err != nil {
		log.Panic("unmarshal plan config", zap.Error(err))
	}
	engine := &Engine{Behavior: b}
	engine.envs = cfg.Environments
	for _, plan := range cfg.Plans {
		if _, ok := b.trees[plan.TreeName]; !ok {
			log.Panic("plan name not found: " + plan.TreeName)
		}
	}
	engine.plans = cfg.Plans
	b.registerSubTreeLoadFunc()
	return engine
}

func (b *Behavior) BuildTestEngine(envs map[string]string, plan *transfer.Plan) *Engine {
	engine := &Engine{Behavior: b}
	engine.envs = envs
	engine.plans = append(engine.plans, plan)
	b.registerSubTreeLoadFunc()
	return engine
}

func (b *Behavior) registerSubTreeLoadFunc() {
	core.SetSubTreeLoadFunc(func(name string) *core.BehaviorTree {
		if t, ok := b.trees[name]; ok {
			return loader.CreateBevTreeFromConfig(&t, b.registerMap)
		}
		log.Panic("create sub tree not found.")
		return nil
	})
}

type Engine struct {
	*Behavior
	plans []*transfer.Plan
	envs  map[string]string
}
