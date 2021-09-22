package agent

import (
	"context"
	"github.com/LilithGames/agent-go/pkg/transfer"
	"github.com/LilithGames/agent-go/tools/log"
	"github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/config"
	"github.com/magicsea/behavior3go/core"
	"go.uber.org/zap"
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
	defer n.recoverPanic()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)+10))
	name := n.GetName()
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

func (n *Action) recoverPanic() {
	if r := recover(); r != nil {
		log.Error("recover panic error", zap.Any("recover", r))
	}
}
