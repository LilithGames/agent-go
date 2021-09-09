package agent

import (
	"context"
	"fmt"
	"os"

	"github.com/LilithGames/agent-go/pkg/transfer"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type proxyStream struct {
	ctx context.Context
	grpc.ClientStream
	index      int
	count      int
	circle     bool
	id         string
	client     transfer.Courier_DeliverMailClient
	viewOpt    *ViewOpt
}

func newProxyStream(client transfer.Courier_DeliverMailClient, viewOpts ...*ViewOpt) *proxyStream {
	return &proxyStream{client: client, id: os.Getenv("ID"), viewOpt: mergeViewOpt(viewOpts...)}
}

func (s *proxyStream) withContext(ctx context.Context) {
	s.ctx = ctx
}

func (s *proxyStream) setPlanCount(count int, circle bool) {
	s.index = 0
	s.count = count
	s.circle = circle
}

func (s *proxyStream) finishPlan(planName string) error {
	if s.client == nil {
		echoLocalData(planName, s.viewOpt)
	}
	select {
	case <- s.ctx.Done():
		return nil
	default:
		s.index++
	}
	if s.index >= s.count {
		if s.circle {
			s.index %= s.count
		} else if s.client != nil{
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
		return pushLocalData(planName, report)
	}
	report.PlanID = s.formatPlanID(planName)
	content, err := proto.Marshal(report)
	if err != nil {
		return fmt.Errorf("proto marshal data: %w", err)
	}
	mail := &transfer.Mail{Action: transfer.ACTION_REPORT_DATA, Content: content}
	select {
	case <- s.ctx.Done():
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
