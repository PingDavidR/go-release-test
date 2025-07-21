# GitHub Copilot Instructions for Go Release Testing Repository

## Repository Purpose
This repository is designed for testing Go release processes. It serves as a sandbox for verifying release workflows, versioning strategies, and deployment pipelines for Go applications.

## Code Standards

## Operations
- operate off branches when making changes
- use pull requests for code reviews
- always use MCP servers whenever possible for any operations

### Go Version
- Use Go 1.22 or later for all development
- Ensure backwards compatibility is maintained where appropriate

### Code Style
- Follow the official Go style guidelines outlined in [Effective Go](https://golang.org/doc/effective_go)
- Use `gofmt` to format all code before committing
- Maintain a clean package structure with clear separation of concerns
- Keep functions small and focused on a single responsibility
- Use meaningful variable and function names that describe their purpose

### Code security
- use https://github.com/github/codeql-action for static code analysis to identify potential security vulnerabilities on a schedule and pre-release
- use securego/gosec (https://github.com/securego/gosec) for static code analysis to identify potential security issues in Go code during development and pre-merge

### Testing
- Write unit tests for all new functionality
- Aim for at least 80% test coverage for core functionality
- Use table-driven tests when testing multiple input/output combinations
- Utilize Go's standard testing package and avoid third-party testing frameworks unless absolutely necessary

### Documentation
- Document all exported functions, types, and variables
- Follow Go's documentation standards with clear examples
- Update README.md with any new features or significant changes

## Release Process Guidelines

### Versioning
- Follow Semantic Versioning (SemVer) for all releases
- Increment version numbers appropriately:
  - MAJOR version for incompatible API changes
  - MINOR version for backward-compatible functionality additions
  - PATCH version for backward-compatible bug fixes

### Release Workflow
- Create feature branches from `main` branch
- Submit PRs for code review before merging
- Tag releases with appropriate version numbers
- Use goreleaser (https://github.com/goreleaser/goreleaser) for building and publishing releases
- Ensure all tests pass before merging to `main`
- Use GitHub Actions for CI/CD workflows to automate testing and release processes

### Artifact Management
- Build binaries for multiple platforms (Linux, macOS, Windows, Docker
- Sign all release artifacts
- Generate checksums for all release binaries
- Store release artifacts in a consistent location

## CI/CD Expectations
- Run tests on all PRs
- Perform static code analysis using golangci-lint
- Verify module dependencies are properly maintained
- Check for security vulnerabilities in dependencies
- Ensure releases pass all tests before publishing

## Module Management
- Maintain a clean go.mod and go.sum
- Regularly update dependencies to their latest versions
- Avoid unnecessary dependencies
- Use Go workspaces for multi-module development if needed

---

*Note: These instructions are intended to provide guidance for GitHub Copilot when assisting with this repository. They should be updated as the repository evolves and processes change.*
