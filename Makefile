
all: test health

# Run tests
test: generate fmt vet 
	go test ./shipyard/... -coverprofile shipyard.cover.out
	go test ./drydock/... -coverprofile drydock.cover.out
	go test ./armada/... -coverprofile armada.cover.out
	go test ./deckhand/... -coverprofile deckhand.cover.out
	go test ./promenade/... -coverprofile promenade.cover.out

health: generate fmt vet
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/shipyard github.com/kubekit99/airship-go-api/shipyard/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/drydock github.com/kubekit99/airship-go-api/drydock/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/armada github.com/kubekit99/airship-go-api/armada/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/deckhand github.com/kubekit99/airship-go-api/deckhand/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/promenade github.com/kubekit99/airship-go-api/promenade/

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet
	go run ./armada/main.go

# Run go fmt against code
fmt:
	go fmt ./shipyard/... ./drydock/... ./armada/... ./deckhand/... ./promenade/...

# Run go vet against code
vet:
	go vet ./shipyard/... ./drydock/... ./armada/... ./deckhand/... ./promenade/...

# Generate code
generate:
ifndef GOPATH
	$(error GOPATH not defined, please define GOPATH. Run "go help gopath" to learn more about GOPATH)
endif
	go generate ./shipyard/... ./drydock/... ./armada/... ./deckhand/... ./promenade/...

# Build the docker image
docker-build:
	docker build . -f deckhand/Dockerfile -t kubekit99/apisim-deckhand:0.1.0
	docker build . -f promenade/Dockerfile -t kubekit99/apisim-promenade:0.1.0
	docker build . -f shipyard/Dockerfile -t kubekit99/apisim-shipyard:0.1.0
	docker build . -f drydock/Dockerfile -t kubekit99/apisim-drydock:0.1.0
	docker build . -f armada/Dockerfile -t kubekit99/apisim-armada:0.1.0

# Push the docker image
docker-push:
	docker push kubekit99/apisim-deckhand:0.1.0
	docker push kubekit99/apisim-promenade:0.1.0
	docker push kubekit99/apisim-shipyard:0.1.0
	docker push kubekit99/apisim-drydock:0.1.0
	docker push kubekit99/apisim-armada:0.1.0
