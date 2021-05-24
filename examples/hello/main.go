package main

import (
	_ "embed"
	"github.com/LilithGames/agent-go/pkg/agent"
)

//go:embed task.yaml
var task []byte

//go:embed hello.b3
var helloB3 []byte

func main() {
	engine := agent.NewEngine()
	engine.RegisterHandlers(HelloHandlers())
	engine.SetTreesFromConfig(helloB3)
	engine.BuildEngineFromConfig(task)
	a := agent.NewAgent(engine)
	a.Start()
}
