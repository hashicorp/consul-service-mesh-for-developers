FROM golang:alpine

COPY . /go/src/github.com/nicholasjackson/consul-service-mesh-for-developer/service

# Build with debug flags
WORKDIR /go/src/github.com/nicholasjackson/consul-service-mesh-for-developer/service
RUN CGO_ENABLED=0 GOOS=linux go build -gcflags "-N -l" -o ./bin/service ./main.go

FROM alpine:latest

RUN mkdir /service
COPY --from=0 /go/src/github.com/nicholasjackson/consul-service-mesh-for-developer/service/bin/service /app/service

WORKDIR /tmp

ENTRYPOINT ["/app/service"]
