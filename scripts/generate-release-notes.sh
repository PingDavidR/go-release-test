#!/usr/bin/env bash

# Check if a version argument is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <version>"
  echo "Example: $0 v1.0.0"
  exit 1
fi

VERSION=$1
CHANGELOG_DIR=".changelog"
RELEASE_NOTES_DIR="release-notes/$VERSION"
GITHUB_OUTPUT_FILE="GITHUB_RELEASE_NOTES.md"
ASCIIDOC_OUTPUT_FILE="$RELEASE_NOTES_DIR/RELEASE_NOTES.adoc"
TEMP_GITHUB_FILE=$(mktemp)
TEMP_ADOC_FILE=$(mktemp)

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

# Create release notes directory
mkdir -p "$RELEASE_NOTES_DIR"

# Ensure the changelog directory exists
if [ ! -d "$CHANGELOG_DIR" ]; then
  echo "Error: Changelog directory '$CHANGELOG_DIR' not found."
  exit 1
fi

# Check if there are any changelog files
FILE_COUNT=$(find "$CHANGELOG_DIR" -maxdepth 1 -name "*.txt" | wc -l)
if [ "$FILE_COUNT" -eq 0 ]; then
  echo "Error: No changelog files found in '$CHANGELOG_DIR'."
  exit 1
fi

# Create release notes headers
# GitHub release notes - simple header
echo "# Release Notes for $VERSION" >"$TEMP_GITHUB_FILE"
echo "" >>"$TEMP_GITHUB_FILE"

# AsciiDoc release notes - more detailed header
{
  echo "= Release Notes for $VERSION"
  echo ":toc:"
  echo ":toclevels: 3"
  echo ":sectnums:"
  echo ""
  echo "Release Date: $(date +"%Y-%m-%d")"
  echo ""
} >"$TEMP_ADOC_FILE"

# Initialize sections
BREAKING_CHANGES=""
FEATURES=""
ENHANCEMENTS=""
BUGS=""
NOTES=""
SECURITY=""
DEPRECATION=""

# Process each changelog file
for FILE in "$CHANGELOG_DIR"/*.txt; do
  if [ -f "$FILE" ]; then
    # Get the filename without extension for the PR link
    FILENAME=$(basename "$FILE" .txt)

    # Extract content from the file
    while IFS= read -r line || [ -n "$line" ]; do
      if [[ $line == '```release-note:breaking-change' ]]; then
        SECTION="BREAKING_CHANGES"
        continue
      elif [[ $line == '```release-note:feature' ]]; then
        SECTION="FEATURES"
        continue
      elif [[ $line == '```release-note:enhancement' ]]; then
        SECTION="ENHANCEMENTS"
        continue
      elif [[ $line == '```release-note:bug' ]]; then
        SECTION="BUGS"
        continue
      elif [[ $line == '```release-note:note' ]]; then
        SECTION="NOTES"
        continue
      elif [[ $line == '```release-note:security' ]]; then
        SECTION="SECURITY"
        continue
      elif [[ $line == '```release-note:deprecation' ]]; then
        SECTION="DEPRECATION"
        continue
      elif [[ $line == '```' ]]; then
        SECTION=""
        continue
      fi

      if [ -n "$SECTION" ] && [ -n "$line" ]; then
        # Generate a mock commit hash for demonstration purposes
        # In a real implementation, you would get this from git
        COMMIT_HASH=$(echo "$FILENAME-$SECTION-$RANDOM" | md5sum | cut -c1-8)

        # Extract Jira ticket if present (CDI-## or PDI-##)
        JIRA_TICKET=""
        if [[ $line =~ (CDI-[0-9]+|PDI-[0-9]+) ]]; then
          JIRA_TICKET="${BASH_REMATCH[0]}"
        fi

        # If the line doesn't contain a Jira ticket, keep the original content
        DISPLAY_LINE="$line"

        # If the line contains a Jira ticket, remove it from the display line
        if [ -n "$JIRA_TICKET" ]; then
          DISPLAY_LINE=$(echo "$line" | sed -E "s/(CDI-[0-9]+|PDI-[0-9]+)//g" | sed -E "s/^[[:space:]]+|[[:space:]]+$//g")
        fi

        # Format for each section: commit hash link, description, PR number link, Jira ticket
        ENTRY_FORMAT="* [\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DISPLAY_LINE [#$FILENAME](https://github.com/${REPO_OWNER}/${REPO_NAME}/pull/$FILENAME)"

        # Add Jira ticket if it exists
        if [ -n "$JIRA_TICKET" ]; then
          ENTRY_FORMAT="$ENTRY_FORMAT $JIRA_TICKET"
        fi

        case "$SECTION" in
          BREAKING_CHANGES)
            BREAKING_CHANGES+="$ENTRY_FORMAT\n"
            # GitHub format: commit hash link, description, PR link, Jira ticket
            echo "[\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DISPLAY_LINE [#$FILENAME](https://github.com/${REPO_OWNER}/${REPO_NAME}/pull/$FILENAME) $JIRA_TICKET" >>"$TEMP_GITHUB_FILE"
            ;;
          FEATURES)
            FEATURES+="$ENTRY_FORMAT\n"
            echo "[\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DISPLAY_LINE [#$FILENAME](https://github.com/${REPO_OWNER}/${REPO_NAME}/pull/$FILENAME) $JIRA_TICKET" >>"$TEMP_GITHUB_FILE"
            ;;
          ENHANCEMENTS)
            ENHANCEMENTS+="$ENTRY_FORMAT\n"
            echo "[\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DISPLAY_LINE [#$FILENAME](https://github.com/${REPO_OWNER}/${REPO_NAME}/pull/$FILENAME) $JIRA_TICKET" >>"$TEMP_GITHUB_FILE"
            ;;
          BUGS)
            BUGS+="$ENTRY_FORMAT\n"
            echo "[\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DISPLAY_LINE [#$FILENAME](https://github.com/${REPO_OWNER}/${REPO_NAME}/pull/$FILENAME) $JIRA_TICKET" >>"$TEMP_GITHUB_FILE"
            ;;
          NOTES)
            NOTES+="$ENTRY_FORMAT\n"
            echo "[\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DISPLAY_LINE [#$FILENAME](https://github.com/${REPO_OWNER}/${REPO_NAME}/pull/$FILENAME) $JIRA_TICKET" >>"$TEMP_GITHUB_FILE"
            ;;
          SECURITY)
            SECURITY+="$ENTRY_FORMAT\n"
            echo "[\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DISPLAY_LINE [#$FILENAME](https://github.com/${REPO_OWNER}/${REPO_NAME}/pull/$FILENAME) $JIRA_TICKET" >>"$TEMP_GITHUB_FILE"
            ;;
          DEPRECATION)
            DEPRECATION+="$ENTRY_FORMAT\n"
            echo "[\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DISPLAY_LINE [#$FILENAME](https://github.com/${REPO_OWNER}/${REPO_NAME}/pull/$FILENAME) $JIRA_TICKET" >>"$TEMP_GITHUB_FILE"
            ;;
        esac
      fi
    done <"$FILE"
  fi
done

# Add sections to the AsciiDoc release notes if they have content
if [ -n "$BREAKING_CHANGES" ]; then
  {
    echo "== BREAKING CHANGES"
    echo -e "$BREAKING_CHANGES"
    echo ""
  } >>"$TEMP_ADOC_FILE"
fi

if [ -n "$SECURITY" ]; then
  {
    echo "== SECURITY"
    echo -e "$SECURITY"
    echo ""
  } >>"$TEMP_ADOC_FILE"
fi

if [ -n "$FEATURES" ]; then
  {
    echo "== FEATURES"
    echo -e "$FEATURES"
    echo ""
  } >>"$TEMP_ADOC_FILE"
fi

if [ -n "$ENHANCEMENTS" ]; then
  {
    echo "== ENHANCEMENTS"
    echo -e "$ENHANCEMENTS"
    echo ""
  } >>"$TEMP_ADOC_FILE"
fi

if [ -n "$BUGS" ]; then
  {
    echo "== BUG FIXES"
    echo -e "$BUGS"
    echo ""
  } >>"$TEMP_ADOC_FILE"
fi

if [ -n "$DEPRECATION" ]; then
  {
    echo "== DEPRECATIONS"
    echo -e "$DEPRECATION"
    echo ""
  } >>"$TEMP_ADOC_FILE"
fi

if [ -n "$NOTES" ]; then
  {
    echo "== NOTES"
    echo -e "$NOTES"
    echo ""
  } >>"$TEMP_ADOC_FILE"
fi

# Write the final release notes
mv "$TEMP_GITHUB_FILE" "$GITHUB_OUTPUT_FILE"
mv "$TEMP_ADOC_FILE" "$ASCIIDOC_OUTPUT_FILE"

echo "GitHub release notes generated in $GITHUB_OUTPUT_FILE"
echo "AsciiDoc release notes generated in $ASCIIDOC_OUTPUT_FILE"

# Generate the formatted release notes files
echo "Generating formatted release notes files..."
"$(dirname "$0")/format-release-notes.sh" "$VERSION"
