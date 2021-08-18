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

type Config struct {
	Plans    []*transfer.Plan
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
	plan     *transfer.Plan
	tree     *core.BehaviorTree
	metadata map[string]string
}

type Behavior struct {
	trees       map[string]config.BTTreeCfg
	registerMap *core.RegisterStructMaps
}

func NewBehavior() *Behavior {
	return &Behavior{
		trees: map[string]config.BTTreeCfg{},
		registerMap: core.NewRegisterStructMaps(),
	}
}

func (b *Behavior) RegisterHandler(name string, handler Handler) {
	action := &Action{handler: handler}
	b.registerMap.Register(name, action)
}

func (b *Behavior) RegisterHandlers(handlers Handlers) {
	for name, handler := range handlers {
		b.RegisterHandler(name, handler)
	}
}

func (b *Behavior) RegisterNode(name string, node core.IBaseNode) {
	b.registerMap.Register(name, node)
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
	engine.metadata = cfg.Environments
	for _, plan := range cfg.Plans {
		if _, ok := b.trees[plan.TreeName]; !ok {
			log.Panic("plan name not found: " + plan.TreeName)
		}
	}
	engine.plans = cfg.Plans
	core.SetSubTreeLoadFunc(func(name string) *core.BehaviorTree {
		if t, ok := b.trees[name]; ok {
			return loader.CreateBevTreeFromConfig(&t, b.registerMap)
		}
		log.Panic("create sub tree not found.")
		return nil
	})
	return engine
}

func (b *Behavior) BuildTestEngine(envs map[string]string, plan *transfer.Plan) *Engine {
	engine := &Engine{Behavior: b}
	engine.metadata = envs
	engine.plans = append(engine.plans, plan)
	return engine
}

type Engine struct {
	*Behavior
	plans       []*transfer.Plan
	metadata    map[string]string
}

func (e *Engine) setMetadata(envs map[string]string) {
	if len(e.metadata) == 0 {
		e.metadata = envs
		return
	}
	for name, value := range envs {
		e.metadata[name] = value
	}
}
