// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"net/http"
	"os"

	hclog "github.com/hashicorp/go-hclog"
)

// global logger
var logger hclog.Logger

func main() {
	logger = hclog.Default()
	logger.Info("Starting service 1.0")

	// create the tracing setup
	err := createTracingClient(os.Getenv("TRACING_ZIPKIN"), "payment", "localhost")
	if err != nil {
		logger.Error("Error creating tracing client", "error", err)
	}

	// wire up the http handler
	http.HandleFunc("/", handler)

	// start the server
	http.ListenAndServe(":8080", nil)
}

/*
squashctl --crisock /run/k3s/containerd/containerd.sock --machine
kubectl port-forward <squash pod> 10001:37579 -n squash-debugger
kubectl port-forward svc/payment-service 8081:8080
start debugging in vscode
add a breakpoint line 15
curl localhost:8081

to update:
docker build -t nicholasjackson/broken-service:v4 .
update image in application/payments.yml
yard push --image nicholasjackson/broken-service:v4
kubectl apply -f ../application/payments.yml
*/
