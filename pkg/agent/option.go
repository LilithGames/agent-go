package agent

import (
	"fmt"

	"github.com/olekukonko/tablewriter"
)

type ViewOpt struct {
	tableApply []func(*tablewriter.Table)
}

func mergeViewOpt(opts ...*ViewOpt) *ViewOpt {
	res := &ViewOpt{}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		res.tableApply = append(res.tableApply, opt.tableApply...)
	}
	return res
}

func (o *ViewOpt) apply(t *tablewriter.Table) {
	for _, f := range o.tableApply {
		if f == nil {
			continue
		}
		f(t)
	}
}

func ViewColWidth(v int) *ViewOpt {
	return &ViewOpt{
		tableApply: []func(*tablewriter.Table){
			func(t *tablewriter.Table) {
				t.SetColWidth(v)
			},
		},
	}
}

type ErrMsg struct {
	Name   string `json:"name,omitempty"`
	Intro  string `json:"intro,omitempty"`
	Detail string `json:"detail,omitempty"`
}

func (e *ErrMsg) Error() string {
	return fmt.Sprintf("name: %s, intro: %s, detail: %s", e.Name, e.Intro, e.Detail)
}

type Alert interface {
	SendMsg(msg *ErrMsg) error
}

type Option func(agent *Agent)

func WithAlert(alert Alert) Option {
	return func(agent *Agent) {
		agent.alert = alert
	}
}

func WithViewer(views ...*ViewOpt) Option {
	view := mergeViewOpt(views...)
	return func(agent *Agent) {
		agent.view = view
	}
}
