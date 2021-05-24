package agent

import (
	"encoding/json"
	"fmt"
	"github.com/LilithGames/agent-go/internal/transfer"
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
	Trees    []config.BTTreeCfg
	Metadata map[string]string
}

func (c *Config) MarshalBinary() (data []byte, err error) {
	return json.Marshal(c)
}

func (c *Config) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

func (c *Config) UnmarshalRawConfig(rawTree, rawOtherCfg []byte) error {
	var err error
	if len(rawTree) != 0 {
		var projectCfg config.RawProjectCfg
		err = json.Unmarshal(rawTree, &projectCfg)
		if err != nil {
			return fmt.Errorf("unmarshal tree config err: %w", err)
		}
		c.Trees = projectCfg.Data.Trees
	}
	if len(rawOtherCfg) != 0 {
		err = yaml.Unmarshal(rawOtherCfg, c)
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

type Engine struct {
	plans       []*transfer.Plan
	metadata    map[string]string
	handlers    Handlers
	trees       map[string]*config.BTTreeCfg
	registerMap *behavior3go.RegisterStructMaps
}

func NewEngine() *Engine {
	return &Engine{
		trees:       map[string]*config.BTTreeCfg{},
		handlers:    map[string]Handler{},
		registerMap: behavior3go.NewRegisterStructMaps(),
	}
}

func (e *Engine) RegisterHandler(name string, handler Handler) {
	e.handlers[name] = handler
	e.registerMap.Register(name, new(Action))
}

func (e *Engine) RegisterHandlers(handlers Handlers) {
	for name, handler := range handlers {
		e.RegisterHandler(name, handler)
	}
}

func (e *Engine) RegisterNode(name string, node interface{}) {
	e.registerMap.Register(name, node)
}

func (e *Engine) SetTreesFromConfig(conf []byte) {
	var rawCfg config.RawProjectCfg
	err := json.Unmarshal(conf, &rawCfg)
	if err != nil {
		log.Panic("unmarshal behavior tree: ", zap.Error(err))
	}
	e.setBehaviorTrees(rawCfg.Data.Trees)
}

func (e *Engine) setBehaviorTrees(trees []config.BTTreeCfg) {
	for _, tree := range trees {
		var treeCfg = tree
		loader.CreateBevTreeFromConfig(&treeCfg, e.registerMap)
		e.trees[tree.Title] = &treeCfg
	}
}

func (e *Engine) BuildEngineFromConfig(conf []byte) {
	var c Config
	err := yaml.Unmarshal(conf, &c)
	if err != nil {
		log.Panic("unmarshal plan config", zap.Error(err))
	}
	e.setExecPlans(c.Plans)
	e.metadata = c.Metadata
}

func (e *Engine) setExecPlans(plans []*transfer.Plan) {
	for _, plan := range plans {
		if _, ok := e.trees[plan.TreeName]; !ok {
			log.Panic("not found name in executors")
		}
	}
	e.plans = plans
}
