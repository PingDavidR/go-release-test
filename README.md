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
   release-note:type
   Description of the change
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

4. **Archiving Changelog Files**: After release, archive the changelog files:

   ```bash
   ./scripts/archive-changelog.sh v1.0.0
   ```

   This moves all changelog files to `.changelog/archive/v1.0.0/`.

For more details, see the [Changelog README](.changelog/README.md).

## Release Process

The project uses GitHub Actions for continuous integration and release management.

### Creating a New Release

1. Ensure all changes have corresponding changelog entries
2. Run `make devcheck` to verify code quality, security, and tests
3. Update the version in `pkg/version/version.go` (or use `make bump-version`)
4. Generate release notes: `./scripts/generate-release-notes.sh v1.0.0`  
   - This creates GitHub-formatted release notes and AsciiDoc documentation
5. Commit the changes and the generated release notes
6. Create a tag: `git tag -a v1.0.0 -m "Release v1.0.0"` (or use `make tag`)
7. Push the tag: `git push origin v1.0.0`
8. GitHub Actions will automatically build the release and publish it
   - Use the contents of `GITHUB_RELEASE_NOTES.md` for the GitHub release description
9. Archive changelog entries: `./scripts/archive-changelog.sh v1.0.0`

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

## Development Tools

### Code Quality Tools

- **golangci-lint**: For Go code linting
- **shellcheck**: For shell script linting
- **shfmt**: For shell script formatting

### Using the Tools

```bash
# Format all code (Go and shell scripts)
make fmt

# Format just shell scripts
make fmt-shell

# Lint all code (Go and shell scripts)
make lint

# Lint just shell scripts
make lint-shell

# Run security checks
make gosec       # Static code security analysis
make govulncheck # Vulnerability detection in dependencies

# Run all developer checks (formatting, linting, security checks, tests)
# Use this before submitting a PR
make devcheck

# Run shell script checks manually
./scripts/check-scripts.sh

# Fix shell script formatting issues
./scripts/check-scripts.sh --fix
```
# Triggering CI workflow
