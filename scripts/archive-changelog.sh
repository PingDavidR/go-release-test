#!/usr/bin/env bash

# Script to archive changelog entries after a release
# Designed to work both locally and in CI environments

set -euo pipefail

# Check if a version argument is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <version>"
  echo "Example: $0 v1.0.0"
  exit 1
fi

VERSION=$1
CHANGELOG_DIR=".changelog"
ARCHIVE_DIR="$CHANGELOG_DIR/archive/$VERSION"

# Ensure the changelog directory exists
if [ ! -d "$CHANGELOG_DIR" ]; then
  echo "Error: Changelog directory '$CHANGELOG_DIR' not found."
  exit 1
fi

# Check if there are any changelog files to archive
FILE_COUNT=$(find "$CHANGELOG_DIR" -maxdepth 1 -name "*.txt" | wc -l)
if [ "$FILE_COUNT" -eq 0 ]; then
  echo "Error: No changelog files found in '$CHANGELOG_DIR' to archive."
  exit 1
fi

# Create the archive directory
mkdir -p "$ARCHIVE_DIR"

# Move all changelog files to the archive directory
for FILE in "$CHANGELOG_DIR"/*.txt; do
  if [ -f "$FILE" ]; then
    FILENAME=$(basename "$FILE")
    mv "$FILE" "$ARCHIVE_DIR/$FILENAME"
    echo "Archived $FILENAME to $ARCHIVE_DIR"
  fi
done

echo "All changelog files have been archived to $ARCHIVE_DIR"

# Check if we're in a CI environment and in detached HEAD state
# In CI, we'll just archive without committing
if [ -n "${CI:-}" ] || [ "$(git symbolic-ref --short HEAD 2>/dev/null || echo 'detached')" = "detached" ]; then
  echo "Running in CI or detached HEAD state, skipping commit step"
  exit 0
fi

# If running locally with a branch, commit the archive changes
echo "Committing archive changes..."
git add "$ARCHIVE_DIR"
# Also stage deletions from .changelog directory
git add -u "$CHANGELOG_DIR"
git commit -m "archive-changelog-files: Archive changelog entries for $VERSION" || echo "No changes to commit"
echo "Changelog entries archived and committed successfully"
