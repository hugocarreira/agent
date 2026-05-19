# AI Agent Setup - Centralized

**Single source of truth for all AI coding assistants (opencode, codex, claude)**

Philosophy: Just talk to it. Give agent way to verify. Ship fast.

---

## Quick Setup

```bash
./scripts/setup-symlinks.sh
```

Creates symlinks:
```
~/.config/opencode/  → <repo>/stacks/opencode/ (if dir exists)
~/.codex/            → <repo>/stacks/codex/     (if dir exists)
~/.claude/           → <repo>/stacks/claude/    (if dir exists)
skills/              → <repo>/skills/ (all stacks)
```

---

## Structure

```
<repo>/
├── AGENTS.md              # Core rules (start here!)
├── QUICK-REFERENCE.md     # Daily cheatsheet
├── RTK.md                 # Token optimization
├── SETUP.md               # Complete setup guide
├── CHANGELOG.md           # Version history
│
├── skills/                # 26 reusable workflows
│   ├── _references/       # Knowledge base docs
│   └── README.md          # Skills catalog
├── scripts/               # Automation scripts
│   ├── setup-symlinks.sh  # Global setup
│   └── setup-project.sh   # Per-project setup
├── stacks/                # Stack-specific configs
│   ├── opencode/
│   ├── codex/
│   └── claude/
├── research/              # Deep research (optional)
└── docs/                  # Guides and checklists
```

---

## Usage

### Daily Workflow
```bash
# Agent reads:
# 1. AGENTS.md (core rules, via symlink)
# 2. RTK.md (token optimization, via @include)
# 3. Project's AGENTS.md (if exists)
# 4. Your code

# Just talk:
"add login button"
"[screenshot] fix padding"
"look at ../other-project and do same"
```

### Add to Projects

```bash
./scripts/setup-project.sh
# Interactive: asks project name + option (new/existing)
```

---

## Core Principles

**#1 Give Agent Way to Verify (HIGHEST LEVERAGE!)**
- Tests agent can run
- Screenshots to compare
- Benchmarks with clear criteria

**#2 Context is Precious**
- /clear between tasks
- /clear after 2 corrections

**#3 Just Talk To It**
- Short prompts + screenshots
- Natural conversation

**#4 Final Manual Review**
- Always understand what agent did
- Never ship without review

---

## What's Included

**26 Skills** (auto-trigger):
- Dev Workflow: ship, commit-and-push, test-and-verify, pullrequest
- Frontend: frontend-design, accessibility, best-practices
- Next.js: next-best-practices, next-upgrade
- SEO: seo, programmatic-seo, competitor-alternatives
- + 16 more (see skills/README.md)

**3 Harnesses Unified**:
- OpenCode
- Codex
- Claude

**All share**: RTK, skills, core principles, global AGENTS.md

---

## Stats

```
✅ 26 skills (all stacks)
✅ 5 reference docs (knowledge base)
✅ 3 harnesses unified  
✅ 1 source of truth
✅ 0 orchestrators
✅ 0 theater
```

---

## Files to Read

**Essential:**
1. AGENTS.md - Core rules
2. QUICK-REFERENCE.md - Daily cheatsheet
3. SETUP.md - Complete setup guide

**For Reference:**
- skills/README.md - Complete skills catalog
- CHANGELOG.md - Version history
- research/ - Deep research on AI dev practices

**Keep it simple. Ship fast. Stay in control.**
