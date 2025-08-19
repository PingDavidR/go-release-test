# Changelog Management

This directory contains changelog entries for changes made to the repository. The changelog entries are used to generate release notes when a new version is released.

## How It Works

1. When you create a PR, you should create a changelog file named `pr-{PR_NUMBER}.txt` in this directory.
2. If you don't create one, GitHub Actions will automatically create one for you.
3. The changelog file should follow the format shown in the examples below.
4. The description in the changelog entry must be 95 characters or less.
5. During the release process, all changelog files are combined to generate release notes.
6. After release, all changelog files are moved to the `archive/{version}` directory.

## Creating a Changelog Entry

You can create a changelog entry using one of these two methods:

1. **Using the script**:

   ```bash
   ./scripts/create-changelog.sh <PR-NUMBER> <CHANGE-TYPE> "Your change description" [JIRA-TICKET]
   ```

   Example:

   ```bash
   ./scripts/create-changelog.sh 123 feature "Added new calculator function" CDI-456
   ```

2. **Manually creating the file**:
   Create a file named `pr-{PR_NUMBER}.txt` in the `.changelog` directory with the following format:

   ```plaintext
   <triple backticks>release-note:feature
   Your change description [JIRA-TICKET]
   <triple backticks>
   ```

   (Replace `<triple backticks>` with three backtick characters)

   Make sure to include the Jira ticket (CDI-## or PDI-##) at the end of your description if applicable.

## Changelog File Format

Changelog files use a specific format with sections marked by triple backticks and a section type. Each file can contain multiple sections. The recognized section types are:

- `release-note:breaking-change` - Breaking changes that require user action
- `release-note:feature` - New features
- `release-note:enhancement` - Enhancements to existing functionality
- `release-note:bug` - Bug fixes
- `release-note:note` - General notes about the release
- `release-note:security` - Security-related changes
- `release-note:deprecation` - Deprecated features or functionality

### Character Limit

The description text in each changelog entry must be 95 characters or less. This limit ensures that release notes are concise and readable. If your description exceeds this limit, you'll need to shorten it. The script will truncate it for you if necessary.

## Example

```plaintext
<triple backticks>release-note:feature
Added new calculator function `Multiply` that allows multiplication of two numbers. CDI-123
<triple backticks>
```

```plaintext
<triple backticks>release-note:bug
Fixed issue where division by zero would cause the application to crash instead of returning an error. PDI-456
<triple backticks>
```

## Release Process

The release process for managing changelog entries involves the following steps:

1. **Generate Release Notes**: Use the `generate-release-notes.sh` script to create release notes for the new version:

   ```bash
   ./scripts/generate-release-notes.sh vX.Y.Z
   ```

   This creates:
   - `GITHUB_RELEASE_NOTES.md` in the repository root
   - `release-notes/vX.Y.Z/RELEASE_NOTES.adoc` for human-readable documentation
   - `release-notes/vX.Y.Z/GITHUB_RELEASE.md` for GitHub releases

2. **Review the Generated Files**: Verify that the generated files contain the correct information:
   - Confirm that all changes are properly categorized
   - Check that the `release-notes/vX.Y.Z/` directory contains all expected files
   - Verify the content of `GITHUB_RELEASE_NOTES.md` (or `.adoc`)

3. **Archive Changelog Entries**: When releasing, archive the changelog entries using:

   ```bash
   ./scripts/archive-changelog.sh vX.Y.Z
   ```

   This moves all processed changelog files to the `.changelog/archive/vX.Y.Z/` directory.

This process ensures that each release contains only the changes made since the last release, and that all changelog entries are properly archived for historical reference.
