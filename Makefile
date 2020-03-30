-include .env

PROJECTNAME="calendar"

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## test: run all tests
test:
	@echo "  >  Running all tests"
	go test ./...

## build: build source in dist/
build:
	@echo "  >  Building binary..."
	go build -o dist/$(PROJECTNAME) cmd/$(PROJECTNAME)/main.go

## install: run go get
install:
	@echo "  >  Checking if there is any missing dependencies..."
	go get ./...

gen:
	protoc --go_out=plugins=grpc:pkg api/*.proto

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo