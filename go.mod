module github.com/LilithGames/agent-go

go 1.16

require (
	github.com/AsynkronIT/protoactor-go v0.0.0-20210520041424-43065ace108f
	github.com/ghodss/yaml v1.0.0
	github.com/hasura/go-graphql-client v0.2.0
	github.com/magicsea/behavior3go v0.0.0-20201106103304-15430dcfecd8
	github.com/olekukonko/tablewriter v0.0.5
	github.com/rs/xid v1.3.0
	go.opentelemetry.io/contrib/instrumentation/runtime v0.22.0
	go.opentelemetry.io/otel/exporters/prometheus v0.22.0
	go.opentelemetry.io/otel/metric v0.22.0
	go.opentelemetry.io/otel/sdk/export/metric v0.22.0
	go.opentelemetry.io/otel/sdk/metric v0.22.0
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
)

replace (
	github.com/hasura/go-graphql-client => github.com/LilithGames/go-graphql-client v1.0.1
	github.com/magicsea/behavior3go => github.com/LilithGames/behavior3go v1.0.1
)
