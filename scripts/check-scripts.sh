#!/usr/bin/env bash
#
# Run shellcheck on all shell scripts in the repository
#
# Usage: ./check-scripts.sh [--fix]
#

set -eo pipefail

# Find all shell scripts in the repository (files with .sh extension and files with shebang line)
find_scripts() {
  # Find files with .sh extension
  find . -name "*.sh" -type f | sort

  # Find files with shebang line starting with #!/bin/bash, #!/usr/bin/env bash, or #!/bin/sh
  find . -type f ! -path "*/\.*" ! -path "*/vendor/*" ! -path "*/node_modules/*" -perm +111 -exec grep -l '^\#\!/bin/bash\|^\#\!/usr/bin/env bash\|^\#\!/bin/sh' {} \; 2>/dev/null | sort -u || true
}

# Check if shellcheck is installed
if ! command -v shellcheck &>/dev/null; then
  echo "shellcheck not found. Please install shellcheck:"
  echo "  - macOS: brew install shellcheck"
  echo "  - Ubuntu: apt-get install shellcheck"
  echo "  - Visit https://github.com/koalaman/shellcheck#installing for more options"
  exit 1
fi

# Check if shfmt is installed
if ! command -v shfmt &>/dev/null; then
  echo "shfmt not found. Please install shfmt:"
  echo "  - macOS: brew install shfmt"
  echo "  - go install mvdan.cc/sh/v3/cmd/shfmt@latest"
  echo "  - Visit https://github.com/mvdan/sh#shfmt for more options"
  exit 1
fi

FIX=false
if [[ ${1:-} == "--fix" ]]; then
  FIX=true
fi

ERRORS=0
SCRIPTS=$(find_scripts)

# Run shellcheck on all scripts
echo "Running shellcheck on shell scripts..."
for script in $SCRIPTS; do
  if ! shellcheck -x "$script"; then
    ERRORS=$((ERRORS + 1))
    echo "shellcheck failed for $script"
  else
    echo "shellcheck passed for $script"
  fi
done

# Run shfmt on all scripts
echo "Running shfmt on shell scripts..."
for script in $SCRIPTS; do
  if $FIX; then
    echo "Formatting $script with shfmt"
    shfmt -i 2 -ci -bn -s -w "$script"
  else
    if ! shfmt -i 2 -ci -bn -s -d "$script"; then
      ERRORS=$((ERRORS + 1))
      echo "shfmt check failed for $script"
    else
      echo "shfmt check passed for $script"
    fi
  fi
done

if [ "$ERRORS" -gt 0 ]; then
  echo "Found $ERRORS error(s) in shell scripts"
  if ! $FIX; then
    echo "Run './scripts/check-scripts.sh --fix' to automatically fix formatting issues"
  fi
  exit 1
else
  echo "All shell scripts passed checks!"
fi
