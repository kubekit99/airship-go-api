
all: test health

# Run tests
test: generate fmt vet 
	go test ./health/... -coverprofile health.cover.out
	go test ./shipyard/... -coverprofile shipyard.cover.out
	go test ./drydock/... -coverprofile drydock.cover.out
	go test ./armada/... -coverprofile armada.cover.out
	go test ./deckhand/... -coverprofile deckhand.cover.out

health: generate fmt vet
	go build -o bin/health github.com/kubekit99/airship-go-api/health/
	go build -o bin/shipyard github.com/kubekit99/airship-go-api/shipyard/
	go build -o bin/drydock github.com/kubekit99/airship-go-api/drydock/
	go build -o bin/armada github.com/kubekit99/airship-go-api/armada/
	go build -o bin/deckhand github.com/kubekit99/airship-go-api/deckhand/

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet
	go run ./health/main.go

# Run go fmt against code
fmt:
	go fmt ./health/... ./shipyard/... ./drydock/... ./armada/... ./deckhand/...

# Run go vet against code
vet:
	go vet ./health/... ./shipyard/... ./drydock/... ./armada/... ./deckhand/...

# Generate code
generate:
ifndef GOPATH
	$(error GOPATH not defined, please define GOPATH. Run "go help gopath" to learn more about GOPATH)
endif
	go generate ./health/... ./shipyard/... ./drydock/... ./armada/... ./deckhand/...
