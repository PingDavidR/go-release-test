# GitHub Actions Workflow Optimization Plan

## Current State

The repository currently has multiple workflows that overlap in functionality and are triggered by the same events:

1. **ci.yml** - Runs tests and builds the application
2. **gosec.yml** - Runs GoSec security scanner
3. **security-checks.yml** - Runs GoSec, govulncheck, and CodeQL
4. **codeql-analysis.yml** - Runs CodeQL analysis
5. **shell-lint.yml** - Runs shellcheck and shfmt for shell scripts
6. **release.yml** - Creates releases when tags are pushed

## Issues Identified

- Multiple security scanning workflows running on the same events
- Duplicate setup steps across workflows
- No clear dependency chain between related jobs
- Inefficient resource usage due to overlapping jobs

## Optimization Plan

### 1. Consolidated Workflows

#### New `main-pipeline.yml`

Combines the following workflows into a single pipeline with job dependencies:

- `ci.yml`
- `gosec.yml`
- `security-checks.yml`
- `codeql-analysis.yml`

The new workflow:

- Has a clear job dependency chain: test → security scanning → build
- Eliminates duplicate setup steps
- Runs security checks in parallel after tests pass
- Builds only after all tests and security checks pass

#### Retain as Specialized Workflows

- `shell-lint.yml` - Kept separate as it's path-triggered for shell scripts
- `release.yml` - Kept separate as it's triggered on tag pushes

### 2. Workflow Removal Plan

After the consolidated workflow is tested and verified, the following workflows should be removed:

- `ci.yml`
- `gosec.yml`
- `security-checks.yml`
- `codeql-analysis.yml`

### 3. Benefits

- Reduced CI/CD maintenance overhead
- Clearer visualization of the entire pipeline
- More efficient resource usage
- Faster feedback on failures
- Improved pipeline reliability

### 4. Implementation Steps

1. Create and test the new `main-pipeline.yml` workflow
2. Update documentation to reflect the new CI/CD process
3. After verification, remove the redundant workflows
4. Monitor the consolidated pipeline for any issues
