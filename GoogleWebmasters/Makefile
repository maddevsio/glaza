GOPATH=$(shell go env GOPATH)

run:
	export GOPATH=$(GOPATH):${PWD}; \
	cd src/googlewebmasters; \
	go run main.go
test:
	export GOPATH=$(GOPATH):${PWD}; \
	cd src/googlewebmasters; \
	go test -v ./...