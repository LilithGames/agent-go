package agent

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/LilithGames/agent-go/pkg/transfer"
	"github.com/LilithGames/agent-go/tools/log"
	"go.uber.org/zap"
)

type reporter struct {
	index    int
	capacity int
	eventNum int64
	planName string
	stream   *proxyStream
	outcomes []*transfer.Outcome
}

func reporterFactory(planName string, stream *proxyStream) func() actor.Actor {
	capacity := 100
	return func() actor.Actor {
		return &reporter{
			capacity: capacity,
			planName: planName,
			stream:   stream,
			outcomes: make([]*transfer.Outcome, capacity),
		}
	}
}

func (r *reporter) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *transfer.Outcome:
		r.appendOutcome(msg)
	case *actor.Stopped:
		r.stopReport()
	}
}

func (r *reporter) appendOutcome(outcome *transfer.Outcome) {
	if r.index >= r.capacity {
		r.sendStatMail()
		r.index, r.eventNum = 0, 0
		r.outcomes = make([]*transfer.Outcome, r.capacity)
	}
	r.outcomes[r.index] = outcome
	r.index++
	if outcome.Class == transfer.CLASS_EVENT {
		r.eventNum++
	}
}

func (r *reporter) stopReport() {
	r.sendStatMail()
	err := r.stream.finishPlan(r.planName)
	if err != nil {
		log.Error("receive error where send finish", zap.Error(err))
	}
}

func (r *reporter) sendStatMail() {
	if r.index < r.capacity {
		r.outcomes = r.outcomes[:r.index]
	}
	report := &transfer.Report{Outcomes: r.outcomes, EventNum: r.eventNum}
	err := r.stream.sendReport(r.planName, report)
	if err != nil {
		log.Error("receive error where send stat", zap.Error(err))
	}
}
