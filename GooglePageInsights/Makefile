GOPATH=$(shell go env GOPATH)

run:
	export GOPATH=$(GOPATH):${PWD}; \
	cd src/googlepageinsights; \
	go run main.go
test:
	export GOPATH=$(GOPATH):${PWD}; \
	cd src/googlepageinsights; \
	go test -v ./...