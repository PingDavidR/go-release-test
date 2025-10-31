# Changelog & Release Notes System Bundle

**Version:** 2.0 (Updated October 2025)  
**Compatible with:** Any Git-based project

## What's Included

This bundle provides a complete changelog and release notes management system that can be integrated into any repository. It includes:

### 📁 Core Scripts (7 files)
- `create-changelog.sh` - Create changelog entries for PRs
- `generate-release-notes.sh` - Generate release notes from changelog entries
- `archive-changelog.sh` - Archive processed changelog files
- `format-release-notes.sh` - Format release notes output
- `check-scripts.sh` - Validate script syntax
- `pre-commit-noagent-check.sh` - Pre-commit hook for agent protection
- `run-integration-tests.sh` - Integration test runner

### 📚 Documentation (8 files)
- `README.md` - Complete changelog system documentation
- `FORMAT_EXAMPLES.md` - Examples for all changelog types
- `QUICK_START.md` - Get started in 5 minutes
- `INSTALLATION.md` - Step-by-step installation guide
- `GITHUB_ACTIONS_SNIPPETS.md` - CI/CD integration examples
- `README_SNIPPETS.md` - Ready-to-paste README sections
- `QUICK_REFERENCE.md` - Command cheat sheet
- `AGENT-PROTECTION.md` - AI agent safety documentation

## Key Features

✅ **8 Change Types:** breaking, feature, enhancement, bugfix, note, security, deprecation, internal  
✅ **Internal Tracking:** Track internal changes without publishing them  
✅ **Jira Integration:** Automatic ticket validation (CDI-###/PDI-### format)  
✅ **Multiple Formats:** Markdown and AsciiDoc output  
✅ **GitHub Actions:** Ready-to-use workflow snippets  
✅ **No Length Limits:** Describe changes as thoroughly as needed  
✅ **Optional Changelog:** Won't block PRs if changelog is missing  

## Quick Start

1. **Extract the bundle:**
   ```bash
   unzip changelog-bundle.zip -d /tmp/changelog-system
   cd /tmp/changelog-system
   ```

2. **Copy scripts to your project:**
   ```bash
   cp -r scripts/* YOUR_PROJECT/scripts/
   chmod +x YOUR_PROJECT/scripts/*.sh
   ```

3. **Copy documentation:**
   ```bash
   cp -r .changelog YOUR_PROJECT/
   ```

4. **Create your first changelog entry:**
   ```bash
   cd YOUR_PROJECT
   ./scripts/create-changelog.sh 123 feature "Added amazing new feature"
   ```

5. **Add GitHub Actions integration** (see `GITHUB_ACTIONS_SNIPPETS.md`)

## Integration Points

### For README.md
See `README_SNIPPETS.md` for ready-to-paste sections including:
- Changelog process overview
- Developer workflow
- Change type descriptions
- Example commands

### For GitHub Actions
See `GITHUB_ACTIONS_SNIPPETS.md` for:
- PR changelog validation (optional)
- Format validation workflow
- Release notes generation

### For Contributing Guidelines
See `QUICK_START.md` for contributor-focused instructions

## Change Type Reference

| Type | Purpose | Published in Release Notes |
|------|---------|---------------------------|
| `breaking` | Breaking changes requiring migration | ✅ Yes |
| `feature` | New features and capabilities | ✅ Yes |
| `enhancement` | Improvements to existing features | ✅ Yes |
| `bugfix` | Bug fixes | ✅ Yes |
| `note` | Important notices or announcements | ✅ Yes |
| `security` | Security-related changes | ✅ Yes |
| `deprecation` | Deprecation notices | ✅ Yes |
| `internal` | Internal changes (refactoring, tests) | ❌ No |

## Changelog Entry Format

```
type: <change-type>
description: <detailed description of the change>
pr: <pull-request-number>
ticket: <JIRA-ticket-or-N/A>
```

### Example
```
type: feature
description: Added multi-factor authentication support with TOTP and SMS backup options
pr: 456
ticket: CDI-789
```

## Customization Options

### Change Jira Ticket Format
Edit `create-changelog.sh` line ~51:
```bash
if ! [[ $JIRA_TICKET =~ ^(YOUR|PREFIX)-[0-9]+$ ]]; then
```

### Add Custom Change Types
Edit both files:
- `create-changelog.sh` line ~76: Add to `VALID_TYPES` array
- `generate-release-notes.sh` line ~109: Add to type grouping logic

### Modify Output Formats
Edit `generate-release-notes.sh`:
- Markdown generation: ~lines 150-250
- AsciiDoc generation: ~lines 260-350

## Files Location After Installation

```
YOUR_PROJECT/
├── scripts/
│   ├── create-changelog.sh
│   ├── generate-release-notes.sh
│   ├── archive-changelog.sh
│   ├── format-release-notes.sh
│   ├── check-scripts.sh
│   ├── pre-commit-noagent-check.sh
│   └── run-integration-tests.sh
├── .changelog/
│   ├── README.md
│   ├── FORMAT_EXAMPLES.md
│   └── archive/  (created when first release is archived)
└── .github/
    └── workflows/
        └── verify-changelog.yml  (optional)
```

## Support & Troubleshooting

### Common Issues

**Q: Scripts won't execute**  
A: Run `chmod +x scripts/*.sh` to make them executable

**Q: Jira validation failing**  
A: Tickets must match format `CDI-###` or `PDI-###`. Customize in `create-changelog.sh` if needed.

**Q: Want to skip changelog for a PR**  
A: This is now supported! The pipeline won't fail if changelog is missing.

**Q: Need to track internal work**  
A: Use `type: internal` - it will be tracked but excluded from release notes.

## Version History

- **v2.0** (Oct 2025): Removed character limits, made changelog optional, added internal type
- **v1.0** (Oct 2025): Initial release with key-value format

## License

Same license as the parent project (see LICENSE file)

## Credits

Created for flexible changelog management across multiple projects.
