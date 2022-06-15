package agent

import (
	"context"
	"log"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/magicsea/behavior3go/core"
)

type Signal struct {
	market   *Market
	terminal chan interface{}
}

func newSignal(m *Market) *Signal {
	terminal := make(chan interface{})
	return &Signal{terminal: terminal, market: m}
}

func (s *Signal) Wait() {
	<-s.terminal
}

func (s *Signal) Close() {
	close(s.terminal)
	s.market.releaseSeat()
}

func (s *Signal) Slave() <-chan One {
	return s.market.getSlave()
}

type One interface {
	ID() string
}

type Market struct {
	idx    int64
	hub    chan One
	acc    chan One
	master chan One
	slave  chan One
	amount int
	used   map[string]One
	signal *Signal
	sync.Mutex
}

func newMarket(amount int) *Market {
	amount = amount * 2
	return &Market{
		amount: amount,
		hub:    make(chan One, amount),
		acc:    make(chan One, amount),
		used:   make(map[string]One),
		master: make(chan One, 1),
		slave:  make(chan One, 0),
	}
}

func (h *Market) PushOne(one One) {
	h.hub <- one
}

func (h *Market) RequireOne(like func(One) bool) One {
	var i int
	for {
		one := <-h.hub
		h.hub <- one
		if like == nil {
			return one
		}
		if like(one) {
			return one
		}
		i++
		if i >= h.amount {
			return nil
		}
	}
}

func (h *Market) JoinOne(one One) {
	h.hub <- one
}

func (h *Market) UseAcc(one One) {
	h.Lock()
	h.used[one.ID()] = one
	h.Unlock()
}

func (h *Market) InviteLikeOne(like func(One) bool) One {
	for i := 0; i < h.amount; i++ {
		select {
		case one := <-h.hub:
			if like == nil {
				return one
			}
			if like(one) {
				return one
			} else {
				h.hub <- one
			}
		case <-time.After(time.Millisecond * 10):
			return nil
		}
	}
	return nil
}

func (h *Market) InviteAcc() One {
	select {
	case one := <-h.acc:
		h.UseAcc(one)
		return one
	case <-time.After(time.Millisecond * 10):
		return nil
	}
}

func (h *Market) InviteOne() One {
	return h.InviteLikeOne(nil)
}

func (h *Market) GrabSeat(one One) (*Signal, bool) {
	select {
	case h.master <- one:
		h.signal = newSignal(h)
		return h.signal, true
	case h.slave <- one:
		return h.signal, false
	}
}

func (h *Market) releaseSeat() bool {
	select {
	case <-h.master:
		return true
	default:
		return false
	}
}

func (h *Market) getSlave() <-chan One {
	return h.slave
}

func (h *Market) reset() {
	h.idx = 0
	for _, o := range h.used {
		h.acc <- o
	}
	h.hub = make(chan One, h.amount)
	h.used = map[string]One{}
	h.master = make(chan One, 1)
	h.slave = make(chan One, 0)
}

func (h *Market) Index() int {
	return int(atomic.AddInt64(&h.idx, 1))
}

type Ticker interface {
	core.Ticker
	Market() *Market
	context() context.Context
	stat() *actor.PID
	actorCtx() *actor.RootContext
	RecvTime(unixNano string)
	SendTime(unixNano string)
}

type Tick struct {
	core.Tick
	market           *Market
	ctx              context.Context
	actorRootContext *actor.RootContext
	statPID          *actor.PID
	recvTime         int64
	sendTime         int64
	alert            Alert
}

func NewTick() *Tick {
	tick := &Tick{}
	tick.Initialize()
	return tick
}

// 用于并行的情况，分裂解决并发问题，一个行为树协程上下文使用一个tick
func (t *Tick) Tear(ticker core.Ticker) {
	tick := ticker.(*Tick)
	tick.market = t.market
	tick.ctx = t.ctx
	tick.statPID = t.statPID
	tick.actorRootContext = t.actorRootContext
	tick.alert = t.alert
	t.Tick.Tear(&tick.Tick)
}

func (t *Tick) TearTick() core.Ticker {
	tick := NewTick()
	t.Tear(tick)
	return tick
}

func (t *Tick) Market() *Market {
	return t.market
}

func (t *Tick) context() context.Context {
	return t.ctx
}

func (t *Tick) stat() *actor.PID {
	return t.statPID
}

func (t *Tick) actorCtx() *actor.RootContext {
	return t.actorRootContext
}

func (t *Tick) RecvTime(unixNano string) {
	t.recvTime = strToInt64(unixNano)
}

func (t *Tick) SendTime(unixNano string) {
	t.sendTime = strToInt64(unixNano)
}

func (t *Tick) SendAlertMsg(msg *ErrMsg) error {
	if t.alert != nil {
		return t.alert.SendMsg(msg)
	}
	log.Println("alert object is nil...")
	return nil
}

func strToInt64(s string) int64 {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return v
}
