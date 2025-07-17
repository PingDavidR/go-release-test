#!/bin/bash

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
