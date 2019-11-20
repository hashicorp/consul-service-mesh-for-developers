package main

import (
	hclog "github.com/hashicorp/go-hclog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var called = false

func setupTests() (*httptest.Server, func()) {
	logger = hclog.Default()
	called = false

	ts := httptest.NewServer(
		http.HandlerFunc(
			func(rw http.ResponseWriter, r *http.Request) {
				called = true
				rw.Write([]byte("ok"))
			},
		),
	)

	os.Setenv("CURRENCY_ADDR", ts.URL)

	return ts, func() {
		os.Unsetenv("CURRENCY_ADDR")
		ts.Close()
	}
}

func TestHandlerCallsUpstream(t *testing.T) {
	_, done := setupTests()
	defer done()

	rr, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/", nil)

	handler(rr, r)

	if rr.Code != http.StatusOK {
		t.Fatalf("Expected status OK, got: %d", rr.Code)
	}

	if !called {
		t.Fatal("Expected upstream to be called")
	}
}
