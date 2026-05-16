# Setup Guide

Centralized AI agent environment for opencode, codex, and claude.

**Philosophy:** Just talk to it. Give agent way to verify. Ship fast.

---

## Quick Start

### 1. Global Setup (Once)

```bash
cd ~/work/agent
./scripts/setup-symlinks.sh
```

Creates symlinks:
```
~/.config/opencode/ → ~/work/agent/stacks/opencode/
~/.codex/           → ~/work/agent/stacks/codex/
~/.claude/          → ~/work/agent/stacks/claude/
skills/             → ~/work/agent/skills/ (all stacks)
```

**Verify:**
```bash
ls -la ~/.claude/CLAUDE.md      # or ~/.config/opencode/ or ~/.codex/
ls -la ~/.claude/skills/
# Should point to ~/work/agent/...
```

---

### 2. Per-Project Setup

**Interactive (recommended):**
```bash
cd ~/work/agent
./scripts/setup-project.sh

# Prompts:
# - Project name
# - Option 1: New project (create AGENTS.md)
# - Option 2: Existing project (add global reference)
```

**Or use individual scripts:**
```bash
# New project
./scripts/setup-new-project.sh <project-name>

# Existing project
./scripts/setup-existing-project.sh <project-name>
```

**Note:** All scripts assume your projects are in `~/work/`

---

## How It Works

When you start your AI coding session in any project:

```
1. Agent reads its harness config (opencode/codex/claude)
   ↓ (symlinked to ~/work/agent/stacks/<harness>/)
   ↓ (which references: "READ ~/work/agent/AGENTS.md")

2. Reads: ~/work/agent/AGENTS.md (global rules)
   ↓
3. Reads: ~/work/agent/RTK.md (token optimization)
   ↓
4. Reads: ./AGENTS.md (project-specific rules)
   ↓
5. Reads: your code
   ↓
6. Implements following ALL rules
```

**Priority:** Project > Global > Defaults

---

## What to Put in Project AGENTS.md

### ✅ Include (agent can't infer)

```markdown
## Stack
- Next.js 16, PostgreSQL, bun (not npm!)

## Commands
rtk bun test
rtk bun run dev

## Out of Scope
- Never touch: lib/billing/ (external team)
- Never modify: .env files

## Project-Specific
- API base: http://localhost:8080/api/v1
- Migrations: make migrate-up before schema changes
- Brand voice: casual PT-BR (see .guidelines/brand.md)
```

### ❌ Exclude (agent infers from code)

```markdown
❌ "Components in src/components/" (obvious)
❌ "Use TypeScript" (sees from imports)
❌ "Never commit secrets" (global covers this)
❌ Full API docs (link is enough)
```

**Goal:** 50-100 lines, not 500.

---

## Daily Usage

```bash
cd ~/work/my-project
# Start your AI agent (opencode/codex/claude)

"add login button"

# Agent understands:
# - Global: "Just talk to it, use RTK, give way to verify"
# - RTK: "Prefix commands with rtk"
# - Project: "Stack: Next.js, Commands: npm run dev"
# - Code: "Button component exists, reuse pattern"

# Agent does:
# 1. Creates LoginButton.tsx (project pattern)
# 2. Writes tests
# 3. Runs: rtk npm test
# 4. Atomic commit
```

---

## Common Commands

### Update Global (affects ALL projects)
```bash
cd ~/work/agent
# Edit AGENTS.md, RTK.md, etc
rtk git commit -m "feat: add new rule"
# Changes reflect in all harnesses (symlinks!)
```

### Update Project (single project)
```bash
cd ~/work/my-project
# Edit ./AGENTS.md
rtk git commit -m "docs: update project rules"
```

### Test if Working
```bash
cd ~/work/my-project
# Start your AI agent

"what's the API base URL?"
# If answers correctly → working!
# If wrong → check symlinks
```

---

## Troubleshooting

### Agent doesn't read project AGENTS.md
```bash
ls -la ./AGENTS.md  # Must exist
```

### Agent ignores global rules
```bash
ls -la ~/.claude/CLAUDE.md  # (or ~/.config/opencode/ or ~/.codex/)
# Must be symlink
cd ~/work/agent
./scripts/setup-symlinks.sh  # Recreate
```

### Skills don't appear
```bash
ls -la ~/.claude/skills/  # (or ~/.config/opencode/skills/ or ~/.codex/skills/)
# Must point to ~/work/agent/skills/
```

---

## Scripts Reference

All scripts in `~/work/agent/scripts/`:

**setup-symlinks.sh**
- Run once to setup global environment
- Creates symlinks for all 3 harnesses
- Usage: `./scripts/setup-symlinks.sh`

**setup-project.sh** (recommended)
- Interactive project setup
- Prompts for project name and option
- Detects existing AGENTS.md automatically
- Usage: `./scripts/setup-project.sh`

**setup-new-project.sh <name>**
- For new projects (no AGENTS.md)
- Copies AGENTS.template.md to ~/work/<name>/
- Usage: `./scripts/setup-new-project.sh my-app`

**setup-existing-project.sh <name>**
- For existing projects (has AGENTS.md)
- Adds global reference line at top
- Backs up original file
- Usage: `./scripts/setup-existing-project.sh my-app`

---

## Summary

**Setup once:**
1. `./scripts/setup-symlinks.sh` (global)
2. Verify symlinks

**Per project:**
1. `./scripts/setup-new-project.sh <name>` OR `./scripts/setup-existing-project.sh <name>`
2. Edit AGENTS.md (pick stack, add commands)
3. Keep short (50-100 lines)

**Daily use:**
1. `cd ~/work/project`
2. Start your AI agent
3. Natural conversation
4. Agent follows all rules

Simple. Direct. Effective.
