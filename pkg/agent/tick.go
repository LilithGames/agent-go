package agent

import (
	"context"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/magicsea/behavior3go/core"
	"sync/atomic"
	"time"
)

type One interface {
	ID() string
}

type Market struct {
	idx    int64
	pre    chan One
	cur    chan One
	hub    chan One
	amount int
}

func newMarket(amount int) *Market {
	return &Market{cur: make(chan One, amount), amount: amount, hub: make(chan One, amount)}
}

func (h *Market) PushOne(one One) {
	h.cur <- one
}

func (h *Market) RequireOne(like func(One) bool) One {
	var (
		one  One
		flag bool
		i    int
	)
	for !flag {
		one = <-h.cur
		flag = like(one)
		h.cur <- one
		i++
		if i >= h.amount {
			return nil
		}
	}
	return one
}

func (h *Market) JoinOne(one One) {
	h.hub <- one
}

func (h *Market) InviteOne() One {
	select {
	case one := <-h.hub:
		return one
	case <-time.After(time.Second):
		return nil
	}
}

func (h *Market) reset() {
	h.pre = h.cur
	h.cur = make(chan One, h.amount)
	h.hub = make(chan One, h.amount)
}

func (h *Market) Index() int {
	return int(atomic.AddInt64(&h.idx, 1))
}

type Ticker interface {
	core.Ticker
	Marget() *Market
	setMarket(market *Market)
	context() context.Context
	setContext(ctx context.Context)
	stat() *actor.PID
	setStatPID(pid *actor.PID)
	RecvTime(timestamp time.Duration)
	SendTime(timestamp time.Duration)
}

type Tick struct {
	core.Tick
	market  *Market
	ctx     context.Context
	statPID *actor.PID
	recvTime time.Duration
	sendTime time.Duration
}

func NewTicker() *Tick {
	tick := &Tick{}
	tick.Initialize()
	return tick
}

func (t *Tick) Marget() *Market {
	return t.market
}

func (t *Tick) setMarket(market *Market) {
	t.market = market
}

func (t *Tick) context() context.Context {
	return t.ctx
}

func (t *Tick) setContext(ctx context.Context) {
	t.ctx = ctx
}

func (t *Tick) stat() *actor.PID {
	return t.statPID
}

func (t *Tick) setStatPID(pid *actor.PID) {
	t.statPID = pid
}

func (t *Tick) RecvTime(timestamp time.Duration) {
	t.recvTime = timestamp
}

func (t *Tick) SendTime(timestamp time.Duration) {
	t.sendTime = timestamp
}