package metric

import (
	"log"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/sdk/export/metric/aggregation"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	selector "go.opentelemetry.io/otel/sdk/metric/selector/simple"
)

func MetricsExport() *prometheus.Exporter {
	config := prometheus.Config{}
	c := controller.New(
		processor.NewFactory(
			selector.NewWithHistogramDistribution(),
			aggregation.CumulativeTemporalitySelector(),
			processor.WithMemory(true),
		),
		controller.WithCollectPeriod(time.Second * 5),
	)
	if err := runtime.Start(runtime.WithMinimumReadMemStatsInterval(time.Second * 5)); err != nil {
		log.Panic("create runtime metric error", err)
	}
	exporter, err := prometheus.New(config, c)
	if err != nil {
		log.Panic("create prometheus metric error", err)
	}
	global.SetMeterProvider(exporter.MeterProvider())
	return exporter
}
