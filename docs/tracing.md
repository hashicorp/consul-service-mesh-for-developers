---
layout: default
title: Tracing
nav_order: 2
---

# Tracing

* Explain how tracing works with spans
* Show Jaeger and see proxy spans and application spans
* Add tracing for Payments
* Deploy Payments

```go
import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
  "github.com/opentracing/opentracing-go/log"
)

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
```