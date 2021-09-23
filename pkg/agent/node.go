package agent

import (
	"context"
	"fmt"
	"github.com/LilithGames/agent-go/pkg/transfer"
	"github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/config"
	"github.com/magicsea/behavior3go/core"
	"math/rand"
	"time"
)

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
	}
}
