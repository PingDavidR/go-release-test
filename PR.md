# PR: Refine Release Notes Format

This PR improves the release notes process to support two different formats:

## GitHub Release Format

- Creates a simple GitHub release note (`GITHUB_RELEASE.md`) without headers
- Only shows essential content (changes, enhancements, etc.)
- Adds a clean commit hash link at the bottom
- Formatted specifically for GitHub releases UI

## Human-Readable Format

- Creates a more detailed human-friendly release note (`HUMAN_RELEASE_NOTES.md`)
- Includes proper headings and sections
- Preserves all details from the original RELEASE_NOTES.adoc
- Maintained in the repository for future reference

## Changes Made

- Added new `format-release-notes.sh` script to transform AsciiDoc to both formats
- Updated the Makefile with a `generate-release-notes` target
- Updated GitHub Actions workflow to use the new release notes format
- Integrated this process with the existing release workflows

## Testing

The script has been tested with the existing v0.1.0 release notes.
