package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
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

	serverSpan := opentracing.StartSpan(
		"handle_request",
		ext.RPCServerOption(wireContext))
	serverSpan.LogFields(log.String("service.type", "http"))

	defer serverSpan.Finish()

	// create the request
	req, err := http.NewRequest(http.MethodGet, os.Getenv("CURRENCY_ADDR"), nil)
	if err != nil {
		logger.Error("Error creating request", "error", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	// execute the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error("Error calling upstream", "error", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		logger.Error("Expected status OK, got", "status", resp.StatusCode)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprint(rw, "Hello World")
}
