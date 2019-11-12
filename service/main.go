package main

import (
	"fmt"
	"net/http"

	hclog "github.com/hashicorp/go-hclog"
)

func main() {
	l := hclog.Default()
	l.Info("Starting service 1.3")

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello World")
	})

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
