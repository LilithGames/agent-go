package agent

import (
	"context"
	"fmt"

	"github.com/LilithGames/agent-go/pkg/transfer"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type proxyStream struct {
	ctx context.Context
	grpc.ClientStream
	index   int
	count   int
	circle  bool
	hasErr  bool
	id      string
	client  transfer.Courier_DeliverMailClient
	viewOpt *ViewOpt
}

func newProxyFromAgent(agent *Agent, clients ...transfer.Courier_DeliverMailClient) *proxyStream {
	proxy := &proxyStream{
		id:      agent.id,
		viewOpt: agent.view,
		ctx:     agent.ctx,
	}
	if len(clients) == 1 {
		proxy.client = clients[0]
	}
	return proxy
}

func (s *proxyStream) setPlanCount(count int, circle bool) {
	s.index = 0
	s.count = count
	s.circle = circle
}

func (s *proxyStream) finishPlan(planName string) error {
	if s.client == nil {
		s.hasErr = hasLocalError()
		echoLocalData(planName, s.viewOpt)
	}
	select {
	case <-s.ctx.Done():
		return nil
	default:
		s.index++
	}
	if s.index >= s.count {
		if s.circle {
			s.index %= s.count
		} else if s.client != nil {
			mail := &transfer.Mail{Action: transfer.ACTION_FINISH_PLAN, Content: []byte(s.id)}
			err := s.client.Send(mail)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *proxyStream) sendReport(planName string, report *transfer.Report) error {
	if s.client == nil {
		return pushLocalData(report)
	}
	report.PlanID = s.formatPlanID(planName)
	content, err := proto.Marshal(report)
	if err != nil {
		return fmt.Errorf("proto marshal data: %w", err)
	}
	mail := &transfer.Mail{Action: transfer.ACTION_REPORT_DATA, Content: content}
	select {
	case <-s.ctx.Done():
		return nil
	default:
		err = s.client.Send(mail)
		if err != nil {
			return fmt.Errorf("send data message: %w", err)
		}
	}
	return nil
}

func (s *proxyStream) Recv() (*transfer.Mail, error) {
	if s.client != nil {
		return s.client.Recv()
	}
	return nil, nil
}

func (s *proxyStream) formatPlanID(planName string) string {
	return fmt.Sprintf("%s-%d-%s", s.id, s.index, planName)
}

func (s *proxyStream) checkResult() error {
	if s.hasErr {
		return fmt.Errorf("Test Failed")
	}
	return nil
}
