TARGETNAME=monitoring-service
TARGETPATH=$(GOPATH)/bin/$(TARGETNAME)

## GolangCI-Lint version
GOLANGCI_VERSION=1.26.0
GOLANGCI_COMMIT=6bd10d01fde78697441d9c11e2235f0dbb1e2822

export GO111MODULE=on
export GOSUMDB=off

ifndef LDFLAGS
LDFLAGS="${LDFLAGBUILDVERSION} ${LDFLAGBUILDCOMMIT}" 
endif

ifndef CGOENABLED
CGOENABLED=1
endif

ifndef BINARY_PATH
BINARY_PATH=$(GLOBALBINPATH)
endif

ifndef GO_OS
GO_OS=linux
endif

ifndef GO_ARCH
GO_ARCH=amd64
endif

.PHONY: all
all: dependencies localbuild test code-quality

.PHONY: build
build:
	CGO_ENABLED=${CGOENABLED} GOOS=${GO_OS} GOARCH=${GO_ARCH} go build -mod=vendor -o ${BINARY_PATH}

.PHONY: local
local: dependencies localbuild test code-quality

.PHONY: localbuild
localbuild: ## Build bin for local env
	go build -mod=vendor -o $(TARGETPATH)

.PHONY: dependencies
dependencies: ## Install dependencies via `go mod vendor`
	@echo "Installing dependencies"
	go mod vendor

.PHONY: test
test: ## Run tests for all packages (Used by CI)
	go test ./... -cover -coverprofile cover-all.out

.PHONY: code-quality
code-quality: bin/golangci-lint ## Run golang-cilint with write result to file
	./bin/golangci-lint --exclude-use-default=false --out-format checkstyle run ./... > static-analysis.xml

bin/golangci-lint: bin/golangci-lint-${GOLANGCI_VERSION}
	@ln -sf golangci-lint-${GOLANGCI_VERSION} bin/golangci-lint

bin/golangci-lint-${GOLANGCI_VERSION}:
	@mkdir -p bin
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/${GOLANGCI_COMMIT}/install.sh | BINARY=golangci-lint bash -s -- v${GOLANGCI_VERSION}
	@mv bin/golangci-lint $@

up-docker-env: ## Run in docker service, prometheus and grafana
	docker-compose \
		-f docker-compose.yml \
		up -d \