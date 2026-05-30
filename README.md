# AI Agent Setup — Centralized

Single source of truth for 6 AI coding agents:

- OpenCode
- Codex (ChatGPT)
- Claude Code
- GitHub Copilot
- Hermes
- Gemini CLI

Just talk to it. Give agent way to verify. Ship fast.

## Setup

```bash
# One-liner (no deps)
curl -fsSL https://raw.githubusercontent.com/hugocarreira/agentrc/main/cli/install.sh | bash
agentrc setup

# Or tell your AI: "install the harness from github.com/hugocarreira/agentrc"
```

## CLI

```bash
agentrc setup      # Link global config + skills (one-time)
agentrc init <n>   # Create AGENTS.md + FEATURES.yaml + PROGRESS.md
agentrc link <n>   # Reference existing project
agentrc verify     # Check global setup + current project pack
agentrc status     # Overview of all agents and config
agentrc guide      # Friendly walkthrough
```

## Structure

```
agentrc/
├── AGENTS.md  RTK.md  SETUP.md
├── cli/       # Go CLI + install.sh
├── skills/    # 16 focused workflows
└── plugins/   # OpenCode plugins
```

Agent reads: harness config → global rules → project rules → current state → your code.

## Project Pack

`agentrc init` creates:

- `AGENTS.md` for project-specific rules, stack, commands, and boundaries
- `FEATURES.yaml` for active feature scope and acceptance criteria
- `PROGRESS.md` for handoff notes, blockers, and exact restart point

`FEATURES.yaml` keeps feature state lightweight:

```yaml
features:
  - id: pricing-page
    behavior: Site has a pricing page with monthly and yearly plans
    done_when:
      - /pricing route renders
      - toggle changes displayed prices
      - primary CTA links to signup
    check:
      type: manual
      note: Run app and verify in browser
    state: active
```

Update `PROGRESS.md` at the end of substantial sessions. Update `FEATURES.yaml` only when feature state, scope, or acceptance changes.

## Principles

1. **Give Agent Way to Verify** — Tests, screenshots, benchmarks
2. **Context is Precious** — `/clear` often, `/btw` for quick questions
3. **Just Talk To It** — Short prompts + screenshots, no theater
4. **Final Manual Review** — Always understand what agent did

16 skills · 6 agents · 1 source of truth · 0 theater
