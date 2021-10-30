module github.com/LilithGames/agent-go

go 1.17

require (
	github.com/AsynkronIT/protoactor-go v0.0.0-20211018041209-5fdd594ca443
	github.com/ghodss/yaml v1.0.0
	github.com/hasura/go-graphql-client v0.4.0
	github.com/magicsea/behavior3go v0.0.0-20201106103304-15430dcfecd8
	github.com/olekukonko/tablewriter v0.0.5
	github.com/rs/xid v1.3.0
	go.opentelemetry.io/contrib/instrumentation/runtime v0.26.0
	go.opentelemetry.io/otel/exporters/prometheus v0.24.0
	go.opentelemetry.io/otel/metric v0.24.0
	go.opentelemetry.io/otel/sdk/export/metric v0.24.0
	go.opentelemetry.io/otel/sdk/metric v0.24.0
	go.uber.org/zap v1.19.1
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/Workiva/go-datastructures v1.0.53 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/klauspost/compress v1.10.3 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/orcaman/concurrent-map v0.0.0-20190107190726-7ed82d9cb717 // indirect
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.26.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	go.opentelemetry.io/otel v1.0.1 // indirect
	go.opentelemetry.io/otel/internal/metric v0.24.0 // indirect
	go.opentelemetry.io/otel/sdk v1.0.1 // indirect
	go.opentelemetry.io/otel/trace v1.0.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20210603081109-ebe580a85c40 // indirect
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	nhooyr.io/websocket v1.8.6 // indirect
)

replace (
	github.com/hasura/go-graphql-client => github.com/LilithGames/go-graphql-client v1.0.4
	github.com/magicsea/behavior3go => github.com/LilithGames/behavior3go v1.0.2
)
