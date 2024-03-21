package otel_receiver

import (
	"fmt"
	"time"
)

type Config struct {
	Interval       string `mapstructure:"interval"`
	NumberOfTraces uint   `mapstructure:"number_of_traces"`
}

func (cfg *Config) Validate() error {
	interval, err := time.ParseDuration(cfg.Interval)
	if err != nil {
		return fmt.Errorf("error parse time in interval, please, make it correct")
	}

	if interval.Minutes() < 1 {
		return fmt.Errorf("if interval set, it has to be set at least one minute")
	}

	if cfg.NumberOfTraces < 1 {
		return fmt.Errorf("if number of traces is set, it has to be at least one trace")
	}

	return nil
}
