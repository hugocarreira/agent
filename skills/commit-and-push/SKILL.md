---
name: "commit-and-push"
description: "Commit changes and push to remote"
---

# Commit and Push Skill

Atomic commit workflow with changelog updates.

## Trigger Phrases
- "ship"
- "commit and push"  
- "push this"
- "deploy"

## Workflow

### 1. Check Status
```bash
git status
```

### 2. Review Changes
```bash
git diff
```

### 3. Update Changelog (if user-facing)
Check if changes affect users:
- New features
- Bug fixes  
- API changes
- UI improvements

If yes, add to `CHANGELOG.md`:
```markdown
## [Unreleased]

### Added
- Feature X that does Y

### Fixed  
- Bug where Z happened

### Changed
- Updated component A to B
```

### 4. Stage Files
```bash
git add [specific files that changed]
```

Only stage files you modified. Don't stage unrelated changes.

### 5. Commit
```bash
git commit -m "type(scope): message"
```

**Types**:
- `feat`: new feature
- `fix`: bug fix
- `refactor`: code improvement, no behavior change
- `docs`: documentation only
- `style`: formatting, no code change
- `perf`: performance improvement
- `test`: adding/fixing tests
- `build`: build system changes
- `ci`: CI/CD changes
- `chore`: maintenance

**Examples**:
```bash
git commit -m "feat(auth): add password reset flow"
git commit -m "fix(api): handle null response in user endpoint"  
git commit -m "refactor(components): extract Button to shared"
```

### 6. Push
```bash
git push
```

If branch doesn't exist on remote:
```bash
git push -u origin <branch-name>
```

### 7. Verify
```bash
git log -1 --oneline
git status
```

Confirm push succeeded and working tree is clean.

## Examples

### Simple Feature
```bash
# User: "commit and push the new login button"
git status
git add src/components/LoginButton.tsx
git commit -m "feat(ui): add new login button component"
git push
```

### Bug Fix with Changelog
```bash
# User: "ship this fix"
# 1. Update CHANGELOG.md
git add src/api/users.ts CHANGELOG.md
git commit -m "fix(api): prevent null pointer in getUserById"
git push
```

### Multiple Related Changes
```bash
# Commit each logical unit separately
git add src/components/Header.tsx
git commit -m "feat(ui): add responsive header"

git add src/styles/global.css  
git commit -m "style(ui): update header breakpoints"

git push
```

## Safety Rules

- Never commit secrets, API keys, `.env` files
- Never force push to main/master without explicit permission
- Never commit `node_modules`, `dist`, build artifacts
- Never stage files you didn't modify (except CHANGELOG)
- Always verify git status after push

## Troubleshooting

### Merge Conflicts
```bash
git pull --rebase
# Fix conflicts
git add .
git rebase --continue  
git push
```

### Wrong Commit Message
```bash
git commit --amend -m "corrected message"
git push --force-with-lease  # Only if not pushed yet
```

### Forgot to Update Changelog
```bash
# Edit CHANGELOG.md
git add CHANGELOG.md
git commit --amend --no-edit
git push --force-with-lease  # Only if just pushed
```

---

Keep commits atomic, messages clear, and always verify the push succeeded.
