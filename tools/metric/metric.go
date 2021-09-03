package metric

import (
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric/global"
	export "go.opentelemetry.io/otel/sdk/export/metric"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	selector "go.opentelemetry.io/otel/sdk/metric/selector/simple"
	"log"
	"time"
)

func MetricsExport() *prometheus.Exporter {
	config := prometheus.Config{}
	c := controller.New(
		processor.New(
			selector.NewWithHistogramDistribution(),
			export.DeltaExportKindSelector(),
		),
	)
	if err := runtime.Start(runtime.WithMinimumReadMemStatsInterval(time.Second)); err != nil {
		log.Panic("create runtime metric error", err)
	}
	exporter, err := prometheus.New(config, c)
	if err != nil {
		log.Panic("create prometheus metric error", err)
	}
	global.SetMeterProvider(exporter.MeterProvider())
	return exporter
}
