#!/usr/bin/env bash
# Wrapper to demonstrate create-changelog.sh for all change types

SCRIPT="$(dirname "$0")/create-changelog.sh"
PR_BASE=1000
CHANGELOG_DIR="$(dirname "$0")/../.changelog"

# Ensure changelog directory exists
mkdir -p "$CHANGELOG_DIR"

TYPES=(
  "breaking-change"
  "feature"
  "enhancement"
  "bug"
  "note"
  "security"
  "deprecation"
)

for i in "${!TYPES[@]}"; do
  TYPE="${TYPES[$i]}"
  PR_NUM=$((PR_BASE + i))
  MSG="Demo message for $TYPE type"
  JIRA="CDI-$((200 + i))"
  echo "Creating changelog for PR $PR_NUM, type $TYPE..."
  "$SCRIPT" "$PR_NUM" "$TYPE" "$MSG" "$JIRA"
  echo ""
done
