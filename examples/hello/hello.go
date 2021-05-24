package main

import (
	"fmt"
	"github.com/LilithGames/agent-go/pkg/agent"
	"github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
	"math/rand"
	"time"
)

func HelloHandlers() agent.Handlers {
	handlers := make(map[string]agent.Handler)
	handlers["helloA"] = HelloA
	handlers["helloB"] = HelloB
	handlers["helloC"] = HelloC
	handlers["helloD"] = HelloD
	handlers["helloE"] = HelloE
	return handlers
}

func HelloA(tick *core.Tick) (behavior3go.Status, error) {
	fmt.Println("helloA")
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	return behavior3go.SUCCESS, nil
}

func HelloB(tick *core.Tick) (behavior3go.Status, error) {
	fmt.Println("helloB")
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	return behavior3go.SUCCESS, nil
}

func HelloC(tick *core.Tick) (behavior3go.Status, error) {
	fmt.Println("helloC")
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	return behavior3go.SUCCESS, nil
}

func HelloD(tick *core.Tick) (behavior3go.Status, error) {
	fmt.Println("helloD")
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	return behavior3go.SUCCESS, nil
}

func HelloE(tick *core.Tick) (behavior3go.Status, error) {
	fmt.Println("helloE")
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	return behavior3go.SUCCESS, nil
}
