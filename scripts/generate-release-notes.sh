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
echo "# Release Notes for $VERSION" > "$TEMP_GITHUB_FILE"
echo "" >> "$TEMP_GITHUB_FILE"

# AsciiDoc release notes - more detailed header
{
    echo "= Release Notes for $VERSION"
    echo ":toc:"
    echo ":toclevels: 3"
    echo ":sectnums:"
    echo ""
    echo "Release Date: $(date +"%Y-%m-%d")"
    echo ""
} > "$TEMP_ADOC_FILE"

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

        # Initialize variables for this file
        CHANGE_TYPE=""
        DESCRIPTION=""
        PR_NUMBER=""
        JIRA_TICKET=""

        # Parse the new format: key-value pairs
        while IFS= read -r line || [ -n "$line" ]; do
            # Skip empty lines
            if [ -z "$line" ]; then
                continue
            fi

            # Parse key-value pairs
            if [[ $line =~ ^type:[[:space:]]*(.*) ]]; then
                CHANGE_TYPE="${BASH_REMATCH[1]}"
            elif [[ $line =~ ^description:[[:space:]]*(.*) ]]; then
                DESCRIPTION="${BASH_REMATCH[1]}"
            elif [[ $line =~ ^pr:[[:space:]]*(.*) ]]; then
                PR_NUMBER="${BASH_REMATCH[1]}"
            elif [[ $line =~ ^ticket:[[:space:]]*(.*) ]]; then
                JIRA_TICKET="${BASH_REMATCH[1]}"
            fi
        done < "$FILE"

        # Skip if we don't have the minimum required fields
        if [ -z "$CHANGE_TYPE" ] || [ -z "$DESCRIPTION" ]; then
            echo "Warning: Skipping $FILE - missing type or description"
            continue
        fi

        # Skip internal changes from release notes
        if [ "$CHANGE_TYPE" = "internal" ]; then
            echo "Skipping internal change from $FILE"
            continue
        fi

        # If PR number wasn't specified in the file, try to extract from filename
        if [ -z "$PR_NUMBER" ]; then
            if [[ $FILENAME =~ pr-([0-9]+) ]]; then
                PR_NUMBER="${BASH_REMATCH[1]}"
            fi
        fi

        # Generate a mock commit hash for demonstration purposes
        # In a real implementation, you would get this from git
        COMMIT_HASH=$(echo "$FILENAME-$CHANGE_TYPE-$RANDOM" | md5sum | cut -c1-8)

        # Build the entry format
        ENTRY_FORMAT="* [\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DESCRIPTION"

        # Add PR link if available
        if [ -n "$PR_NUMBER" ]; then
            ENTRY_FORMAT="$ENTRY_FORMAT [#$PR_NUMBER](https://github.com/${REPO_OWNER}/${REPO_NAME}/pull/$PR_NUMBER)"
        fi

        # Add Jira ticket if it exists and is not N/A
        if [ -n "$JIRA_TICKET" ] && [ "$JIRA_TICKET" != "N/A" ]; then
            ENTRY_FORMAT="$ENTRY_FORMAT $JIRA_TICKET"
        fi

        # Add to appropriate section based on change type
        case "$CHANGE_TYPE" in
            breaking)
                BREAKING_CHANGES+="$ENTRY_FORMAT\n"
                # GitHub format
                if [ -n "$PR_NUMBER" ]; then
                    echo "[\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DESCRIPTION [#$PR_NUMBER](https://github.com/${REPO_OWNER}/${REPO_NAME}/pull/$PR_NUMBER) $JIRA_TICKET" >> "$TEMP_GITHUB_FILE"
                else
                    echo "[\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DESCRIPTION $JIRA_TICKET" >> "$TEMP_GITHUB_FILE"
                fi
                ;;
            feature)
                FEATURES+="$ENTRY_FORMAT\n"
                if [ -n "$PR_NUMBER" ]; then
                    echo "[\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DESCRIPTION [#$PR_NUMBER](https://github.com/${REPO_OWNER}/${REPO_NAME}/pull/$PR_NUMBER) $JIRA_TICKET" >> "$TEMP_GITHUB_FILE"
                else
                    echo "[\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DESCRIPTION $JIRA_TICKET" >> "$TEMP_GITHUB_FILE"
                fi
                ;;
            enhancement)
                ENHANCEMENTS+="$ENTRY_FORMAT\n"
                if [ -n "$PR_NUMBER" ]; then
                    echo "[\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DESCRIPTION [#$PR_NUMBER](https://github.com/${REPO_OWNER}/${REPO_NAME}/pull/$PR_NUMBER) $JIRA_TICKET" >> "$TEMP_GITHUB_FILE"
                else
                    echo "[\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DESCRIPTION $JIRA_TICKET" >> "$TEMP_GITHUB_FILE"
                fi
                ;;
            bugfix)
                BUGS+="$ENTRY_FORMAT\n"
                if [ -n "$PR_NUMBER" ]; then
                    echo "[\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DESCRIPTION [#$PR_NUMBER](https://github.com/${REPO_OWNER}/${REPO_NAME}/pull/$PR_NUMBER) $JIRA_TICKET" >> "$TEMP_GITHUB_FILE"
                else
                    echo "[\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DESCRIPTION $JIRA_TICKET" >> "$TEMP_GITHUB_FILE"
                fi
                ;;
            note)
                NOTES+="$ENTRY_FORMAT\n"
                if [ -n "$PR_NUMBER" ]; then
                    echo "[\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DESCRIPTION [#$PR_NUMBER](https://github.com/${REPO_OWNER}/${REPO_NAME}/pull/$PR_NUMBER) $JIRA_TICKET" >> "$TEMP_GITHUB_FILE"
                else
                    echo "[\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DESCRIPTION $JIRA_TICKET" >> "$TEMP_GITHUB_FILE"
                fi
                ;;
            security)
                SECURITY+="$ENTRY_FORMAT\n"
                if [ -n "$PR_NUMBER" ]; then
                    echo "[\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DESCRIPTION [#$PR_NUMBER](https://github.com/${REPO_OWNER}/${REPO_NAME}/pull/$PR_NUMBER) $JIRA_TICKET" >> "$TEMP_GITHUB_FILE"
                else
                    echo "[\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DESCRIPTION $JIRA_TICKET" >> "$TEMP_GITHUB_FILE"
                fi
                ;;
            deprecation)
                DEPRECATION+="$ENTRY_FORMAT\n"
                if [ -n "$PR_NUMBER" ]; then
                    echo "[\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DESCRIPTION [#$PR_NUMBER](https://github.com/${REPO_OWNER}/${REPO_NAME}/pull/$PR_NUMBER) $JIRA_TICKET" >> "$TEMP_GITHUB_FILE"
                else
                    echo "[\`$COMMIT_HASH\`](https://github.com/${REPO_OWNER}/${REPO_NAME}/commit/$COMMIT_HASH) $DESCRIPTION $JIRA_TICKET" >> "$TEMP_GITHUB_FILE"
                fi
                ;;
            *)
                echo "Warning: Unknown change type '$CHANGE_TYPE' in $FILE"
                ;;
        esac
    fi
done

# Add sections to the AsciiDoc release notes if they have content
if [ -n "$BREAKING_CHANGES" ]; then
    {
        echo "== BREAKING CHANGES"
        echo -e "$BREAKING_CHANGES"
        echo ""
    } >> "$TEMP_ADOC_FILE"
fi

if [ -n "$SECURITY" ]; then
    {
        echo "== SECURITY"
        echo -e "$SECURITY"
        echo ""
    } >> "$TEMP_ADOC_FILE"
fi

if [ -n "$FEATURES" ]; then
    {
        echo "== FEATURES"
        echo -e "$FEATURES"
        echo ""
    } >> "$TEMP_ADOC_FILE"
fi

if [ -n "$ENHANCEMENTS" ]; then
    {
        echo "== ENHANCEMENTS"
        echo -e "$ENHANCEMENTS"
        echo ""
    } >> "$TEMP_ADOC_FILE"
fi

if [ -n "$BUGS" ]; then
    {
        echo "== BUG FIXES"
        echo -e "$BUGS"
        echo ""
    } >> "$TEMP_ADOC_FILE"
fi

if [ -n "$DEPRECATION" ]; then
    {
        echo "== DEPRECATIONS"
        echo -e "$DEPRECATION"
        echo ""
    } >> "$TEMP_ADOC_FILE"
fi

if [ -n "$NOTES" ]; then
    {
        echo "== NOTES"
        echo -e "$NOTES"
        echo ""
    } >> "$TEMP_ADOC_FILE"
fi

# Write the final release notes
mv "$TEMP_GITHUB_FILE" "$GITHUB_OUTPUT_FILE"
mv "$TEMP_ADOC_FILE" "$ASCIIDOC_OUTPUT_FILE"

echo "GitHub release notes generated in $GITHUB_OUTPUT_FILE"
echo "AsciiDoc release notes generated in $ASCIIDOC_OUTPUT_FILE"

# Generate the formatted release notes files
echo "Generating formatted release notes files..."
"$(dirname "$0")/format-release-notes.sh" "$VERSION"
