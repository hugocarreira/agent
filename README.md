# AI Agent Setup - Centralized

**Single source of truth for all AI coding assistants (opencode, codex, claude)**

Philosophy: Just talk to it. Give agent way to verify. Ship fast.

---

## Quick Setup

```bash
cd ~/work/agent
./scripts/setup-symlinks.sh
```

Creates symlinks:
```
~/.config/opencode/  → ~/work/agent/stacks/opencode/
~/.codex/            → ~/work/agent/stacks/codex/
~/.claude/           → ~/work/agent/stacks/claude/
skills/              → ~/work/agent/skills/ (all stacks)
```

---

## Structure

```
~/work/agent/
├── AGENTS.md              # Core rules (start here!)
├── QUICK-REFERENCE.md     # Daily cheatsheet
├── RTK.md                 # Token optimization
│
├── skills/                # 32 reusable workflows
├── stacks/                # Stack-specific configs
├── research/              # Deep research (optional reading)
└── setup-symlinks.sh
```

---

## Usage

### Daily Workflow
```bash
# Agent reads:
# 1. ~/work/agent/AGENTS.md (core rules)
# 2. ~/work/agent/RTK.md (token optimization)  
# 3. Project's AGENTS.md (if exists)
# 4. Your code

# Just talk:
"add login button"
"[screenshot] fix padding"
"look at ../other-project and do same"
```

### Add to Projects

```bash
cd ~/work/agent
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

**32 Skills** (auto-trigger):
- Development: commit-and-push, test-and-verify, frontend-design
- Next.js: next-best-practices, next-upgrade
- SEO: seo, seo-audit, programmatic-seo
- Product: page-cro, product-strategy
- + 23 more

**3 Harnesses Unified**:
- OpenCode
- Codex
- Claude

**All share**: RTK, skills, core principles, global AGENTS.md

---

## Stats

```
✅ 32 skills (all stacks)
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

**Optional:**
- research/ - Deep research on AI dev practices

**Keep it simple. Ship fast. Stay in control.**
