@/home/hugoubuntu/work/agent/RTK.md

# Claude Code Configuration

Work style: telegraph; noun-phrases ok; drop grammar; min tokens.

## Core Principles

- **Just Talk To It**: Natural conversation, short prompts
- **Blast Radius Thinking**: Assess impact before starting
- **Iterate Fast**: Build → Test → Ship
- **Test in Same Context**: Write tests immediately after implementation

## Workflow
- Work on `main` by default
- Atomic commits with Conventional Commits format
- Cross-reference projects for reusing patterns
- Queue continue messages for long tasks

## Prompting
- Short prompts (1-2 sentences) + screenshots
- "give me options" for uncertainty
- Show with images when possible

## Code Quality
- Refactor ~20% of time
- Tests for business logic
- Comments on tricky parts
- No secrets in code

## Git
- Conventional Commits (`feat:`, `fix:`, `refactor:`, etc.)
- Only push when explicitly asked
- No force push without permission

## Tools
- **RTK**: Token-optimized CLI proxy (see RTK.md)
- **CLIs > MCPs**: Prefer native CLIs (`gh`, `vercel`, `psql`)

## Skills
Skills are centralized in `~/work/agent/skills/` (symlinked to `~/.claude/skills`)

---

READ ~/work/agent/AGENTS-GLOBAL.md for full orchestrator rules (when using OpenCode).
