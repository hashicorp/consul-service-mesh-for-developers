package main

import (
	"fmt"
	"net/http"
)

func handler(rw http.ResponseWriter, r *http.Request) {
	logger.Info("Handling request")

	fmt.Fprint(rw, "Hello World")
}
