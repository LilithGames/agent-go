package agent

import "github.com/olekukonko/tablewriter"

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

type AgentOpt struct {
	view *ViewOpt
}

func AgentViewOpt(opts ...*ViewOpt) *AgentOpt {
	return &AgentOpt{
		view: mergeViewOpt(opts...),
	}
}

func (o *AgentOpt) getView() *ViewOpt {
	return o.view
}

func mergeAgentOpt(opts ...*AgentOpt) *AgentOpt {
	res := &AgentOpt{}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		res.view = mergeViewOpt(res.view, opt.view)
	}
	return res
}
