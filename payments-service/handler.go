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

	

	

	fmt.Fprint(rw, "Hello World")
}
