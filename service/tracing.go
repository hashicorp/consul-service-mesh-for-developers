package main

import (
	"github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"

	"fmt"
	"github.com/openzipkin/zipkin-go/reporter"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

func createTracingClient(jaegerURI, name, serviceURI string) error {
	var reporter reporter.Reporter

	// create the reporter
	reporter = zipkinhttp.NewReporter(fmt.Sprintf("%s/api/v2/spans", jaegerURI))

	// create our local service endpoint
	endpoint, err := zipkin.NewEndpoint(name, serviceURI)
	if err != nil {
		return fmt.Errorf("unable to create local endpoint: %+v\n", err)
	}

	// initialize our tracer
	nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		return fmt.Errorf("unable to create tracer: %+v\n", err)
	}

	// use zipkin-go-opentracing to wrap our tracer
	tracer := zipkinot.Wrap(nativeTracer)

	// optionally set as Global OpenTracing tracer instance
	opentracing.SetGlobalTracer(tracer)

	return nil
}
