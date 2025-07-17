# Go Release Test

This repository is designed for testing Go release processes. It serves as a sandbox for verifying release workflows, versioning strategies, and deployment pipelines for Go applications.

## Purpose

- Test release workflows for Go applications
- Verify versioning strategies
- Test deployment pipelines
- Document best practices for Go release processes

## Project Structure

```plaintext
.
├── cmd/gorelease/       # Main application entry point
├── pkg/                 # Public packages
│   ├── calculator/      # Calculator package with basic arithmetic operations
│   └── version/         # Version information package
├── internal/            # Private packages
│   └── helpers/         # Helper functions for internal use
├── .github/             # GitHub specific files
│   ├── workflows/       # GitHub Actions workflows
│   └── copilot-instructions.md # GitHub Copilot instructions
└── Makefile             # Build automation
```

## Getting Started

### Prerequisites

- Go 1.22 or higher
- Git
- Make (optional, for using the Makefile)

### Building the Application

```bash
# Build the application
make build

# Run tests
make test

# Build for all supported platforms
make build-all
```

### Running the Application

```bash
# Run with default operation (add)
./bin/gorelease 5 3

# Specify an operation
./bin/gorelease -op=subtract 5 3
./bin/gorelease -op=multiply 5 3
./bin/gorelease -op=divide 5 3

# Show version information
./bin/gorelease -version
```

## Release Process

The project uses GitHub Actions for continuous integration and release management.

### Creating a New Release

1. Update the version in `pkg/version/version.go` (or use `make bump-version`)
2. Commit the changes
3. Create a tag: `git tag -a v0.1.0 -m "Release v0.1.0"` (or use `make tag`)
4. Push the tag: `git push origin v0.1.0`
5. GitHub Actions will automatically build the release and publish it

### CI/CD Workflows

- **CI Workflow**: Runs on every push to main and on pull requests
  - Runs tests
  - Builds the application
  - Uploads the binary as an artifact

- **Release Workflow**: Triggered when a tag is pushed
  - Builds for all supported platforms
  - Creates a GitHub Release
  - Uploads all binaries and checksums

## License

This project is licensed under the MIT License - see the LICENSE file for details.
