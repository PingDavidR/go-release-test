#!/usr/bin/env bash
#
# Script to create a changelog entry for a PR
#
# Usage: ./create-changelog.sh <pr-number> <change-type> <message>
#
# Example: ./create-changelog.sh 123 feature "Added new feature X"
#
# Change types:
#   - breaking-change
#   - feature
#   - enhancement
#   - bug
#   - note
#   - security
#   - deprecation

set -euo pipefail

# Check if the required arguments are provided
if [ $# -lt 3 ]; then
  echo "Usage: $0 <pr-number> <change-type> <message>"
  echo ""
  echo "Change types:"
  echo "  - breaking-change"
  echo "  - feature"
  echo "  - enhancement"
  echo "  - bug"
  echo "  - note"
  echo "  - security"
  echo "  - deprecation"
  echo ""
  echo "Example: $0 123 feature \"Added new feature X\""
  exit 1
fi

PR_NUMBER="$1"
CHANGE_TYPE="$2"
MESSAGE="$3"
CHANGELOG_DIR=".changelog"
CHANGELOG_FILE="${CHANGELOG_DIR}/pr-${PR_NUMBER}.txt"

# Validate change type
VALID_TYPES=("breaking-change" "feature" "enhancement" "bug" "note" "security" "deprecation")
VALID=0
for TYPE in "${VALID_TYPES[@]}"; do
  if [ "$CHANGE_TYPE" = "$TYPE" ]; then
    VALID=1
    break
  fi
done

if [ $VALID -eq 0 ]; then
  echo "Error: Invalid change type '$CHANGE_TYPE'"
  echo "Valid change types: ${VALID_TYPES[*]}"
  exit 1
fi

# Create changelog directory if it doesn't exist
mkdir -p "$CHANGELOG_DIR"

# Check if file already exists
if [ -f "$CHANGELOG_FILE" ]; then
  echo "Warning: Changelog file '$CHANGELOG_FILE' already exists."
  read -r -p "Do you want to overwrite it? (y/n): " CONFIRM
  if [ "$CONFIRM" != "y" ]; then
    echo "Operation cancelled."
    exit 0
  fi
fi

# Create the changelog file
echo '```release-note:'"${CHANGE_TYPE}" > "$CHANGELOG_FILE"
echo "$MESSAGE" >> "$CHANGELOG_FILE"
echo '```' >> "$CHANGELOG_FILE"

echo "Changelog entry created at: $CHANGELOG_FILE"
echo "Content:"
echo "----------------"
cat "$CHANGELOG_FILE"
echo "----------------"
