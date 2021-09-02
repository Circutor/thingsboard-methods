#
# Copyright (c) 2021 Circutor S.A.
#

NAME=thingsboard-methods
GO=go
VERSION=$(shell cat ./VERSION)
GOFLAGS=-ldflags "-s -w -X main.version=$(VERSION)"

build:
	$(GO) build $(GOFLAGS) -o $(NAME) ./cmd

clean:
	rm -f $(NAME)

lint:
	golangci-lint run

test:
	go test -coverprofile=profile.cov ./...
	go tool cover -func profile.cov
	rm profile.cov
	go vet ./...
	gofmt -l .

run:
	./$(NAME)
