package agent

import (
	"context"
	"github.com/LilithGames/agent-go/tools/metric"
	"net/http"
	"os"
	"strconv"

	"github.com/LilithGames/agent-go/pkg/transfer"
	"github.com/LilithGames/agent-go/tools/log"
	"github.com/rs/xid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const masterAddr = "MASTER_ADDR"

type Agent struct {
	engine   *Engine
	endpoint string
	opt      *AgentOpt
}

func IsTestMode() bool {
	endpoint := os.Getenv(masterAddr)
	return endpoint == ""
}

func NewAgent(engine *Engine, opts ...*AgentOpt) *Agent {
	if len(engine.plans) == 0 {
		log.Panic("absent plans")
	}
	endpoint := os.Getenv(masterAddr)
	return &Agent{engine: engine, endpoint: endpoint, opt: mergeAgentOpt(opts...)}
}

func (a *Agent) Start() {
	if a.endpoint == "" {
		a.startDefaultAgent()
	} else {
		go a.startMetricServer()
		a.startClusterAgent()
	}
}

func (a *Agent) startDefaultAgent() {
	c := newProxyStream(nil, a.opt.getView())
	newManager(a.engine, c).startLocalService()
}

func (a *Agent) startClusterAgent() {
	conn := a.dialMaster()
	defer conn.Close()
	var ctx = a.newOutgoingContext()
	client, err := transfer.NewCourierClient(conn).DeliverMail(ctx)
	if err != nil {
		log.Panic("request grpc courier error", zap.Error(err))
	}
	var c = newProxyStream(client, a.opt.getView())
	newManager(a.engine, c).startClusterService()
}

func (a *Agent) dialMaster() *grpc.ClientConn {
	options := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(a.endpoint, options...)
	if err != nil {
		log.Panic("connect grpc error", zap.Error(err))
	}
	return conn
}

func (a *Agent) newOutgoingContext() context.Context {
	agentID := xid.New().String()
	var rc int32
	for _, plan := range a.engine.plans {
		if plan.RobotNum > rc {
			rc = plan.RobotNum
		}
	}
	robotNum := strconv.Itoa(int(rc))
	data := map[string]string{"agentID": agentID, "ID": os.Getenv("ID"), "robotNum": robotNum}
	return metadata.NewOutgoingContext(context.Background(), metadata.New(data))
}

func (a *Agent) startMetricServer() {
	exporter := metric.MetricsExport()
	http.HandleFunc("/metrics", exporter.ServeHTTP)
	err := http.ListenAndServe(":2222", nil)
	log.Panic("start metric error", zap.Error(err))
}
