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
	metadata  map[string]string
	tree      *core.BehaviorTree
	waitGroup *sync.WaitGroup
	ctx       context.Context
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

func (t *job) withMetadata(metadata map[string]string) *job {
	t.metadata = metadata
	return t
}

type robot struct {
}

func newRobot() actor.Actor {
	return &robot{}
}

func (r *robot) Receive(ctx actor.Context) {
	if msg, ok := ctx.Message().(*job); ok {
		r.execute(ctx, msg)
	}
}

func (r *robot) execute(ctx actor.Context, task *job) {
	start := time.Now()
	outcome := transfer.Outcome{Name: "whole_process"}
	board := core.NewBlackboard()
	board.SetMem("actorCtx", ctx)
	board.SetMem("metadata", task.metadata)
	var status behavior3go.Status
	for {
		status := task.tree.Tick(task, board)
		if status != behavior3go.RUNNING {
			break
		}
		time.Sleep(time.Second)
	}
	task.waitGroup.Done()
	outcome.Status = transfer.STATUS(status)
	outcome.Consume = time.Since(start).Nanoseconds()
	ctx.Send(task.statPID, &outcome)
}
