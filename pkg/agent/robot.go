package agent

import (
	"context"
	"sync"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/LilithGames/agent-go/pkg/transfer"
	"github.com/magicsea/behavior3go"
	"github.com/magicsea/behavior3go/core"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
)

var oc = metric.Must(global.Meter("agent")).NewInt64Counter("agent_stress_task_outcome")

type job struct {
	statPID   *actor.PID
	tree      *core.BehaviorTree
	waitGroup *sync.WaitGroup
	ctx       context.Context
	market    *Market
	alert     Alert
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

func (t *job) withAlert(alert Alert) *job {
	t.alert = alert
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
	tick := NewTick()
	tick.market = r.task.market
	tick.ctx = r.task.ctx
	tick.statPID = r.task.statPID
	tick.actorRootContext = rootCtx
	tick.alert = r.task.alert
	var status behavior3go.Status
	for {
		status = r.task.tree.Tick(tick, board)
		if status != behavior3go.RUNNING {
			break
		}
		time.Sleep(time.Second)
	}
	r.addMetric(r.task.tree.GetTitile(), status)
	outcome.Status = transfer.STATUS(status)
	outcome.Consume = time.Since(start).Nanoseconds()
	outcome.Class = transfer.CLASS_HANDLER
	rootCtx.Send(r.task.statPID, &outcome)
	r.task.waitGroup.Done()
}

func (r *robot) addMetric(name string, status behavior3go.Status) {
	tl := attribute.String("task", name)
	ctx := context.Background()
	switch status {
	case behavior3go.FAILURE, behavior3go.ERROR:
		oc.Add(ctx, 1, tl, attribute.String("status", "failure"))
	case behavior3go.SUCCESS:
		oc.Add(ctx, 1, tl, attribute.String("status", "success"))
	}
}
