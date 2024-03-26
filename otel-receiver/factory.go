package otel_receiver

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const (
	typeAndy = "andyfilya"
)

// NewFactory creates a factory for file receiver
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeAndy,
		createDefaultConfig,
		receiver.WithTraces(createTracesReceiver, component.StabilityLevelAlpha))
}

func createDefaultConfig() component.Config {
	return &Config{
		Path: "trace.json",
	}
}

func createTracesReceiver(_ context.Context, settings receiver.CreateSettings, configuration component.Config, consumer consumer.Traces) (receiver.Traces, error) {
	logger := settings.Logger
	cfg := configuration.(*Config)

	tr := &andyfilyaReceiver{
		logger:       logger,
		nextConsumer: consumer,
		cfg:          cfg,
	}

	return tr, nil
}
