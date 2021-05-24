package agent

import (
	"context"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/LilithGames/agent-go/pkg/transfer"
	"github.com/LilithGames/agent-go/tools/log"
	"github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/config"
	"github.com/magicsea/behavior3go/core"
	"math/rand"
	"time"
)

type Handler func(tick *core.Tick) (behavior3go.Status, error)

type Handlers map[string]Handler

type Action struct {
	core.Action
	name string
}

func (n *Action) Initialize(params *config.BTNodeCfg) {
	n.Action.Initialize(params)
	n.name = params.Name
}

func (n *Action) OnTick(tick *core.Tick) behavior3go.Status {
	defer n.recoverPanic()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	job := tick.GetTarget().(*job)
	handler := job.handlers[n.name]
	var outcome transfer.Outcome
	start := time.Now()
	status, err := handler(tick)
	outcome.Status = transfer.STATUS(status)
	if err != nil {
		outcome.Err = err.Error()
	}
	outcome.Name = n.name
	if status == behavior3go.RUNNING {
		outcome.Name = "polling_" + n.name
	}
	outcome.Consume = time.Since(start).Nanoseconds()

	beginKey := "begin:" + n.name
	begin := tick.Blackboard.GetMem(beginKey)
	if begin != nil && status != behavior3go.RUNNING {
		outcome.Consume = time.Since(begin.(time.Time)).Nanoseconds()
	}
	if begin == nil && status == behavior3go.RUNNING {
		tick.Blackboard.SetMem(n.name, start)
	}
	actorCtx := tick.Blackboard.GetMem("actorCtx").(actor.Context)
	actorCtx.Send(job.statPID, &outcome)
	return n.currentStatus(job.ctx, status)
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
		log.Error("recover panic error")
	}
}
