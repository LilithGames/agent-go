package agent

import (
	"context"
	"fmt"
	"github.com/LilithGames/agent-go/internal/transfer"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"os"
)

type proxyStream struct {
	grpc.ClientStream
	index      int
	count      int
	id         string
	quantities *quantities
	ctx        context.Context
	cancel     context.CancelFunc
	client     transfer.Courier_DeliverMailClient
}

func newProxyStream(client transfer.Courier_DeliverMailClient) *proxyStream {
	proxy := &proxyStream{client: client, id: os.Getenv("ID")}
	proxy.ctx, proxy.cancel = context.WithCancel(context.Background())
	return proxy
}

func (s *proxyStream) setPlanCount(count int) {
	s.index = 1
	s.count = count
}

func (s *proxyStream) sendFinish(planName string) error {
	defer func() {
		if s.index >= s.count {
			s.cancel()
		}
		s.index++
	}()
	if s.client == nil {
		return s.echoLocalData()
	}
	planID := s.formatPlanID(planName)
	mail := &transfer.Mail{Action: transfer.ACTION_FINISH_PLAN, Content: []byte(planID)}
	err := s.client.Send(mail)
	if err != nil {
		return fmt.Errorf("send finish message: %w", err)
	}
	if s.index >= s.count {
		err = s.client.CloseSend()
		if err != nil {
			return fmt.Errorf("send close message: %w", err)
		}
	}
	return nil
}

func (s *proxyStream) sendReport(planName string, report *transfer.Report) error {
	if s.client == nil {
		return s.pushLocalData(report)
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

func (s *proxyStream) pushLocalData(report *transfer.Report) error {
	if s.quantities == nil {
		s.quantities = newQuantities()
	}
	putReportData(s.quantities, report.Outcomes)
	return nil
}

func (s *proxyStream) echoLocalData() error {
	printQuantities(s.quantities)
	s.quantities = newQuantities()
	return nil
}

func (s *proxyStream) formatPlanID(planName string) string {
	return fmt.Sprintf("%s-%d-%s", s.id, s.index, planName)
}
