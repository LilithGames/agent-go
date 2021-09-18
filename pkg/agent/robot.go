package agent

import (
	"context"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/LilithGames/agent-go/pkg/transfer"
	"github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
	"sync"
	"time"
)

type job struct {
	statPID   *actor.PID
	tree      *core.BehaviorTree
	waitGroup *sync.WaitGroup
	ctx       context.Context
	market    *Market
}

func newJob() *job {
	return &job{}
}

func (t *job) withCancelCtx(ctx context.Context) *job {
	t.ctx = ctx
	return t
}

func (t *job) withWaitGroup(group *sync.WaitGroup) *job {
	t.waitGroup = group
	return t
}

func (t *job) withBeTree(tree *core.BehaviorTree) *job {
	t.tree = tree
	return t
}

func (t *job) withStatPID(pid *actor.PID) *job {
	t.statPID = pid
	return t
}

func (t *job) withMarket(market *Market) *job {
	t.market = market
	return t
}

type robot struct {
	task *job
}

func newRobot(j *job) *robot {
	return &robot{task: j}
}

func (r *robot) execute(rootCtx *actor.RootContext) {
	start := time.Now()
	outcome := transfer.Outcome{Name: "whole_process"}
	board := core.NewBlackboard()
	tick := NewTicker()
	tick.market = r.task.market
	tick.ctx = r.task.ctx
	tick.statPID = r.task.statPID
	tick.actorRootContext = rootCtx
	var status behavior3go.Status
	for {
		status = r.task.tree.Tick(tick, nil, board)
		if status != behavior3go.RUNNING {
			break
		}
		time.Sleep(time.Second)
	}
	outcome.Status = transfer.STATUS(status)
	outcome.Consume = time.Since(start).Nanoseconds()
	outcome.Class = transfer.CLASS_HANDLER
	rootCtx.Send(r.task.statPID, &outcome)
	r.task.waitGroup.Done()
}
