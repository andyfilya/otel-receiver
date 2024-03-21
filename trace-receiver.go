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
	config       *Config
}

// Start function represents a signal of the Collector telling the component to
// start its processing. context used for creating a new context to support you receiver
// component.Host can be useful during the whole lifecycle of the receiver
func (otRvr *andyfilyaReceiver) Start(ctx context.Context, host component.Host) error {
	otRvr.host = host
	ctx = context.Background()
	ctx, otRvr.cancel = context.WithCancel(ctx)

	interval, _ := time.ParseDuration(otRvr.config.Interval) // err is check in config
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				otRvr.logger.Info("start processing traces...")
			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}

// Shutdown function is cancel context
func (otRvr *andyfilyaReceiver) Shutdown(ctx context.Context) error {
	otRvr.cancel()
	return nil
}
