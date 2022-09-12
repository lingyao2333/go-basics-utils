package trace

import (
	"github.com/lingyao2333/go-basics-utils/options"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type config struct {
	tracerProvider trace.TracerProvider
	propagator     propagation.TextMapPropagator
}

// WithPropagator with tracer propagator.
func WithPropagator(propagator propagation.TextMapPropagator) options.Option {
	return func(o interface{}) {
		o.(*config).propagator = propagator
	}
}

// WithTracerProvider with tracer provider.
func WithTracerProvider(provider trace.TracerProvider) options.Option {
	return func(o interface{}) {
		o.(*config).tracerProvider = provider
	}
}
