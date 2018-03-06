GOPATH=$(shell go env GOPATH)

run:
	export GOPATH=$(GOPATH):${PWD}; \
	cd src/googlesheets; \
	go run main.go