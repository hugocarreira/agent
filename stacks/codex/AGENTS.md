@/home/hugoubuntu/work/agent/RTK.md

# Codex Configuration

Work style: telegraph; noun-phrases ok; drop grammar; min tokens.

## Core Principles (from Peter Steinberger workflow)

- **Just Talk To It**: Short prompts (1-2 sentences) + screenshots
- **Shipping at Inference-Speed**: Expect code works out-of-the-box
- **Blast Radius Thinking**: Know how many files a task will touch before starting
- **Iterate Fast, Refactor Smart**: 80% building, 20% refactoring

## Workflow
- Work on `main` by default (no worktrees unless asked)
- Atomic commits: one feature/fix per commit
- Test in same context (finds bugs automatically)
- Cross-reference projects: "look at ../other-project and do same"

## Prompting Style
- Short prompts + screenshots (~50% of prompts should include images)
- "give me options before making changes" for big tasks
- Show don't tell: drag screenshot + "fix padding"

## Code Quality
- Refactor when prompts get slow (~20% of time)
- Add inline comments for tricky logic
- Write tests after each feature (in same context)

## Git
- Conventional Commits format
- Only push when asked
- No destructive ops without consent

## Tools
- **CLIs > MCPs**: Prefer `gh`, `vercel`, `psql`, `bun`
- **RTK**: Always prefix shell commands (see RTK.md)

## Skills
Skills are centralized in `~/work/agent/skills/` (symlinked to `~/.codex/skills`)

---

READ ~/work/agent/AGENTS-GLOBAL.md for orchestrator rules (when using OpenCode).
