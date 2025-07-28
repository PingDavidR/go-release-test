#!/usr/bin/env bash

# This script processes a release notes AsciiDoc file and generates:
# 1. GitHub release notes (without header, commit hash as link)
# 2. Human-friendly release notes (with header and sections)
# Updated with proper spacing for redirect operators

set -euo pipefail

# Check if a version argument is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <version>"
  echo "Example: $0 v1.0.0"
  exit 1
fi

VERSION=$1
RELEASE_NOTES_DIR="release-notes/$VERSION"
ASCIIDOC_INPUT_FILE="$RELEASE_NOTES_DIR/RELEASE_NOTES.adoc"
GITHUB_OUTPUT_FILE="$RELEASE_NOTES_DIR/GITHUB_RELEASE.md"
HUMAN_OUTPUT_FILE="$RELEASE_NOTES_DIR/HUMAN_RELEASE_NOTES.md"

# Check if the AsciiDoc file exists
if [ ! -f "$ASCIIDOC_INPUT_FILE" ]; then
  echo "Error: AsciiDoc file '$ASCIIDOC_INPUT_FILE' not found."
  exit 1
fi

# Ensure the release notes directory exists
mkdir -p "$RELEASE_NOTES_DIR"

# Get the commit hash for the version tag
COMMIT_HASH=$(git rev-list -n 1 "$VERSION" 2>/dev/null || echo "")
if [ -z "$COMMIT_HASH" ]; then
  echo "Warning: No git tag found for $VERSION, using placeholder hash"
  COMMIT_HASH="HEAD"
fi

# Short hash for display
SHORT_HASH=${COMMIT_HASH:0:7}

# Get the repository owner and name from the git remote
REPO_URL=$(git config --get remote.origin.url)
REPO_OWNER=$(echo "$REPO_URL" | sed -n 's/.*[:/]\([^/]*\)\/[^/]*\.git/\1/p')
REPO_NAME=$(echo "$REPO_URL" | sed -n 's/.*[:/][^/]*\/\([^/]*\)\.git/\1/p')

# If we couldn't extract from URL, use default values
if [ -z "$REPO_OWNER" ]; then
  REPO_OWNER="PingDavidR"
fi
if [ -z "$REPO_NAME" ]; then
  REPO_NAME="go-release-test"
fi

# Extract the release date from the AsciiDoc file
RELEASE_DATE=$(grep -i "Release Date:" "$ASCIIDOC_INPUT_FILE" | sed -E 's/Release Date: (.*)/\1/g')
if [ -z "$RELEASE_DATE" ]; then
  RELEASE_DATE=$(date +"%Y-%m-%d")
fi

# ---------------------------------------
# Create GitHub release notes file (simple format)
# ---------------------------------------
echo "Generating GitHub release notes format..."

# Start with an empty file
: >"$GITHUB_OUTPUT_FILE"

# Extract and format content sections
# More robust extraction of FEATURES section
if grep -q "== FEATURES" "$ASCIIDOC_INPUT_FILE"; then
  # Find content between FEATURES and the next section
  grep -A 100 "== FEATURES" "$ASCIIDOC_INPUT_FILE" \
    | awk '/^==/{ if (p) {exit}; p=1; next} p' \
    | grep -v "^$" \
    | sed 's/^* //' >>"$GITHUB_OUTPUT_FILE" || true
  echo "" >>"$GITHUB_OUTPUT_FILE"
fi

# More robust extraction of ENHANCEMENTS section
if grep -q "== ENHANCEMENTS" "$ASCIIDOC_INPUT_FILE"; then
  grep -A 100 "== ENHANCEMENTS" "$ASCIIDOC_INPUT_FILE" \
    | awk '/^==/{ if (p) {exit}; p=1; next} p' \
    | grep -v "^$" \
    | sed 's/^* //' >>"$GITHUB_OUTPUT_FILE" || true
  echo "" >>"$GITHUB_OUTPUT_FILE"
fi

# Add any remaining sections (NOTES, SECURITY, etc.) - customize as needed
if grep -q "== NOTES" "$ASCIIDOC_INPUT_FILE"; then
  grep -A 100 "== NOTES" "$ASCIIDOC_INPUT_FILE" \
    | awk '/^==/{ if (p) {exit}; p=1; next} p' \
    | grep -v "^$" \
    | sed 's/^* //' >>"$GITHUB_OUTPUT_FILE" || true
  echo "" >>"$GITHUB_OUTPUT_FILE"
fi

# Add bug fixes if they exist
if grep -q "== BUG FIXES" "$ASCIIDOC_INPUT_FILE"; then
  grep -A 100 "== BUG FIXES" "$ASCIIDOC_INPUT_FILE" \
    | awk '/^==/{ if (p) {exit}; p=1; next} p' \
    | grep -v "^$" \
    | sed 's/^* //' >>"$GITHUB_OUTPUT_FILE" || true
  echo "" >>"$GITHUB_OUTPUT_FILE"
fi

# Add security fixes if they exist
if grep -q "== SECURITY" "$ASCIIDOC_INPUT_FILE"; then
  grep -A 100 "== SECURITY" "$ASCIIDOC_INPUT_FILE" \
    | awk '/^==/{ if (p) {exit}; p=1; next} p' \
    | grep -v "^$" \
    | sed 's/^* //' >>"$GITHUB_OUTPUT_FILE" || true
  echo "" >>"$GITHUB_OUTPUT_FILE"
fi

# Add the commit hash link at the end
echo "[${SHORT_HASH}](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/${COMMIT_HASH})" >>"$GITHUB_OUTPUT_FILE"

# ---------------------------------------
# Create Human-friendly release notes (with sections)
# ---------------------------------------
echo "Generating human-friendly release notes..."

# Create header
{
  echo "# Release Notes for $VERSION"
  echo ""
  echo "Release Date: $RELEASE_DATE"
  echo ""
} >"$HUMAN_OUTPUT_FILE"

# Process FEATURES section
if grep -q "== FEATURES" "$ASCIIDOC_INPUT_FILE"; then
  echo "## Features" >>"$HUMAN_OUTPUT_FILE"
  echo "" >>"$HUMAN_OUTPUT_FILE"
  grep -A 100 "== FEATURES" "$ASCIIDOC_INPUT_FILE" \
    | awk '/^==/{ if (p) {exit}; p=1; next} p' \
    | grep -v "^$" \
    | sed 's/\* /- /' >>"$HUMAN_OUTPUT_FILE" || true
  echo "" >>"$HUMAN_OUTPUT_FILE"
fi

# Process ENHANCEMENTS section
if grep -q "== ENHANCEMENTS" "$ASCIIDOC_INPUT_FILE"; then
  echo "## Enhancements" >>"$HUMAN_OUTPUT_FILE"
  echo "" >>"$HUMAN_OUTPUT_FILE"
  grep -A 100 "== ENHANCEMENTS" "$ASCIIDOC_INPUT_FILE" \
    | awk '/^==/{ if (p) {exit}; p=1; next} p' \
    | grep -v "^$" \
    | sed 's/\* /- /' >>"$HUMAN_OUTPUT_FILE" || true
  echo "" >>"$HUMAN_OUTPUT_FILE"
fi

# Process NOTES section
if grep -q "== NOTES" "$ASCIIDOC_INPUT_FILE"; then
  echo "## Notes" >>"$HUMAN_OUTPUT_FILE"
  echo "" >>"$HUMAN_OUTPUT_FILE"
  grep -A 100 "== NOTES" "$ASCIIDOC_INPUT_FILE" \
    | awk '/^==/{ if (p) {exit}; p=1; next} p' \
    | grep -v "^$" \
    | sed 's/\* /- /' >>"$HUMAN_OUTPUT_FILE" || true
  echo "" >>"$HUMAN_OUTPUT_FILE"
fi

# Process SECURITY section (if exists)
if grep -q "== SECURITY" "$ASCIIDOC_INPUT_FILE"; then
  echo "## Security" >>"$HUMAN_OUTPUT_FILE"
  echo "" >>"$HUMAN_OUTPUT_FILE"
  grep -A 100 "== SECURITY" "$ASCIIDOC_INPUT_FILE" \
    | awk '/^==/{ if (p) {exit}; p=1; next} p' \
    | grep -v "^$" \
    | sed 's/\* /- /' >>"$HUMAN_OUTPUT_FILE" || true
  echo "" >>"$HUMAN_OUTPUT_FILE"
fi

# Process BUG FIXES section (if exists)
if grep -q "== BUG FIXES" "$ASCIIDOC_INPUT_FILE"; then
  echo "## Bug Fixes" >>"$HUMAN_OUTPUT_FILE"
  echo "" >>"$HUMAN_OUTPUT_FILE"
  grep -A 100 "== BUG FIXES" "$ASCIIDOC_INPUT_FILE" \
    | awk '/^==/{ if (p) {exit}; p=1; next} p' \
    | grep -v "^$" \
    | sed 's/\* /- /' >>"$HUMAN_OUTPUT_FILE" || true
  echo "" >>"$HUMAN_OUTPUT_FILE"
fi

# Add commit section
{
  echo "## Commit"
  echo ""
  echo "[${SHORT_HASH}](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/${COMMIT_HASH})"
} >>"$HUMAN_OUTPUT_FILE"

echo "Release note files created successfully:"
echo "- GitHub Release: $GITHUB_OUTPUT_FILE"
echo "- Human-Friendly: $HUMAN_OUTPUT_FILE"
