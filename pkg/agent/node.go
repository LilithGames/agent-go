package agent

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/LilithGames/agent-go/pkg/transfer"
	"github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/config"
	"github.com/magicsea/behavior3go/core"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
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
	defer n.recoverPanic(name, tick)
	var outcome transfer.Outcome
	outcome.Class = transfer.CLASS_HANDLER
	start := time.Now()
	status, err := n.handler(tick)
	outcome.Status = transfer.STATUS(status)
	if err != nil {
		outcome.Err = err.Error()
	}

	outcome.Name = name
	if start.UnixNano() < tick.recvTime {
		outcome.Consume = tick.sendTime - tick.recvTime
	} else {
		outcome.Consume = time.Since(start).Nanoseconds()
	}
	tick.actorCtx().Send(tick.stat(), &outcome)
	n.addMetric(tick.GetTree().GetTitile(), name, status)
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

func (n *Action) recoverPanic(name string, tick *Tick) {
	if r := recover(); r != nil {
		var outcome transfer.Outcome
		outcome.Name = name
		outcome.Status = transfer.STATUS_FAILURE
		outcome.Class = transfer.CLASS_HANDLER
		outcome.Err = fmt.Sprintf("receive panic: %v", r)
		tick.actorCtx().Send(tick.stat(), &outcome)
		n.addMetric(tick.GetTree().GetTitile(), name, behavior3go.ERROR)
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
