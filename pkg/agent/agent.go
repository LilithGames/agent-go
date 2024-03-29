package agent

import (
	"context"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"strings"

	"github.com/LilithGames/agent-go/tools/metric"
	"github.com/spf13/viper"

	"github.com/LilithGames/agent-go/pkg/transfer"
	"github.com/LilithGames/agent-go/tools/log"
	"github.com/rs/xid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const masterAddr = "MASTER_ADDR"

type Agent struct {
	id       string
	ctx      context.Context
	cancel   context.CancelFunc
	engine   *Engine
	variable *variable
	endpoint string
	view     *ViewOpt
	alert    Alert
	stream   *proxyStream
}

func NewAgent(engine *Engine, cfg *viper.Viper, opts ...Option) *Agent {
	if len(engine.plans) == 0 {
		log.Panic("absent plans")
	}
	if cfg == nil {
		cfg = viper.New()
		cfg.SetConfigType("prop")
		envs := os.Environ()
		for _, env := range envs {
			parts := strings.SplitN(env, "=", 2)
			if len(parts) == 2 {
				cfg.Set(parts[0], parts[1])
			}
		}
	}
	variable := initVariable(cfg)
	id := variable.GetString("ID")
	endpoint := variable.GetString(masterAddr)
	ctx, cancel := context.WithCancel(context.Background())
	at := &Agent{
		id:       id,
		ctx:      ctx,
		cancel:   cancel,
		engine:   engine,
		variable: variable,
		endpoint: endpoint,
	}
	for _, opt := range opts {
		opt(at)
	}
	return at
}

func (a *Agent) Start() {
	go a.startMetricServer()
	if a.endpoint == "" {
		a.startDefaultAgent()
	} else {
		a.startClusterAgent()
	}
}

func (a *Agent) CheckResult() error {
	return a.stream.checkResult()
}

func (a *Agent) startDefaultAgent() {
	a.stream = newProxyFromAgent(a)
	newManagerFromAgent(a).startLocalService()
}

func (a *Agent) startClusterAgent() {
	conn := a.dialMaster()
	defer conn.Close()
	ctx := a.newOutgoingContext()
	client, err := transfer.NewCourierClient(conn).DeliverMail(ctx)
	if err != nil {
		log.Panic("request grpc courier error", zap.Error(err))
	}
	a.stream = newProxyFromAgent(a, client)
	newManagerFromAgent(a).startClusterService()
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
	data := map[string]string{"agentID": agentID, "ID": a.id, "robotNum": robotNum}
	return metadata.NewOutgoingContext(context.Background(), metadata.New(data))
}

func (a *Agent) startMetricServer() {
	exporter := metric.MetricsExport()
	http.HandleFunc("/metrics", exporter.ServeHTTP)
	err := http.ListenAndServe(":6060", nil)
	log.Panic("start metric error", zap.Error(err))
}
