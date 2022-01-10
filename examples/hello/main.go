package main

import (
	_ "embed"
	"github.com/LilithGames/agent-go/pkg/agent"
	"encoding/json"
	"fmt"
	"github.com/hasura/go-graphql-client"
	"github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
	"github.com/rs/xid"
	"log"
	"math/rand"
	"time"
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

func HelloHandlers() agent.Handlers {
	handlers := make(map[string]agent.Handler)
	handlers["helloA"] = HelloA
	handlers["helloB"] = HelloB
	handlers["helloC"] = HelloC
	handlers["helloD"] = HelloD
	handlers["helloE"] = HelloE
	handlers["newUser"] = newUser
	handlers["addFriend"] = addFriend
	handlers["buildTeam"] = buildTeam
	return handlers
}

func HelloA(tick agent.Ticker) (behavior3go.Status, error) {
	var p *Player
	one := tick.Marget().InviteOne()
	if one == nil {
		p = NewPlayer(xid.New().String())
		tick.Marget().UseOne(p) 
		fmt.Println("new player id", p.ID())
	} else {
		p = one.(*Player)
	}
	fmt.Println("current player id: ", p.ID())
	return behavior3go.SUCCESS, nil
}

func HelloB(tick agent.Ticker) (behavior3go.Status, error) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	return behavior3go.SUCCESS, nil
}

func HelloC(tick agent.Ticker) (behavior3go.Status, error) {
	return behavior3go.SUCCESS, nil
}

func HelloD(tick agent.Ticker) (behavior3go.Status, error) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	return behavior3go.SUCCESS, nil
}

func HelloE(tick agent.Ticker) (behavior3go.Status, error) {
	return behavior3go.SUCCESS, nil
}

func newUser(tick agent.Ticker) (behavior3go.Status, error) {
	player := NewPlayer(xid.New().String())
	tick.Marget().PushOne(player)
	tick.Blackboard().SetMem("userId", player.id)
	return behavior3go.SUCCESS, nil
}

func addFriend(tick agent.Ticker) (behavior3go.Status, error) {
	userID := tick.Blackboard().GetMem("userId").(string)
	one := tick.Marget().RequireOne(func(one agent.One) bool {
		player := one.(*Player)
		return player.ID() != userID
	})
	if one != nil {
		player := one.(*Player)
		fmt.Println("userID: " + userID + " has friend: " + player.id)
		return behavior3go.SUCCESS, nil
	}
	fmt.Println("not found friend")
	return behavior3go.SUCCESS, nil
}

func buildTeam(tick agent.Ticker) (behavior3go.Status, error) {
	index := tick.Marget().Index()
	userID := tick.Blackboard().GetMem("userId").(string)
	if index%2 == 1 {
		tick.Marget().JoinOne(NewPlayer(userID))
	} else {
		one := tick.Marget().InviteOne()
		if one == nil {
			fmt.Println("no found player to build team")
		} else {
			player := one.(*Player)
			fmt.Println("userID: " + userID + " build team with: " + player.id)
		}
	}
	return behavior3go.SUCCESS, nil
}

func NewSubscription() core.IBaseNode {
	subscription := agent.NewGqlSubscription(agent.WithLog(log.Println))
	return subscription
}

func NewSubscriber() core.IBaseNode {
	var query struct {
		Message struct {
			id   graphql.String
			text graphql.String
		} `graphql:"messageAdded(roomName:$roomName)"`
	}
	variables := map[string]interface{}{
		"roomName": graphql.String("#gophers"),
	}
	handler := func(tick agent.Ticker, message *json.RawMessage, err error) error {
		if err != nil {
			log.Println(err)
		}
		return err
	}
	return agent.NewGqlSubscriber("test", query).WithVariables(variables).WithCallback(handler)
}

type Player struct {
	agent.One
	id string
}

func NewPlayer(id string) *Player {
	return &Player{id: id}
}

func (p *Player) ID() string {
	return p.id
}