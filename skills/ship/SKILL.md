---
name: ship
description: Complete shipping workflow - branch, commit by context, push, and create PR in one command
summary: Ship changes - auto branch (if main), commit files by context, push, open PR
read_when: User says "ship", "ship this", "ship it"
triggers:
  - ship
  - ship this
  - ship it
---

# Ship

Complete shipping workflow in one command.

## What It Does

1. **Branch** - If on main, create feature branch (feat/fix/refactor/docs)
2. **Commit** - Group files by context, commit separately with Conventional Commits
3. **Push** - Push to remote (set upstream if needed)
4. **PR** - Create pull request with `gh pr create`

## Smart File Grouping

**Group together:**
- Component + test file
- Related functionality

**Commit separately:**
- Config files (package.json, tsconfig.json)
- Documentation (README.md, CHANGELOG.md)
- Dependencies
- Unrelated features

## Commit Message Format

Use Conventional Commits:
```
feat: add user authentication
fix: memory leak in WebSocket handler
refactor: simplify API error handling
docs: update README with setup instructions
chore: update dependencies
```

## Branch Naming

```
feat/description    # New features
fix/description     # Bug fixes
refactor/description # Refactoring
docs/description    # Documentation
```

## Workflow

```bash
# 1. Check branch
current_branch=$(git branch --show-current)
if [ "$current_branch" = "main" ]; then
  # Create feature branch (ask user or auto-generate)
  git checkout -b feat/new-feature
fi

# 2. Group and commit files
# Analyze git status, group by context
# For each group:
#   - Show files + generated message
#   - Ask confirmation
#   - git add <files> && git commit -m "message"

# 3. Push
git push -u origin $(git branch --show-current)

# 4. Create PR
gh pr create --title "..." --body "..." --base main
```

## Options

**Simple:**
```
"ship"
"ship this"
```

**Custom branch:**
```
"ship on branch feat/new-feature"
```

**Skip PR:**
```
"ship without PR"
```

## vs. Other Skills

**Use `ship`** - Complete automation (branch + commits + push + PR)  
**Use `commit-and-push`** - Manual commit messages, no PR  
**Use `pullrequest`** - Already committed/pushed, just need PR

## Example

```
User modifies:
  src/Button.tsx
  src/Button.test.tsx
  package.json
  README.md

User: "ship"

Agent:
  ✓ Create feat/add-button (was on main)
  ✓ Commit #1: Button.tsx + test
    "feat: add Button component with tests"
  ✓ Commit #2: package.json
    "chore: add button dependencies"
  ✓ Commit #3: README.md
    "docs: document Button component"
  ✓ Push feat/add-button
  ✓ PR #123 created
  
  https://github.com/user/repo/pull/123
```

---

**One command. Full workflow. Done.**
