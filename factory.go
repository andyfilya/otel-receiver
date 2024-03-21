package otel_receiver

import (
	"context"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"time"
)

const (
	typeStr         = "otel-receiver"
	defaultInterval = 1 * time.Minute
)

func createDefaultConfig() component.Config {
	return &Config{
		Interval: string(defaultInterval),
	}
}

func createTracerReceiver(_ context.Context, params receiver.CreateSettings, baseCfg component.Config, consumer consumer.Traces) (receiver.Traces, error) {
	logger := params.Logger
	otrvrCfg := baseCfg.(*Config)

	trRvr := &oteltracerReceiver{
		logger:       logger,
		nextConsumer: consumer,
		config:       otrvrCfg,
	}

	return trRvr, nil
}

// NewFactory creates a new factory for otel receiver
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithTraces(createTracerReceiver, component.StabilityLevelAlpha),
	)
}
