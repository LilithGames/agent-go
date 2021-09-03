package main

import (
	"encoding/json"
	"fmt"
	"github.com/LilithGames/agent-go/pkg/agent"
	"github.com/hasura/go-graphql-client"
	"github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
	"log"
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
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	return behavior3go.SUCCESS, nil
}

func HelloB(tick *core.Tick) (behavior3go.Status, error) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	return behavior3go.SUCCESS, nil
}

func HelloC(tick *core.Tick) (behavior3go.Status, error) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	return behavior3go.SUCCESS, nil
}

func HelloD(tick *core.Tick) (behavior3go.Status, error) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	return behavior3go.SUCCESS, nil
}

func HelloE(tick *core.Tick) (behavior3go.Status, error) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	return behavior3go.SUCCESS, nil
}


func NewSubscription() core.IBaseNode {
	url := "ws://host.docker.internal:8085/query"
	// url := "ws://127.0.0.1:8085/query"
	subscription := agent.NewGqlSubscription(url, agent.WithLog(log.Println))
	fmt.Printf("subscription node %p \n", subscription)
	return subscription
}

func NewSubscriber() core.IBaseNode {
	var query struct{
		Message struct {
			id   graphql.String
			text graphql.String
		} `graphql:"messageAdded(roomName:$roomName)"`
	}
	variables := map[string]interface{}{
		"roomName": graphql.String("#gophers"),
	}
	handler := func(message *json.RawMessage, err error) error {
		if err != nil {
			log.Println(err)
		}
		return err
	}
	return agent.NewGqlSubscriber("test", query, variables, handler)
}