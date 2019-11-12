build_linux:
	CGO_ENABLED=0 GOOS=linux go build -gcflags "-N -l" -o ./bin/service ./main.go

build_docker:
	docker build -t nicholasjackson/broken-service:latest .

build: build_linux build_docker
