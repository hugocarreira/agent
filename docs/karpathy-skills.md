# Karpathy-Inspired Coding Guidelines

**Source:** https://github.com/multica-ai/andrej-karpathy-skills (146k stars)

Behavioral guidelines derived from [Andrej Karpathy's observations](https://x.com/karpathy/status/2015883857489522876) on LLM coding pitfalls. Already embedded in our [AGENTS.md](../AGENTS.md) as core principles.

## The 4 Principles

| Principle | Prevents |
|-----------|----------|
| **Think Before Coding** | Wrong assumptions, hidden confusion, silent tradeoffs |
| **Simplicity First** | Overcomplication, bloat, speculative abstractions |
| **Surgical Changes** | Orthogonal edits, touching code you shouldn't |
| **Goal-Driven Execution** | Vague instructions, no verification loop |

## Key Insight

> "Don't tell it what to do, give it success criteria and watch it go."

Transform imperative tasks → declarative goals with verification loops. This is why our setup gives agents a way to verify.

## Install Options

**Claude Code plugin:**
```
/plugin marketplace add forrestchang/andrej-karpathy-skills
/plugin install andrej-karpathy-skills@karpathy-skills
```

**Standalone CLAUDE.md:**
```bash
curl -o CLAUDE.md https://raw.githubusercontent.com/forrestchang/andrej-karpathy-skills/main/CLAUDE.md
```

**Cursor:** Uses `.cursor/rules/karpathy-guidelines.mdc` — see [CURSOR.md](https://github.com/multica-ai/andrej-karpathy-skills/blob/main/CURSOR.md).

## Tradeoff

Biases toward **caution over speed**. For trivial tasks (typo fixes, one-liners) skip full rigor.
