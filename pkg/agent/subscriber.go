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
	"strconv"
	"time"
)

type ClientOption func(client *graphql.SubscriptionClient)

type MessageCallback func(tick Ticker, message *json.RawMessage, err error) error

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

func (g *GqlSubscription) OnOpen(tick core.Ticker) {
	rawToken := tick.Blackboard().GetMem("token")
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
		for _, option := range options {
			option(client)
		}
		return client
	}
	return subscription
}

type GqlSubscriber struct {
	actions.Subscriber
	actorID   *actor.PID
	actorCtx  actor.Context
	variables map[string]interface{}
	callback  MessageCallback
	reply     interface{}
}

type Message struct {
	Data       *json.RawMessage `json:"data"`
	Extensions struct {
		Debug struct {
			SendTime string `json:"send_time"`
			RecvTime string `json:"recv_time"`
		} `json:"debug"`
	} `json:"extensions"`
}

func (g *GqlSubscriber) OnOpen(ticker core.Ticker) {
	tick := ticker.(*Tick)
	g.actorCtx = tick.Blackboard().GetMem("actorCtx").(actor.Context)
	g.actorID = tick.stat()
}

func (g *GqlSubscriber) GqlSubscriberWrapHandler(name string, tick Ticker) MessageHandler {
	return func(message *json.RawMessage, err error) error {
		if err != nil {
			log.Error("receive event data error", zap.Error(err))
			return err
		}
		var rawMsg Message
		err = json.Unmarshal(*message, &rawMsg)
		if err != nil {
			log.Error("unmarshal message error", zap.Error(err))
			return err
		}
		if g.reply != nil {
			err = json.Unmarshal(*rawMsg.Data, g.reply)
			if err != nil {
				log.Error("unmarshal message error", zap.Error(err))
				return err
			}
		}
		var outcome transfer.Outcome
		start, err := strconv.ParseInt(rawMsg.Extensions.Debug.SendTime, 10, 64)
		if err != nil {
			log.Error("parse send time error", zap.Error(err))
			return err
		}
		outcome.Consume = time.Now().UnixNano() - start
		if g.callback != nil {
			err = g.callback(tick, rawMsg.Data, err)
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

func (g *GqlSubscriber) WithVariables(variables map[string]interface{}) *GqlSubscriber {
	g.variables = variables
	return g
}

func (g *GqlSubscriber) WithCallback(callback MessageCallback) *GqlSubscriber {
	g.callback = callback
	return g
}

func (g *GqlSubscriber) WithReply(reply interface{}) *GqlSubscriber {
	g.reply = reply
	return g
}

func NewGqlSubscriber(name string, query interface{}) *GqlSubscriber {
	subscriber := &GqlSubscriber{}
	subscriber.SubTopic = func(ticker core.Ticker, client interface{}) error {
		tick := ticker.(*Tick)
		subClient := client.(*graphql.SubscriptionClient)
		wrapHandler := subscriber.GqlSubscriberWrapHandler(name, tick)
		_, err := subClient.NamedSubscribe(name, query, subscriber.variables, wrapHandler)
		return err
	}
	return subscriber
}
