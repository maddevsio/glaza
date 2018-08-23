GOPATH=$(shell go env GOPATH)

run:
	export GOPATH=$(GOPATH):${PWD}; \
	cd src/googleanalytics; \
	go run main.go
test:
	export GOPATH=$(GOPATH):${PWD}; \
	cd src/googleanalytics; \
	go test -v ./...