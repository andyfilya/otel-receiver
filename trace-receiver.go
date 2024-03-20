package otel_receiver

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
)

type oteltracerReceiver struct {
	host         component.Host
	cancel       context.CancelFunc
	logger       *log.Logger
	nextConsumer consumer.Traces
	config       *Config
}

func (otRvr *oteltracerReceiver) Start(ctx context.Context, host component.Host) error {
	otRvr.host = host
	ctx = context.Background()
	ctx, otRvr.cancel = context.WithCancel(ctx)

	return nil
}

func (otRvr *oteltracerReceiver) Shutdown(ctx context.Context) error {
	otRvr.cancel()

	return nil
}
