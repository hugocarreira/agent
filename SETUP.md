# Setup

## Quick Start

```bash
agentrc setup      # Link global config + skills (auto-detects agents)
agentrc verify     # Check everything
agentrc init foo   # New project → foo/AGENTS.md
agentrc link bar   # Existing project → prepend global ref
```

## Project AGENTS.md Template

```markdown
## Stack
- Next.js 16, PostgreSQL, bun

## Commands
rtk bun test
rtk bun run dev

## Out of Scope
- Never touch: lib/billing/
```

Keep it SHORT. Only what agent can't infer from code.

## Common

```bash
# Update global (affects ALL projects)
cd <agent-root>
# Edit AGENTS.md, RTK.md → commit

# Config or skills broken?
agentrc setup   # refreshes all global config
```

See `agentrc --help` for all commands.
