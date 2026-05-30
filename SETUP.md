# Setup

## Quick Start

```bash
agentrc setup      # Link global config + skills (auto-detects agents)
agentrc verify     # Check setup + current project pack
agentrc init foo   # New project → foo/{AGENTS.md,FEATURES.yaml,PROGRESS.md}
agentrc link bar   # Existing project → prepend global ref
```

## Project Pack

```markdown
# AGENTS.md
## Stack
- Next.js 16, PostgreSQL, bun

## Commands
rtk bun test
rtk bun run dev

## Out of Scope
- Never touch: lib/billing/
```

```yaml
# FEATURES.yaml
features:
  - id: signup
    behavior: User can create an account
    done_when:
      - signup page renders
      - form submits successfully
      - success path redirects to app
    state: active
```

```markdown
# PROGRESS.md
## In Progress
- Signup backend done

## Next Session Start Here
- Run app, finish submit flow, verify redirect
```

Keep `AGENTS.md` short. Keep `FEATURES.yaml` concrete. Update `PROGRESS.md` only after substantial sessions.

## Common

```bash
# Update global (affects ALL projects)
cd <agent-root>
# Edit AGENTS.md, RTK.md → commit

# Config or skills broken?
agentrc setup   # refreshes all global config
```

See `agentrc --help` for all commands.
