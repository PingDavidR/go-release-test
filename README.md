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
├── cmd/mathreleaser/    # Main application entry point
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

## Calculator Features

The calculator package provides the following operations:

| Operation | Description | Example |
|-----------|-------------|---------|
| Add | Adds two numbers | `./bin/mathreleaser -op=add 5 3` |
| Subtract | Subtracts the second number from the first | `./bin/mathreleaser -op=subtract 5 3` |
| Multiply | Multiplies two numbers | `./bin/mathreleaser -op=multiply 5 3` |
| Divide | Divides the first number by the second | `./bin/mathreleaser -op=divide 10 2` |
| Power | Raises the first number to the power of the second | `./bin/mathreleaser -op=power 2 3` |
| SquareRoot | Calculates the square root of a number | `./bin/mathreleaser -op=sqrt 16` |
| Sin | Calculates the sine of an angle (in radians) | `./bin/mathreleaser -op=sin 0` |
| Cos | Calculates the cosine of an angle (in radians) | `./bin/mathreleaser -op=cos 0` |
| Tan | Calculates the tangent of an angle (in radians) | `./bin/mathreleaser -op=tan 0.7853981634` |
| Random | Generates a cryptographically secure random number between two values | `./bin/mathreleaser -op=random 10 20` |

### Building the Application

```bash
# Build the application
make build

# Run unit tests
make test

# Run integration tests
make integration-test

# Run all tests (unit + integration)
make test-all

# Build for all supported platforms
make build-all
```

## Testing

### Unit Tests

Run the unit tests using:

```bash
make unit-test
```

This runs tests with the `-short` flag, which skips integration tests for faster execution.

### Integration Tests

The project includes comprehensive integration tests that run the application with predefined inputs and verify the outputs. These tests ensure that the application works correctly as a whole.

Integration tests are located in the `integration_tests/` directory and include:

1. Basic operation tests with various inputs
2. Workflow tests that simulate real user scenarios
3. Edge case tests with extreme inputs
4. Batch operation tests that verify sequences of calculations

To run the integration tests:

```bash
make integration-test
```

To run both unit and integration tests:

```bash
make test-all
```

The integration test results are saved to `integration-test-results.log` in the project root directory.

### CI Pipeline

The CI pipeline runs both unit tests and integration tests in parallel, providing faster feedback:

1. The `unit-test` job runs unit tests quickly with the `-short` flag
2. The `integration-test` job runs the full integration test suite
3. Subsequent jobs like security scanning and code analysis only proceed when both test suites pass

This parallel testing approach ensures comprehensive test coverage while minimizing pipeline execution time.

## Running the Application

```bash
# Run with default operation (add)
./bin/mathreleaser 5 3

# Specify an operation
./bin/mathreleaser -op=subtract 5 3
./bin/mathreleaser -op=multiply 5 3
./bin/mathreleaser -op=divide 5 3
./bin/mathreleaser -op=power 2 3
./bin/mathreleaser -op=sqrt 16
./bin/mathreleaser -op=sin 0
./bin/mathreleaser -op=cos 0
./bin/mathreleaser -op=tan 0.7853981634  # π/4 radians (45 degrees)
./bin/mathreleaser -op=random 10 20      # Generate random number between 10 and 20
```

### Git Hooks

This repository includes git hooks to ensure code quality standards are met before pushing changes:

- **pre-push**: Runs `make devcheck` before allowing a push to proceed, ensuring all code formatting, linting, security checks, and tests pass.

To install the git hooks:

```bash
# Install the git hooks
make install-hooks
```

Alternatively, you can manually copy the hooks from `.githooks/` to your local `.git/hooks/` directory and make them executable:

```bash
mkdir -p .git/hooks
cp .githooks/* .git/hooks/
chmod +x .git/hooks/*
```

```bash
# Show version information
./bin/mathreleaser -version
```

## Changelog Process

This repository uses a structured changelog process to track changes and generate release notes.

### How It Works

1. **Adding Changes**: When making a change, create a new file in the `.changelog` directory with a unique name matching the PR (e.g., `pr-1.txt`).

2. **Changelog Format**: Each changelog file should contain one or more sections using this format:

   ```plaintext
   <triple-backticks>release-note:type
   Description of the change CDI/PDI-###  <--- JIRA ticket
   <triple-backticks
   ```

   Where `type` can be one of:
   - `breaking-change`
   - `feature`
   - `enhancement`
   - `bug`
   - `note`
   - `security`
   - `deprecation`

3. **Generating Release Notes**: When creating a release, run:

   ```bash
   ./scripts/generate-release-notes.sh v1.0.0
   ```

   This will create:
   - `GITHUB_RELEASE_NOTES.md` with the format `<commit hash> <description> (PR #)` for GitHub releases
   - `release-notes/v1.0.0/RELEASE_NOTES.adoc` with a human-readable AsciiDoc version

4. **Archiving Changelog Files**: After generating the release notes, archive the changelog files:

   ```bash
   ./scripts/archive-changelog.sh v1.0.0
   ```

   The command above would move all changelog files to `.changelog/archive/v1.0.0/`.

For more details, see the [Changelog README](.changelog/README.md).

## Release Process

The project uses GitHub Actions for continuous integration and release management.

### Creating a New Release

1. Ensure all changes have corresponding changelog entries
2. Run `make full-check` to verify code quality, security, and tests
3. Update the version in `pkg/version/version.go` (or use `make bump-version`)
4. Generate release notes: `./scripts/generate-release-notes.sh v1.0.0`  
   - This creates GitHub-formatted release notes and AsciiDoc documentation
   - Review the files for accuracy and content
5. Archive the change log files by running `./scripts/archive-changelog.sh v1.0.0`
6. Add and commit the changes and the generated release notes with a comment "Release v1.0.0"
6. Create a tag: `git tag -a v1.0.0 -m "Release v1.0.0"` (or use `make tag`)
7. Push the tag: `git push origin main --tags
8. GitHub Actions will automatically build the release and publish it
   - Use the contents of `GITHUB_RELEASE_NOTES.md` for the GitHub release description

### CI/CD Workflows

- **CI Workflow**: Runs on every push to main and on pull requests
  - Runs tests
  - Performs security checks including static analysis for weak random number generation
  - Verifies code quality and formatting

- **Release Workflow**: Triggered when a tag is pushed
  - Builds for all supported platforms
  - Creates a GitHub Release
  - Uploads all binaries and checksums

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Development Tools

### Code Quality Tools

- **golangci-lint**: For Go code linting
- **shellcheck**: For shell script linting
- **shfmt**: For shell script formatting

### Using the Tools

```bash
# Build, test, and run the application
make build              # Build the application
make run                # Run the application
make clean              # Clean built binaries
make install            # Install the application to GOPATH/bin
make all                # Clean, build, and test

# Testing
make test               # Run all tests
make unit-test          # Run only unit tests (faster)
make integration-test   # Run integration tests
make test-all           # Run all tests including integration tests

# Code formatting
make fmt                # Format all code (Go and shell scripts)
make fmt-go             # Format just Go code
make fmt-shell          # Format just shell scripts

# Code quality checks
make vet                # Run go vet
make lint               # Lint all code (Go and shell scripts)
make lint-go            # Lint just Go code
make lint-shell         # Lint just shell scripts
make lint-all-shell     # Run both shellcheck and shfmt on shell scripts
make shellcheck-scripts # Check scripts with shellcheck
make shfmt-scripts      # Check scripts with shfmt

# Security checks
make gosec              # Static code security analysis
make govulncheck        # Vulnerability detection in dependencies

# Comprehensive checks
make devcheck           # Run all developer checks (formatting, linting)
make full-check         # Full check including security and tests

# Release management
make release            # Create a new release (builds for all platforms)
make build-all          # Build for all supported platforms
make generate-release-notes # Generate release notes
make tag                # Tag a new release
make bump-version       # Bump the version number

# Documentation
make docs               # Generate and serve documentation

# Git hooks
make install-hooks      # Install git hooks
```
