package main

import (
	"fmt"
	sleepy "github.com/nicholasjackson/sleepy-client"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"net/http"
)

func handler(rw http.ResponseWriter, r *http.Request) {
	logger.Info("Handling request")

	// attempt to create a span using a parent span defined in http headers
	wireContext, err := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header),
	)

	if err != nil {
		// if there is no span in the headers an error will be raised, log
		// this error
		logger.Debug("Error obtaining context, creating new span", "error", err)
	}

	// Create the span referring to the RPC client if available.
	// If wireContext == nil, a root span will be created.
	serverSpan := opentracing.StartSpan(
		"handle_request",
		ext.RPCServerOption(wireContext))
	serverSpan.LogFields(log.String("service.type", "http"))

	defer serverSpan.Finish()

	// create the upstream span
	upstreamSpan := serverSpan.Tracer().StartSpan("call_upstream",
		opentracing.ChildOf(serverSpan.Context()),
	)
	defer upstreamSpan.Finish()

	// call the upstream
	c := &sleepy.HTTP{}
	_, err = c.GET("http://some.servce.somewhere/")
	if err != nil {
		serverSpan.SetTag("error", true)
		serverSpan.LogFields(log.Error(err))
	}

	fmt.Fprint(rw, "Hello World")
}
