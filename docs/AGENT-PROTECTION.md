# Agent Protection System

This repository implements a protection system to prevent automated modification of critical files by AI agents or bots.

## How It Works

### 1. The `.noagent` Configuration File

The core of the protection system is the `.noagent` file in the repository root. This file lists patterns for files and directories that should not be modified by automated processes:

```plaintext
# Example entries in .noagent
go.mod
go.sum
Makefile
pkg/version/version.go
scripts/*.sh
```

### 2. Protection Enforcement

Protection is enforced through:

#### GitHub Workflow

The `.github/workflows/agent-protection.yml` workflow runs on pull requests and:

- Identifies files changed in the PR
- Checks if any changed files match patterns in `.noagent`
- Looks for commit messages that suggest automated changes
- Fails the workflow if protected files were modified by what appears to be an automated process

#### File Headers (Optional)

You can also add a special comment header to critical files:

```go
// NOAGENT: This file should not be modified by automated processes
```

## Maintaining Protected Files

### Adding Protection to Files

To protect additional files from agent modification:

1. Add the file path or glob pattern to `.noagent`
2. Optionally add the `NOAGENT` comment header to the file

### Removing Protection

To allow agents to modify previously protected files:

1. Remove the file path or pattern from `.noagent`
2. Remove any `NOAGENT` comment headers from the file

## Human Edit Override

The protection system is designed to allow human-initiated changes to protected files while blocking agent-initiated changes.

### How Human Edits Are Detected

The system primarily looks for evidence of agent or bot involvement in the commits:

- Commit messages containing terms like "automated by", "copilot change", "agent modification", etc.
- PR metadata suggesting automated generation

### Manual Override Tag

If you're making legitimate human edits to protected files and the system still blocks your PR:

1. Add `[HUMAN EDIT]` to your PR title or description
2. This tag signals to the workflow that the changes were made by a human
3. The protection check will automatically pass

### Examples

**PR Title with Override:**

```plaintext
[HUMAN EDIT] Update go.mod with critical security dependencies
```

**PR Description with Override:**

```plaintext
This PR updates the release scripts to fix a critical bug.

[HUMAN EDIT] - I'm manually updating these protected files to address security issues.
```
