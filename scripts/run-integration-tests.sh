#!/bin/bash

# Script to run integration tests for the mathreleaser application
# This script builds the binary and runs all integration tests

set -e

# Go to the project root directory
cd "$(dirname "$0")/.."

# Check if running in CI environment
if [ -n "$CI" ]; then
  # Simpler output for CI
  echo "Running integration tests in CI environment..."
else
  echo "====================================================="
  echo "        MATHRELEASER INTEGRATION TEST RUNNER         "
  echo "====================================================="
  echo
fi

# Build the binary if it doesn't exist
if [ ! -f bin/mathreleaser-test ]; then
  echo "Building the application binary..."
  go build -o bin/mathreleaser-test ./cmd/mathreleaser
  echo "Binary built successfully at bin/mathreleaser-test"
  echo
fi

# Run the integration tests
echo "Running integration tests..."

# Set log file based on environment
if [ -n "$CI" ]; then
  # In CI, output directly to console for better visibility
  go test -v ./integration_tests/...
  TEST_EXIT_CODE=$?
else
  # For local runs, save to log file
  go test -v ./integration_tests/... | tee integration-test-results.log
  TEST_EXIT_CODE=${PIPESTATUS[0]}
  echo
  echo "Test results saved to integration-test-results.log"
fi

# Check if tests passed
if [ $TEST_EXIT_CODE -eq 0 ]; then
  if [ -z "$CI" ]; then
    echo
    echo "====================================================="
    echo "      ✅ INTEGRATION TESTS PASSED SUCCESSFULLY       "
    echo "====================================================="
  else
    echo "Integration tests passed successfully."
  fi
  exit 0
else
  if [ -z "$CI" ]; then
    echo
    echo "====================================================="
    echo "      ❌ INTEGRATION TESTS FAILED                    "
    echo "====================================================="
  else
    echo "Integration tests failed."
  fi
  exit 1
fi
