package agent

import (
	"encoding/json"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/LilithGames/agent-go/pkg/transfer"
	"github.com/LilithGames/agent-go/tools/log"
	"github.com/hasura/go-graphql-client"
	"github.com/magicsea/behavior3go/actions"
	"github.com/magicsea/behavior3go/composites"
	"github.com/magicsea/behavior3go/core"
	"go.uber.org/zap"
	"os"
	"time"
)

type ClientOption func(client *graphql.SubscriptionClient)

type MessageHandler func(message *json.RawMessage, err error) error

func WithLog(log func(args ...interface{})) ClientOption {
	return func(client *graphql.SubscriptionClient) {
		client.WithLog(log)
	}
}

func WithOnErr(onError func(sc *graphql.SubscriptionClient, err error) error) ClientOption {
	return func(client *graphql.SubscriptionClient) {
		client.OnError(onError)
	}
}

func WithOnConnected(fn func()) ClientOption {
	return func(client *graphql.SubscriptionClient) {
		client.OnConnected(fn)
	}
}

func WithOnDisconnected(fn func()) ClientOption {
	return func(client *graphql.SubscriptionClient) {
		client.OnDisconnected(fn)
	}
}

func WithoutLogTypes(types ...graphql.OperationMessageType) ClientOption {
	return func(client *graphql.SubscriptionClient) {
		client.WithoutLogTypes(types...)
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(client *graphql.SubscriptionClient) {
		client.WithTimeout(timeout)
	}
}

type GqlSubscription struct {
	composites.Subscription
	token string
}

func (g *GqlSubscription) OnOpen(tick *core.Tick) {
	rawToken := tick.Blackboard.GetMem("token")
	if rawToken != nil {
		g.token = rawToken.(string)
	}
}

func NewGqlSubscription(options ...ClientOption) *GqlSubscription {
	subscription := &GqlSubscription{}
	subscription.ClientCreator = func() composites.SubClient {
		backend := os.Getenv("backend")
		if backend == "" {
			log.Panic("graphql backend not found")
		}
		client := graphql.NewSubscriptionClient(backend)
		if subscription.token != "" {
			client.WithConnectionParams(map[string]interface{}{
				"Authorization": subscription.token,
			})
		}
		for _, option :=  range options {
			option(client)
		}
		return client
	}
	return subscription
}

type GqlSubscriber struct {
	actions.Subscriber
	actorID *actor.PID
	actorCtx actor.Context
}

type Message struct {
	Data *json.RawMessage `json:"data"`
	Extensions struct{
		Debug struct{
			SendTime int64 `json:"send_time"`
		} `json:"debug"`
	} `json:"extensions"`
}

func (g *GqlSubscriber) OnOpen(tick *core.Tick) {
	g.actorCtx = tick.Blackboard.GetMem("actorCtx").(actor.Context)
	job := tick.GetTarget().(*job)
	g.actorID = job.statPID
}

func (g *GqlSubscriber) GqlSubscriberWrapHandler(name string, handler MessageHandler) MessageHandler {
	return func(message *json.RawMessage, err error) error {
		var rawMsg Message
		errj := json.Unmarshal(*message, &rawMsg)
		if errj != nil {
			log.Panic("unmarshal message error", zap.Error(errj))
		}
		var outcome transfer.Outcome
		start := time.Unix(rawMsg.Extensions.Debug.SendTime, 0).Unix()
		outcome.Consume = time.Now().Unix() - start
		if handler != nil {
			err = handler(rawMsg.Data, err)
		}
		outcome.Class = transfer.CLASS_EVENT
		outcome.Name = name
		outcome.Status = transfer.STATUS_SUCCESS
		if err != nil {
			outcome.Err = err.Error()
			outcome.Status = transfer.STATUS_ERROR
		}
		g.actorCtx.Send(g.actorID, &outcome)
		return err
	}
}

func NewGqlSubscriber(name string, query interface{}, variables map[string]interface{}, handler MessageHandler) *GqlSubscriber {
	subscriber := &GqlSubscriber{}
	subscriber.SubTopic = func(client interface{}) error {
		subClient := client.(*graphql.SubscriptionClient)
		wrapHandler := subscriber.GqlSubscriberWrapHandler(name, handler)
		_, err := subClient.NamedSubscribe(name, query, variables, wrapHandler)
		return err
	}
	return subscriber
}


