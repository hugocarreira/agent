---
name: changelog
summary: Create or update CHANGELOG.md following Keep a Changelog format
read_when: User asks to create changelog, update changelog, generate release notes, or prepare version release
triggers:
  - changelog
  - update changelog
  - create changelog
  - generate changelog
  - release notes
  - prepare release
  - version bump
---

# Changelog Generator

Create or update `CHANGELOG.md` following [Keep a Changelog](https://keepachangelog.com/en/1.0.0/) format with automatic commit analysis.

## Purpose

Automate changelog creation by analyzing git commits since last release and categorizing changes into standard sections.

## Workflow

### 1. Analyze Commits

**Without Git Tags (default):**

```bash
# Since last entry in CHANGELOG (by date)
# Extract date from last changelog entry, get commits since then
last_date=$(grep -m1 "## \[.*\] - " CHANGELOG.md | grep -oP '\d{4}-\d{2}-\d{2}')
git log --since="$last_date" --oneline

# Since specific date
git log --since="2025-01-01" --oneline

# Last N commits
git log -n 50 --oneline

# Since specific commit hash
git log abc123..HEAD --oneline

# All commits (for initial changelog)
git log --oneline
```

**With Git Tags (optional):**

```bash
# If you use tags, get commits since last tag
git log $(git describe --tags --abbrev=0)..HEAD --oneline
```

### 2. Categorize by Conventional Commits

Use commit message prefixes to categorize:

| Prefix | Category | Description |
|--------|----------|-------------|
| `feat:` | **Added** | New features |
| `fix:` | **Fixed** | Bug fixes |
| `refactor:` | **Changed** | Code refactoring |
| `docs:` | **Changed** | Documentation updates |
| `chore:` | **Changed** | Maintenance tasks |
| `perf:` | **Changed** | Performance improvements |
| `test:` | **Changed** | Test updates |
| `style:` | **Changed** | Code style changes |
| `ci:` | **Changed** | CI/CD changes |
| `build:` | **Changed** | Build system changes |
| `revert:` | **Fixed** | Reverted changes |
| **BREAKING CHANGE** | **Changed** | Breaking changes (highlight!) |

### 3. Format Structure

```markdown
# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

---

## [Unreleased]

### Added
- New feature X
- New feature Y

### Changed
- Refactored Z for better performance
- Updated documentation for A

### Fixed
- Bug in B that caused C
- Issue with D

### Removed
- Deprecated feature E

## 2025-01-30

### Added
- Feature set for this release

### Changed
- Improvements and refactoring

## 2025-01-29

### Added
- Initial release
```

### 4. Update Workflow

**Default (no git tags):**
1. Read existing CHANGELOG.md (if exists)
2. Ask user for scope:
   - "Last N commits" (e.g., last 20)
   - "Since date" (e.g., since 2025-01-01)
   - "Since last CHANGELOG entry" (extract date, get commits after)
   - "All commits" (for initial changelog)
   - "Since commit hash" (e.g., since abc123)
3. Categorize commits by prefix
4. Insert new date section (YYYY-MM-DD)
5. Preserve all previous entries

**With git tags (optional):**
1. Detect if project uses tags (`git tag -l`)
2. Get commits since last tag
3. Use semantic version format `[1.0.0]` instead of date
4. Include comparison links

**For unreleased changes:**
1. Add to `[Unreleased]` section
2. Update regularly during development
3. Move to dated section when releasing

## Best Practices

### Group Related Changes
```markdown
### Added
- **Authentication**: OAuth2 provider, JWT tokens, refresh flow
- **API**: New endpoints for user management, bulk operations
- **UI**: Dashboard redesign, mobile responsive nav
```

### Be Specific
```markdown
❌ "Fixed bugs"
✅ "Fixed memory leak in WebSocket connection handler"

❌ "Updated dependencies"
✅ "Updated React 18 → 19 for improved concurrent rendering"
```

### Highlight Breaking Changes
```markdown
### Changed
- **BREAKING**: Authentication now requires API key in header instead of query param
- **BREAKING**: Removed deprecated `/v1/users` endpoint (use `/v2/users`)
```

### Link to Issues/PRs
```markdown
### Fixed
- Memory leak in connection pool (#123)
- Race condition in cache invalidation (#456)
```

## Special Cases

### Initial Release
Keep it simple:
```markdown
## 2025-01-29

### Added
- Initial release with core features
- User authentication
- API endpoints
- Documentation
```

### Major Milestone (optional version marker)
```markdown
## 2025-02-15 (v1.0.0)

### Added
- Production-ready release
- Full feature set complete
```

### Hotfixes
```markdown
## 2025-01-30

### Fixed
- Critical security vulnerability in auth
- Production crash in payment processing
```

### Using Versions (if you start using git tags)
```markdown
## [1.0.0] - 2025-01-29

### Added
- Initial release
```

## Automation Tips

### Extract from Commits
```bash
# Get all 'feat:' commits
git log --oneline --grep="^feat:" v1.0.0..HEAD

# Get all 'fix:' commits  
git log --oneline --grep="^fix:" v1.0.0..HEAD

# Check for breaking changes
git log --grep="BREAKING CHANGE" v1.0.0..HEAD
```

### Auto-categorize
1. Parse each commit message
2. Extract prefix (`feat`, `fix`, etc.)
3. Map to changelog category
4. Group by category
5. Format as bullet points

### Date Format
Use ISO date: `YYYY-MM-DD` (e.g., `2025-01-29`)

## Output

### Ask Before Writing
1. **Ask for scope**: "How far back? (last N commits / since date / since last entry / all)"
2. Show categorized changes to user
3. Ask for confirmation
4. Request additional context if needed
5. Let user edit/reorder before writing

### Preserve History
- Never delete previous versions
- Always append new sections at top (below Unreleased)
- Keep comparison links updated

### Commit Message
```bash
git commit -m "docs: update CHANGELOG.md for v1.1.0"
```

## Example Invocations

**Without git tags (default):**
```
"update changelog with last 30 commits"
"create changelog since January 1st"
"generate changelog since last entry"
"update changelog with all new commits"
"add changelog entry for today's work"
```

**With git tags (if you use them):**
```
"update changelog for v1.1.0"
"create changelog from commits since v1.0.0"
```

## Related

- Git log and commit history analysis
- Conventional Commits spec
- Semantic versioning
- Release automation

---

**Remember**: Changelog is for humans, not machines. Focus on what changed and why it matters to users.
