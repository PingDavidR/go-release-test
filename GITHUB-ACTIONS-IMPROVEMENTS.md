# GitHub Actions Improvements

This PR optimizes the GitHub Actions workflows in this repository to improve efficiency and maintainability.

## Changes

1. **Created a unified pipeline workflow** (`main-pipeline.yml`) that consolidates four previously separate workflows:
   - `ci.yml` (CI testing and building)
   - `gosec.yml` (GoSec security scanning)
   - `security-checks.yml` (Multiple security checks)
   - `codeql-analysis.yml` (CodeQL analysis)

2. **Improved workflow efficiency** by:
   - Eliminating duplicate setup steps
   - Creating a logical job dependency chain
   - Running related checks in parallel where appropriate

3. **Retained specialized workflows** that have different triggers:
   - `shell-lint.yml` (path-triggered for shell scripts)
   - `release.yml` (tag-triggered for releases)

## Benefits

- **Reduced maintenance overhead**: Fewer workflow files to maintain
- **Clearer pipeline visualization**: Single view of the entire CI/CD process
- **More efficient resource usage**: Shared setup steps and optimized job dependencies
- **Faster feedback on failures**: Parallel execution where possible
- **Improved reliability**: Better job dependency management

## Implementation Plan

See the detailed implementation plan in [.github/WORKFLOW-OPTIMIZATION.md](/.github/WORKFLOW-OPTIMIZATION.md).
