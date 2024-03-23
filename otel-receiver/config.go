package otel_receiver

import (
	"fmt"
	"os"
)

type Config struct {
	Path string `mapstructure:"path"`
}

func (c *Config) Validate() error {
	if c.Path == "" {
		return fmt.Errorf("path is empty")
	}

	f, err := os.Open(c.Path)
	if err != nil {
		fmt.Errorf("can't open file")
	}

	f.Close()
	return nil
}
