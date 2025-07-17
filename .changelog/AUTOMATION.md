# Changelog Automation

This document explains the automated changelog generation process in this repository.

## How It Works

The repository has an automated system to ensure that each pull request includes a changelog entry. This automation helps maintain a consistent and thorough changelog for each release.

## Process Overview

### When Opening a PR

1. When you open a PR, the GitHub Action checks if either:
   - A changelog file exists at `.changelog/pr-{PR_NUMBER}.txt`
   - The PR has the `skip-changelog` label

2. If neither condition is met, the action adds a comment to your PR with instructions for creating a changelog entry.

### Creating a Changelog Entry Manually

You can create a changelog entry manually by:

1. Creating a file at `.changelog/pr-{PR_NUMBER}.txt`
2. Adding one or more entries in the following format:

   ```plaintext
   release-note:type
   Description of your change
   ```

   Where `type` can be one of:
   - `breaking-change` - Breaking changes that require user action
   - `feature` - New features
   - `enhancement` - Enhancements to existing functionality
   - `bug` - Bug fixes
   - `note` - General notes about the release
   - `security` - Security-related changes
   - `deprecation` - Deprecated features or functionality

### Skipping Changelog

For PRs that don't need a changelog entry (like documentation-only changes), add the `skip-changelog` label to the PR.

### Automatic Generation

When a PR is merged without a changelog entry and without the `skip-changelog` label, the system automatically:

1. Generates a changelog entry based on the PR title
2. Determines the entry type by analyzing keywords in the PR title
3. Creates and commits the changelog file

## How Entry Types Are Determined

The system uses keywords in the PR title to determine the entry type:

| Keywords | Entry Type |
|----------|------------|
| fix, bug, issue, resolve | `bug` |
| feature, add, implement | `feature` |
| improve, enhance, refactor, update | `enhancement` |
| security, secure, vulnerability | `security` |
| deprecate, deprecation | `deprecation` |
| break, breaking | `breaking-change` |
| (default) | `note` |

## Integration with Release Process

These changelog entries are automatically collected during the release process using the existing tools:

- `scripts/generate-release-notes.sh` - Generates two types of release notes from the changelog entries
- `scripts/archive-changelog.sh` - Archives changelog files after a release
