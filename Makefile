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

run:
	./$(NAME)
