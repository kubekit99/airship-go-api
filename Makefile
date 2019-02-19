
all: test health

# Run tests
test: generate fmt vet 
	go test ./shipyard/... -coverprofile shipyard.cover.out
	go test ./drydock/... -coverprofile drydock.cover.out
	go test ./armada/... -coverprofile armada.cover.out
	go test ./deckhand/... -coverprofile deckhand.cover.out
	go test ./promenade/... -coverprofile promenade.cover.out

health: generate fmt vet
	go build -o bin/shipyard github.com/kubekit99/airship-go-api/shipyard/
	go build -o bin/drydock github.com/kubekit99/airship-go-api/drydock/
	go build -o bin/armada github.com/kubekit99/airship-go-api/armada/
	go build -o bin/deckhand github.com/kubekit99/airship-go-api/deckhand/
	go build -o bin/promenade github.com/kubekit99/airship-go-api/promenade/

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
