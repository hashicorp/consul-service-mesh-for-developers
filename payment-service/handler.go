package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

func handler(rw http.ResponseWriter, r *http.Request) {
	logger.Info("Handling request")

	serverSpan := opentracing.StartSpan("handle_request")
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
