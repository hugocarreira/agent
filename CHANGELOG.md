# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

---

## [2.0.0] - 2026-05-22

### Added

**Go CLI (`cli/`)**
- `agentrc setup` - global config + skills setup (replaces `setup-symlinks.sh`)
- `agentrc init <name>` - create AGENTS.md from template (replaces `setup-project.sh` new)
- `agentrc link <name>` - add global reference to existing project (replaces `setup-project.sh` existing)
- `agentrc verify` - check config and symlinks, detect installed agents, report status
- Zero external deps (cobra, pflag only)
- Built with Go 1.25 + cobra

**Infrastructure**
- `cli/Makefile` for build/install
- Go 1.25.9 installed locally (`~/.local/go/bin/go`)
- `~/.local/go/bin` and `~/go/bin` added to PATH

### Changed
- All scripts migrated from bash to Go CLI
- SETUP.md rewritten for `agentrc` CLI commands

### Removed
- `scripts/setup-symlinks.sh`
- `scripts/setup-project.sh`

---

## [1.0.0] - 2025-01-29

### 🎉 Initial Release

Centralized AI agent environment for 3 harnesses (opencode, codex, claude) with unified configuration, skills, and automation.

### Added

**Core System**
- Global `AGENTS.md` with "Just Talk To It" philosophy (209 lines)
- `RTK.md` for token optimization (60-90% savings)
- `SETUP.md` complete setup guide
- `README.md` quick reference
- `QUICK-REFERENCE.md` daily workflow cheatsheet
- `AGENTS.template.md` with 3 stack options

**Automation**
- `setup-symlinks.sh` - global setup (one-time)
- `setup-project.sh` - interactive project setup (new or existing)

**Skills** (24 total)
- Dev Workflow: 5 skills (agent-browser, commit-and-push, pullrequest, test-and-verify, simplify)
- Frontend/Design: 7 skills (frontend-design, accessibility, web-design-guidelines, vercel-react-view-transitions, best-practices, security-best-practices, next-best-practices)
- Next.js: 2 skills (next-cache-components, next-upgrade)
- SEO/Marketing: 6 skills (seo, programmatic-seo, competitor-alternatives, product-marketing-context, landing-page, signup-flow-cro)
- Growth/CRO: 1 skill (onboarding-cro)
- Performance: 2 skills (core-web-vitals, performance)
- Utilities: 1 skill (find-skills)

**Knowledge Base**
- `_references/` directory with 5 documents (38.5K total)
- marketing-psychology-models.md (50+ mental models)
- cro-experiments.md (comprehensive experiment library)
- ai-writing-detection.md (AI writing patterns to avoid)
- monetization-models.md (7 frameworks)
- product-strategy-canvas.md (9-section canvas)

**Documentation**
- Research findings in `research/` (optional deep-dive)
- Shipping checklist and refactoring guide in `docs/`
- Skills catalog in `skills/README.md`

### Changed

**Architecture**
- Unified 3 harnesses into single source of truth
- Stack wrappers reduced to 3 lines each (93% reduction)
- Symlinks connect all configs to `~/work/agentrc/`

**Philosophy**
- Removed orchestrator patterns (multi-agent, subagents, plan mode)
- Adopted Peter Steinberger's "Just Talk To It" workflow
- Single agent, single terminal approach
- Context management as critical practice

**Skills Cleanup**
- Reduced from 32 to 24 skills (25% reduction)
- Removed 8 bloat skills (2,107 lines)
- Extracted valuable content to 5 reference docs
- Merged best practices into 4 enhanced skills
- Net removal: 852 lines of fluff

### Removed

**Over-Engineering**
- Orchestrator system and agent definitions
- Multi-terminal workflows
- Architect/Implementer pattern
- Parallel agent patterns
- Complex delegation flows

**Bloat Skills**
- premium-saas-design (merged into frontend-design)
- marketing-psychology (converted to reference)
- seo-audit (merged into seo)
- page-cro (merged into landing-page)
- web-quality-audit (merged into core-web-vitals)
- monetization-strategy (converted to reference)
- product-strategy (converted to reference)
- marketing-ideas (pure bloat, deleted)

### Technical Details

**Stats**
- 31 commits
- 24 skills (was 32)
- 5 reference docs
- 2 automation scripts
- 3 unified stacks
- 6 core documentation files

**Key Features**
- ✅ Single source of truth
- ✅ Zero duplication
- ✅ Zero bloat
- ✅ Generic (works with any AI harness)
- ✅ Fully automated setup
- ✅ Token-optimized
- ✅ Research-backed

**Based On**
- Peter Steinberger: "Just Talk To It", "Shipping at Inference-Speed"
- Mitchell Hashimoto: Vibe engineering, harness engineering
- Simon Willison: AI amplifies senior engineering practices
- Anthropic: Claude Code Best Practices (context management)

---

## Future Considerations

- Additional skills based on usage patterns
- Project-specific skill templates
- Integration with more AI harnesses
- Community skill contributions

---

[1.0.0]: https://github.com/yourusername/agent/releases/tag/v1.0.0
