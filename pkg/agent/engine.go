package agent

import (
	"encoding/json"
	"fmt"
	"github.com/LilithGames/agent-go/pkg/transfer"
	"github.com/LilithGames/agent-go/tools/log"
	"github.com/ghodss/yaml"
	"github.com/magicsea/behavior3go"
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
	handlers Handlers
	plan     *transfer.Plan
	tree     *core.BehaviorTree
	metadata map[string]string
}

type Behavior struct {
	handlers    Handlers
	trees       map[string]*config.BTTreeCfg
	registerMap *behavior3go.RegisterStructMaps
}

func NewBehavior() *Behavior {
	return &Behavior{
		handlers: map[string]Handler{},
		trees: map[string]*config.BTTreeCfg{},
		registerMap: behavior3go.NewRegisterStructMaps(),
	}
}

func (b *Behavior) RegisterHandler(name string, handler Handler) {
	b.handlers[name] = handler
	b.registerMap.Register(name, new(Action))
}

func (b *Behavior) RegisterHandlers(handlers Handlers) {
	for name, handler := range handlers {
		b.RegisterHandler(name, handler)
	}
}

func (b *Behavior) RegisterNode(name string, node interface{}) {
	b.registerMap.Register(name, node)
}

func (b *Behavior) RegisterTreeConfig(conf []byte) {
	var rawCfg config.RawProjectCfg
	err := json.Unmarshal(conf, &rawCfg)
	if err != nil {
		log.Panic("unmarshal behavior tree: ", zap.Error(err))
	}
	for _, tree := range rawCfg.Data.Trees {
		var cfg = tree
		loader.CreateBevTreeFromConfig(&cfg, b.registerMap)
		b.trees[cfg.Title] = &cfg
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
			log.Panic("plan name not found")
		}
	}
	engine.plans = cfg.Plans
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

func (e *Engine) setTrees(trees []config.BTTreeCfg) {
	for _, tree := range trees {
		var treeCfg = tree
		loader.CreateBevTreeFromConfig(&treeCfg, e.registerMap)
		e.trees[tree.Title] = &treeCfg
	}
}

func (e *Engine) setPlans(plans []*transfer.Plan) {
	for _, plan := range plans {
		if _, ok := e.trees[plan.TreeName]; !ok {
			log.Panic("not found name in executors")
		}
	}
	e.plans = plans
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
