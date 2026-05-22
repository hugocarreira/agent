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
agentrc init <n>   # Create AGENTS.md for new project
agentrc link <n>   # Reference existing project
agentrc verify     # Check everything's working
agentrc status     # Overview of all agents and config
agentrc guide      # Friendly walkthrough
```

## Structure

```
<repo>/
├── AGENTS.md  RTK.md  SETUP.md
├── cli/       # Go CLI + install.sh
├── skills/    # 26 reusable workflows
└── plugins/   # OpenCode plugins
```

Agent reads: harness config → global rules → project rules → your code.

## Principles

1. **Give Agent Way to Verify** — Tests, screenshots, benchmarks
2. **Context is Precious** — `/clear` often, `/btw` for quick questions
3. **Just Talk To It** — Short prompts + screenshots, no theater
4. **Final Manual Review** — Always understand what agent did

26 skills · 6 agents · 1 source of truth · 0 theater
