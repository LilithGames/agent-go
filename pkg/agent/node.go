package agent

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/LilithGames/agent-go/pkg/transfer"
	"github.com/LilithGames/agent-go/tools/log"
	"github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/config"
	"github.com/magicsea/behavior3go/core"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	"go.uber.org/zap"
)

var hc = metric.Must(global.Meter("agent")).NewInt64Counter("agent_stress_handler_outcome")

type Handler func(ticker Ticker) (behavior3go.Status, error)

type Handlers map[string]Handler

type Action struct {
	core.Action
	handler Handler
}

func (n *Action) Initialize(params *config.BTNodeCfg) {
	n.Action.Initialize(params)
}

func (n *Action) OnTick(ticker core.Ticker) behavior3go.Status {
	tick := ticker.(*Tick)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)+10))
	name := n.GetName()
	task := tick.GetTree().GetTitile()
	defer n.recoverPanic(tick, name)
	var outcome transfer.Outcome
	outcome.Class = transfer.CLASS_HANDLER
	start := time.Now()
	status, err := n.handler(tick)
	outcome.Status = transfer.STATUS(status)
	if err != nil {
		log.Error("handler error", zap.String("handler", name), zap.String("task", task), zap.Error(err))
		outcome.Err = err.Error()
		msg := &ErrMsg{Name: name, Intro: "CI-Error", Detail: outcome.Err}
		err := tick.SendAlertMsg(msg)
		if err != nil {
			log.Error("send alert message error: ", zap.Error(err))
		}
	}
	outcome.Name = name
	if start.UnixNano() < tick.recvTime {
		outcome.Consume = tick.sendTime - tick.recvTime
	} else {
		outcome.Consume = time.Since(start).Nanoseconds()
	}
	tick.actorCtx().Send(tick.stat(), &outcome)
	n.addMetric(task, name, status)
	return n.currentStatus(tick.context(), status)
}

func (n *Action) currentStatus(ctx context.Context, status behavior3go.Status) behavior3go.Status {
	select {
	case <-ctx.Done():
		return behavior3go.FAILURE
	default:
		return status
	}
}

func (n *Action) recoverPanic(tick *Tick, name string) {
	if r := recover(); r != nil {
		var outcome transfer.Outcome
		outcome.Name = name
		outcome.Status = transfer.STATUS_FAILURE
		outcome.Class = transfer.CLASS_HANDLER
		outcome.Err = fmt.Sprintf("receive panic: %v", r)
		tick.actorCtx().Send(tick.stat(), &outcome)
		task := tick.GetTree().GetTitile()
		n.addMetric(task, name, behavior3go.ERROR)
		detail := fmt.Sprintf("panic: %v", r)
		log.Error("handler panic", zap.String("handler", name), zap.String("task", task), zap.Error(fmt.Errorf("%s", detail)))
		msg := &ErrMsg{Name: name, Intro: "CI-Panic", Detail: detail}
		err := tick.SendAlertMsg(msg)
		if err != nil {
			log.Error("send alert message error: ", zap.Error(err))
		}
	}
}

func (n *Action) addMetric(task, node string, status behavior3go.Status) {
	hl := attribute.String("handler", node)
	tl := attribute.String("task", task)
	ctx := context.Background()
	switch status {
	case behavior3go.SUCCESS:
		hc.Add(ctx, 1, tl, hl, attribute.String("status", "success"))
	case behavior3go.ERROR, behavior3go.FAILURE:
		hc.Add(ctx, 1, tl, hl, attribute.String("status", "failure"))
	case behavior3go.RUNNING:
		hc.Add(ctx, 1, tl, hl, attribute.String("status", "running"))
	}
}
