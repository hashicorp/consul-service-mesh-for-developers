---
layout: default
title: Canary Deployments
nav_order: 3
---

* Explain previous deployment was risky as there could be a bug
* Talk about Canary deployments
* Update service to add new feature
* Deploy with Canary, see failures in tracing

```go
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
```