# Variables
BINARY_NAME=mathreleaser
VERSION=$(shell grep -o '"[0-9]\+\.[0-9]\+\.[0-9]\+"' pkg/version/version.go | tr -d '"')
GIT_COMMIT=$(shell git rev-parse --short HEAD)
BUILD_DATE=$(shell date -u '+%Y-%m-%d %H:%M:%S')
LDFLAGS=-ldflags "-X github.com/PingDavidR/go-release-test/pkg/version.Version=${VERSION} -X github.com/PingDavidR/go-release-test/pkg/version.GitCommit=${GIT_COMMIT} -X 'github.com/PingDavidR/go-release-test/pkg/version.BuildDate=${BUILD_DATE}'"
PLATFORMS=darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64

# Function to ensure required tools are installed
define ensure_tool
	@command -v $(1) >/dev/null 2>&1 || { echo "$(1) not found. Installing..."; $(2); }
endef

# Default target
.PHONY: all
all: clean build test

# Build the application
.PHONY: build
build:
	go build ${LDFLAGS} -o bin/${BINARY_NAME} ./cmd/mathreleaser

# Run the application
.PHONY: run
run: build
	./bin/${BINARY_NAME}

# Test the application
.PHONY: test
test:
	go test -v ./...

# Clean the binary
.PHONY: clean
clean:
	rm -rf bin/
	rm -rf dist/

# Create a new release
.PHONY: release
release: clean test build-all generate-release-notes
	mkdir -p dist
	cp bin/* dist/
	cd dist && \
	find . -type f -name "${BINARY_NAME}*" | xargs shasum -a 256 > checksums.txt

# Generate release notes
.PHONY: generate-release-notes
generate-release-notes:
	@TAG="v${VERSION}"; \
	echo "Generating release notes for $$TAG"; \
	./scripts/generate-release-notes.sh "$$TAG"; \
	./scripts/format-release-notes.sh "$$TAG"; \
	cp "release-notes/$$TAG/GITHUB_RELEASE.md" GITHUB_RELEASE_NOTES.md; \
	echo "Release notes generated successfully"

# Tag a new release
.PHONY: tag
tag:
	git tag -a v${VERSION} -m "Release v${VERSION}"
	git push origin v${VERSION}

# Build for all platforms
.PHONY: build-all
build-all: clean
	mkdir -p bin
	$(foreach PLATFORM,$(PLATFORMS),\
		$(eval OS_ARCH := $(subst /, ,$(PLATFORM)))\
		$(eval OS := $(word 1,$(OS_ARCH)))\
		$(eval ARCH := $(word 2,$(OS_ARCH)))\
		$(eval BINARY_SUFFIX := $(if $(filter windows,$(OS)),.exe,))\
		GOOS=$(OS) GOARCH=$(ARCH) go build ${LDFLAGS} -o bin/${BINARY_NAME}-$(OS)-$(ARCH)$(BINARY_SUFFIX) ./cmd/mathreleaser; \
	)

# Install the application
.PHONY: install
install: build
	cp bin/${BINARY_NAME} $(GOPATH)/bin/

# Install git hooks
.PHONY: install-hooks
install-hooks:
	@echo "Installing git hooks..."
	@mkdir -p .git/hooks
	@cp .githooks/* .git/hooks/ 2>/dev/null || true
	@chmod +x .git/hooks/*
	@echo "Git hooks installed successfully."

# Format the code
.PHONY: fmt
fmt: fmt-go fmt-shell

# Format Go code
.PHONY: fmt-go
fmt-go:
	go fmt ./...

# Format shell scripts
.PHONY: fmt-shell
fmt-shell:
	./scripts/check-scripts.sh --fix

# Vet the code
.PHONY: vet
vet:
	go vet ./...

# Lint the code
.PHONY: lint
lint: lint-go lint-shell

# Go linting
.PHONY: lint-go
lint-go:
	golangci-lint run

# Shell script linting
.PHONY: lint-shell
lint-shell:
	./scripts/check-scripts.sh

# Run security checks on Go code
.PHONY: gosec
gosec:
	$(call ensure_tool,gosec,go install github.com/securego/gosec/v2/cmd/gosec@latest)
	gosec -quiet ./...

# Run dependency vulnerability check
.PHONY: govulncheck
govulncheck:
	$(call ensure_tool,govulncheck,go install golang.org/x/vuln/cmd/govulncheck@latest)
	govulncheck ./...

# Developer check - run before submitting PR
.PHONY: devcheck
devcheck: fmt vet lint gosec govulncheck test lint-all-shell
	@echo "✅ All developer checks passed! Ready to submit PR."

# Check scripts with shellcheck
.PHONY: shellcheck-scripts
shellcheck-scripts:
	@echo "Running shellcheck on all shell scripts..."
	$(call ensure_tool,shellcheck,sudo apt-get update && sudo apt-get install -y shellcheck)
	@for script in $$(find . -name "*.sh" -type f -o -type f ! -path "*/.*" ! -path "*/vendor/*" ! -path "*/node_modules/*" -perm +111 -exec grep -l '^\#\!/bin/bash\|^\#\!/usr/bin/env bash' {} \; 2>/dev/null | sort -u); do \
		echo "Checking $$script"; \
		shellcheck -x "$$script" || exit 1; \
	done

# Check scripts with shfmt
.PHONY: shfmt-scripts
shfmt-scripts:
	@echo "Running shfmt on all shell scripts..."
	$(call ensure_tool,shfmt,go install mvdan.cc/sh/v3/cmd/shfmt@latest)
	@for script in $$(find . -name "*.sh" -type f -o -type f ! -path "*/.*" ! -path "*/vendor/*" ! -path "*/node_modules/*" -perm +111 -exec grep -l '^\#\!/bin/bash\|^\#\!/usr/bin/env bash' {} \; 2>/dev/null | sort -u); do \
		echo "Checking $$script"; \
		shfmt -i 2 -ci -bn -s -d "$$script" || exit 1; \
	done

# Lint all shell scripts (shellcheck and shfmt)
.PHONY: lint-all-shell
lint-all-shell: shellcheck-scripts shfmt-scripts

# Generate documentation
.PHONY: docs
docs:
	godoc -http=:6060

# Create a new version
.PHONY: bump-version
bump-version:
	@read -p "Enter new version (current: ${VERSION}): " version; \
	sed -i '' "s/Version = \"[0-9]\+\.[0-9]\+\.[0-9]\+\"/Version = \"$${version}\"/g" pkg/version/version.go
	@echo "Version bumped to $(shell grep -o '"[0-9]\+\.[0-9]\+\.[0-9]\+"' pkg/version/version.go | tr -d '"')"
