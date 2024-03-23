package otel_receiver

import (
	"go.opentelemetry.io/collector/pdata/ptrace"
	"os"
)

func generateTraces(path string) ptrace.Traces {
	file, _ := os.Open(path)
	var bytes []byte
	_, err := file.Read(bytes)
	if err != nil {
		return ptrace.NewTraces()
	}

	jsonUnmarshaler := &ptrace.JSONUnmarshaler{}

	var tr ptrace.Traces
	tr, err = jsonUnmarshaler.UnmarshalTraces(bytes)
	if err != nil {
		return ptrace.NewTraces()
	}

	return tr
}
