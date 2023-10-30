.PHONY: test run

Source= @find ./internal -type f -name "*.go"

GO_FILES := $(wildcard ./internal/osrm/*.go) \
			*.go


all: dist_finder ${GO_FILES}
	docker run -it --name dist_finder --rm -p8080:8080 -v ${PWD}:/app -w /app golang:latest go build


test:
	docker run -it --name dist_finder --rm -p8080:8080 -v ${PWD}:/app -w /app golang:latest go test; go test ./internal/osrm; go test ./internal/api

run: dist_finder
	docker run -it --name dist_finder --rm -p8080:8080 -v ${PWD}:/app -w /app golang:latest ./dist_finder

dist_finder:
	docker run -it --name dist_finder --rm -p8080:8080 -v ${PWD}:/app -w /app golang:latest go build