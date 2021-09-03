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
	behavior := agent.NewBehavior()
	behavior.RegisterHandlers(HelloHandlers())
	behavior.RegisterNode("TestSubscription", NewSubscription)
	behavior.RegisterNode("TestSubscriber", NewSubscriber)
	behavior.RegisterTreeConfig(helloB3)
	engine := behavior.BuildEngineFromConfig(task)
	a := agent.NewAgent(engine)
	a.Start()
}
