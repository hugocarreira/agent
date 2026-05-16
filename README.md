# 🤖 Centralized AI Agent Configuration

**Single source of truth for OpenCode, Codex, and Claude Code**

This repository centralizes ALL agent configurations, skills, and scripts across multiple AI coding harnesses.

---

## 📁 Structure

```
~/work/agent/
├── AGENTS-GLOBAL.md          # Orchestrator-based rules (for OpenCode)
├── RTK.md                     # Token optimization CLI rules
├── AGENTS.template.md         # Template for project-specific configs
├── CHANGELOG.md               # History of changes
│
├── agents/                    # Specialized agent definitions
│   ├── orchestrator.md
│   ├── frontend-developer.md
│   ├── backend-developer.md
│   ├── designer-ui-ux.md
│   ├── devops-engineer.md
│   ├── product-manager.md
│   ├── marketing-growth.md
│   ├── code-reviewer.md
│   ├── security-auditor.md
│   └── technical-writer.md
│
├── skills/                    # Reusable workflow skills (32 skills)
│   ├── accessibility/
│   ├── agent-browser/
│   ├── best-practices/
│   ├── commit-and-push/
│   ├── frontend-design/
│   ├── pullrequest/
│   ├── security-best-practices/
│   ├── test-and-verify/
│   └── ... (28 more)
│
├── plugins/                   # Harness plugins
│   └── rtk.ts                # RTK token optimizer plugin
│
├── docs/                      # Documentation
│   ├── shipping-checklist.md
│   └── refactoring-guide.md
│
├── scripts/                   # Helper scripts
│   └── (portable helpers)
│
├── stacks/                    # Stack-specific configs
│   ├── opencode/
│   │   └── AGENTS.md         # → ~/.config/opencode/AGENTS.md
│   ├── codex/
│   │   └── AGENTS.md         # → ~/.codex/AGENTS.md
│   └── claude/
│       └── CLAUDE.md         # → ~/.claude/CLAUDE.md
│
└── setup-symlinks.sh          # Automated symlink setup
```

---

## 🚀 Quick Setup

### 1. Run Automated Setup
```bash
cd ~/work/agent
./setup-symlinks.sh
```

This will:
- Backup existing configs
- Create symlinks from all stacks to `~/work/agent`
- Centralize skills, agents, and configs

### 2. Verify Setup
```bash
# Check symlinks
ls -la ~/.config/opencode/AGENTS.md
ls -la ~/.codex/AGENTS.md
ls -la ~/.claude/CLAUDE.md
ls -la ~/.claude/skills
```

### 3. Test It
```bash
# From any AI coding harness:
# OpenCode, Codex, or Claude Code should now load centralized configs
```

---

## 🎯 How It Works

### Centralized Configuration
All three stacks now read from the same source:

```
OpenCode reads:
  ~/.config/opencode/AGENTS.md → ~/work/agent/stacks/opencode/AGENTS.md
                                    ↓ (reads global)
                               ~/work/agent/AGENTS-GLOBAL.md

Codex reads:
  ~/.codex/AGENTS.md → ~/work/agent/stacks/codex/AGENTS.md
                         ↓ (includes)
                      ~/work/agent/RTK.md

Claude reads:
  ~/.claude/CLAUDE.md → ~/work/agent/stacks/claude/CLAUDE.md
                          ↓ (includes)
                       ~/work/agent/RTK.md

All stacks share:
  ~/.{opencode,codex,claude}/skills → ~/work/agent/skills/
```

### Stack-Specific Behavior

**OpenCode** (Orchestrator-first)
- Uses specialized agents for domain work
- Follows AGENTS-GLOBAL.md orchestrator rules
- Delegates to `frontend-developer`, `backend-developer`, etc.

**Codex** (Peter Steinberger workflow)
- "Just Talk To It" philosophy
- Short prompts + screenshots
- Blast radius thinking
- 80% building / 20% refactoring

**Claude Code** (Conversational)
- Natural language prompts
- Iterate fast approach
- Test in same context
- Queue continue messages

### Shared Skills
All 32 skills are available to all stacks:

**Development**
- `commit-and-push`, `pullrequest`, `test-and-verify`
- `frontend-design`, `best-practices`, `security-best-practices`

**AI/Browser**
- `agent-browser`, `simplify`

**Next.js**
- `next-best-practices`, `next-cache-components`, `next-upgrade`

**SEO/Marketing**
- `seo`, `seo-audit`, `programmatic-seo`
- `marketing-ideas`, `marketing-psychology`

**CRO/Product**
- `page-cro`, `signup-flow-cro`, `onboarding-cro`
- `product-strategy`, `product-marketing-context`

**Design/Performance**
- `premium-saas-design`, `web-design-guidelines`
- `accessibility`, `core-web-vitals`, `performance`

---

## 📝 Usage Patterns

### In Projects
Create project-specific `AGENTS.md`:

```markdown
READ ~/work/agent/AGENTS-GLOBAL.md BEFORE ANYTHING (skip if missing).

# My Project - Specific Rules

## Stack
- TypeScript + Next.js 15
- PostgreSQL + Prisma
- Tailwind CSS

## Conventions
- Use `pnpm` not `npm`
- API routes in `app/api/`
- Components in `components/`

## Out of Scope
- Never modify `lib/billing/` (legacy system)
```

### Discovering New Patterns
When you find a useful pattern:

```bash
cd ~/work/agent

# Add to global rules
nano AGENTS-GLOBAL.md

# Or create a new skill
mkdir skills/my-new-skill
nano skills/my-new-skill/SKILL.md

# Commit
git add .
git commit -m "feat: add new pattern X"
```

### Cross-Stack Updates
Edit once, applies everywhere:

```bash
# Update RTK config
nano ~/work/agent/RTK.md
# ✅ Both Codex and Claude see the change immediately

# Add new skill
cp -r my-skill ~/work/agent/skills/
# ✅ OpenCode, Codex, and Claude all discover it
```

---

## 🔧 Maintenance

### Update Global Rules
```bash
cd ~/work/agent
nano AGENTS-GLOBAL.md
git commit -am "refactor: update orchestrator rules"
```

### Add New Skill
```bash
cd ~/work/agent/skills
mkdir my-new-workflow
cat > my-new-workflow/SKILL.md <<'EOF'
---
name: "my-new-workflow"
description: "Short trigger phrase"
---

# My New Workflow

[skill content]
EOF
```

### Backup Before Changes
```bash
# The setup script auto-backs up on each run
./setup-symlinks.sh
# Creates .backup.{timestamp} files
```

### Verify Symlinks
```bash
# Check all symlinks are intact
ls -la ~/.config/opencode/{AGENTS.md,skills,agents,plugins}
ls -la ~/.codex/{AGENTS.md,RTK.md,skills}
ls -la ~/.claude/{CLAUDE.md,AGENTS.md,RTK.md,skills}
```

---

## 🎓 Philosophy

### Peter Steinberger Principles
- **Just Talk To It**: Short prompts, natural conversation
- **Shipping at Inference-Speed**: Code works out-of-the-box
- **Blast Radius Thinking**: Know impact before starting
- **Iterate Fast, Refactor Smart**: 80/20 rule
- **CLIs > MCPs**: Native tools over context pollution
- **Test in Same Context**: Find bugs immediately

### Orchestrator Architecture (OpenCode)
- **Single Entry Point**: Orchestrator handles all requests
- **Agents Stay in Lane**: No domain overlap
- **Structured Reporting**: Clear handoffs
- **Quality Gates**: Code review before shipping
- **Contracts Are Law**: Shared interfaces defined upfront

### Universal Rules
- Token efficiency (RTK everywhere)
- Conventional Commits
- Tests for business logic
- No secrets in code
- Security by default

---

## 📊 Stats

- **3 AI Harnesses**: OpenCode (primary), Codex, Claude Code
- **32 Skills**: Accessible to all stacks
- **10 Specialized Agents**: For orchestrated workflows
- **1 Source of Truth**: `~/work/agent`

---

## 🚦 Next Steps

1. ✅ Run `./setup-symlinks.sh`
2. ✅ Verify symlinks are created
3. ✅ Test with your preferred harness
4. 📝 Customize stack-specific configs if needed
5. 🔄 Commit changes to git

---

**Maintainer**: This structure is inspired by [Peter Steinberger's agent-scripts](https://github.com/steipete/agent-scripts) and adapted for multi-stack workflows.
