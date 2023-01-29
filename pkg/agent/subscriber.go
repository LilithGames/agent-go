package agent

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/LilithGames/agent-go/pkg/transfer"
	"github.com/LilithGames/agent-go/tools/log"
	"github.com/gorilla/websocket"
	"github.com/hasura/go-graphql-client"
	"github.com/magicsea/behavior3go/actions"
	"github.com/magicsea/behavior3go/composites"
	"github.com/magicsea/behavior3go/core"
	"go.uber.org/zap"
)

type NewClientOption func(client *graphql.SubscriptionClient, tick core.Ticker)

type MessageCallback func(tick Ticker, message *json.RawMessage, err error) error

type MessageHandler func(message *json.RawMessage, err error) error

func WithLog(log func(args ...interface{})) NewClientOption {
	return func(client *graphql.SubscriptionClient, _ core.Ticker) {
		client.WithLog(log)
	}
}

func WithOnErr(onError func(sc *graphql.SubscriptionClient, err error) error) NewClientOption {
	return func(client *graphql.SubscriptionClient, _ core.Ticker) {
		client.OnError(onError)
	}
}

func WithOnConnected(fn func()) NewClientOption {
	return func(client *graphql.SubscriptionClient, _ core.Ticker) {
		client.OnConnected(fn)
	}
}

func WithOnDisconnected(fn func()) NewClientOption {
	return func(client *graphql.SubscriptionClient, _ core.Ticker) {
		client.OnDisconnected(fn)
	}
}

func WithoutLogTypes(types ...graphql.OperationMessageType) NewClientOption {
	return func(client *graphql.SubscriptionClient, _ core.Ticker) {
		client.WithoutLogTypes(types...)
	}
}

func WithTimeout(timeout time.Duration) NewClientOption {
	return func(client *graphql.SubscriptionClient, _ core.Ticker) {
		client.WithTimeout(timeout)
	}
}

func WithNothing() NewClientOption {
	return func(_ *graphql.SubscriptionClient, _ core.Ticker) {}
}

type InitFunction func(tick core.Ticker) NewClientOption

func WithInitFunc(init InitFunction) NewClientOption {
	return func(client *graphql.SubscriptionClient, tick core.Ticker) {
		option := init(tick)
		option(client, tick)
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

func NewGraphqlSubscription(backend string, options ...NewClientOption) *GqlSubscription {
	subscription := &GqlSubscription{}
	subscription.ClientCreator = func(tick core.Ticker) composites.SubClient {
		client := graphql.NewSubscriptionClient(backend)
		client.WithWebSocket(newWebsocketConn)
		if subscription.token != "" {
			client.WithConnectionParams(map[string]interface{}{
				"Authorization": subscription.token,
			})
		}
		for _, option := range options {
			option(client, tick)
		}
		return client
	}
	return subscription
}

func newWebsocketConn(sc *graphql.SubscriptionClient) (graphql.WebsocketConn, error) {
	rawURL := sc.GetURL()
	if strings.Contains(rawURL, "https") {
		rawURL = strings.ReplaceAll(rawURL, "https", "wss")
	}
	if strings.Contains(rawURL, "http") {
		rawURL = strings.ReplaceAll(rawURL, "http", "ws")
	}
	conn, _, err := websocket.DefaultDialer.Dial(rawURL, nil)
	return conn, err
}

type GqlSubscriber struct {
	actions.Subscriber
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

func (g *GqlSubscriber) GqlSubscriberWrapHandler(name string, tick Ticker) MessageHandler {
	return func(message *json.RawMessage, err error) error {
		if err != nil {
			log.Error("receive event data error", zap.Error(err))
			return err
		}
		if message == nil {
			return fmt.Errorf("receive raw message nil")
		}
		var rawMsg Message
		err = json.Unmarshal(*message, &rawMsg)
		if err != nil {
			log.Error("unmarshal message error", zap.Error(err))
			return err
		}
		if rawMsg.Data == nil {
			return fmt.Errorf("receive raw message.data nil: %s", string(*message))
		}
		if g.reply != nil {
			err = json.Unmarshal(*rawMsg.Data, g.reply)
			if err != nil {
				log.Error("unmarshal message error", zap.Error(err))
				return err
			}
		}
		defer g.recoverPanic(name, tick)
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
		tick.actorCtx().Send(tick.stat(), &outcome)
		return err
	}
}

func (g *GqlSubscriber) recoverPanic(name string, tick Ticker) {
	if r := recover(); r != nil {
		var outcome transfer.Outcome
		outcome.Name = name
		outcome.Class = transfer.CLASS_EVENT
		outcome.Status = transfer.STATUS_ERROR
		outcome.Err = fmt.Sprintf("receive panic: %v", r)
		tick.actorCtx().Send(tick.stat(), &outcome)
		log.Error("event panic", zap.String("event", name), zap.String("task", tick.GetTree().GetTitile()), zap.Error(fmt.Errorf("panic: %v", r)))
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
		_, err := subClient.Subscribe(query, subscriber.variables, wrapHandler, graphql.OperationName(name))
		return err
	}
	return subscriber
}

func NewGqlRawSubscriber(name, query string) *GqlSubscriber {
	subscriber := &GqlSubscriber{}
	subscriber.SubTopic = func(ticker core.Ticker, client interface{}) error {
		tick := ticker.(*Tick)
		subClient := client.(*graphql.SubscriptionClient)
		wrapHandler := subscriber.GqlSubscriberWrapHandler(name, tick)
		_, err := subClient.SubscribeRaw(query, subscriber.variables, wrapHandler)
		return err
	}
	return subscriber
}
