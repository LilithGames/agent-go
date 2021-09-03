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
	grpc.ClientStream
	index      int
	count      int
	circle     bool
	id         string
	ctx        context.Context
	cancel     context.CancelFunc
	client     transfer.Courier_DeliverMailClient
	viewOpt    *ViewOpt
}

func newProxyStream(client transfer.Courier_DeliverMailClient, viewOpts ...*ViewOpt) *proxyStream {
	proxy := &proxyStream{client: client, id: os.Getenv("ID"), viewOpt: mergeViewOpt(viewOpts...)}
	proxy.ctx, proxy.cancel = context.WithCancel(context.Background())
	return proxy
}

func (s *proxyStream) setPlanCount(count int, circle bool) {
	s.index = 0
	s.count = count
	s.circle = circle
}

func (s *proxyStream) sendFinish(planName string) error {
	defer func() {
		s.index++
		if s.index >= s.count && !s.circle {
			s.cancel()
		}
		if s.circle {
			s.index %= s.count
		}
	}()
	if s.client == nil {
		return echoLocalData(planName, s.viewOpt)
	}
	planID := s.formatPlanID(planName)
	mail := &transfer.Mail{Action: transfer.ACTION_FINISH_PLAN, Content: []byte(planID)}
	err := s.client.Send(mail)
	if err != nil {
		return fmt.Errorf("send finish message: %w", err)
	}
	if s.index >= s.count - 1 && !s.circle {
		err = s.client.CloseSend()
		if err != nil {
			return fmt.Errorf("send close message: %w", err)
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
	err = s.client.Send(mail)
	if err != nil {
		return fmt.Errorf("send data message: %w", err)
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
