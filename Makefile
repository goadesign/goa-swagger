#! /usr/bin/make
#
# Makefile for goa-swagger
#
# Targets:
# - "depend" retrieves the Go packages needed to run the linter and tests
# - "lint" runs the linter and checks the code format using goimports
# - "test" runs the tests
# - "build" builds the docker container with the latest master from github
# - "deploy" deploys the latest build to GKE
#
# Meta targets:
# - "all" is the default target, it runs all the targets in the order above.
#
DIRS=$(shell go list -f {{.Dir}} ./...)
DEPEND=golang.org/x/tools/cmd/cover golang.org/x/tools/cmd/goimports \
	github.com/golang/lint/golint github.com/onsi/gomega \
	github.com/onsi/ginkgo github.com/onsi/ginkgo/ginkgo

.PHONY: build deploy

all: depend lint test build deploy

depend:
	@go get $(DEPEND)

lint:
	@for d in $(DIRS) ; do \
		if [ "`goimports -l $$d/*.go | tee /dev/stderr`" ]; then \
			echo "^ - Repo contains improperly formatted go files" && echo && exit 1; \
		fi \
	done
	@if [ "`golint ./... | grep -v app | grep -v "should not use dot imports" | tee /dev/stderr`" ]; then \
		echo "^ - Lint errors!" && echo && exit 1; \
	fi

test:
	@ginkgo -r --randomizeAllSpecs --failOnPending --randomizeSuites --race -skipPackage vendor

build:
	@docker build -t gcr.io/goa-swagger/service-node .

deploy:
	@gcloud docker push gcr.io/goa-swagger/service-node
	@kubectl rolling-update service-node --image=gcr.io/goa-swagger/service-node
