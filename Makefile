# Makefile
PROGRAM_NAME=freeport

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
COMMIT=$(shell git rev-parse --short HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
TAG=$(shell git describe --tags |cut -d- -f1)

LDFLAGS = -ldflags "-X main.gitTag=${TAG} -X main.gitCommit=${COMMIT} -X main.gitBranch=${BRANCH}"

.PHONY: help cleanup dep lint build install uninstall 

.DEFAULT_GOAL := help

help: ## Display this help screen.
	@echo "Makefile available targets:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  * \033[36m%-15s\033[0m %s\n", $$1, $$2}'

dep: ## Download the dependencies.
	go mod download

build: dep ## Build freeport executable.
	mkdir -p ./bin
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o bin/${PROGRAM_NAME} ./cmd/${PROGRAM_NAME}

cleanup: ## Clean up the build directory.
	rm -f ./bin/${PROGRAM_NAME}_${GOARCH}_${GOOS}
	rmdir ./bin

lint: dep ## Lint the source files
	[ -d "$(ROOT_DIR)/golangci-lint" ] || mkdir -p "$(ROOT_DIR)/golangci-lint"
	docker run --rm \
    -v "$(ROOT_DIR)":/app \
    -v "$(ROOT_DIR)/golangci-lint/.cache":/root/.cache \
    -w /app \
    golangci/golangci-lint:v2.7.1 \
        golangci-lint run \
        -c .golangci-lint.yml \
    > "$(ROOT_DIR)/golangci-lint/unformatted-report.txt"

docker-build: ## Build docker image
	docker build -t artemshalin/freeport:${TAG} .
	docker image prune --force --filter label=stage=intermediate

docker-push: ## Push docker image to registry
	docker push artemshalin/freeport:${TAG}