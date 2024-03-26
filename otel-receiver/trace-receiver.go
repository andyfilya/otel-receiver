package otel_receiver

import (
	"context"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.uber.org/zap"
	"time"
)

type andyfilyaReceiver struct {
	host         component.Host
	cancel       context.CancelFunc
	logger       *zap.Logger
	nextConsumer consumer.Traces
	cfg          *Config
}

// Start function represents a signal of the Collector telling the component to
// start its processing. context used for creating a new context to support you receiver
// component.Host can be useful during the whole lifecycle of the receiver
func (file *andyfilyaReceiver) Start(ctx context.Context, host component.Host) error {
	file.host = host
	ctx = context.Background()
	ctx, file.cancel = context.WithCancel(ctx)

	interval, _ := time.ParseDuration("15s")
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				file.logger.Info("I should start processing traces now!")
				traces := generateTraces(file.cfg.Path)
				err := file.nextConsumer.ConsumeTraces(ctx, traces)
				if err != nil {
					file.logger.Error(err.Error())
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}

// Shutdown function is cancel context
func (file *andyfilyaReceiver) Shutdown(ctx context.Context) error {
	file.cancel()

	return nil
}
