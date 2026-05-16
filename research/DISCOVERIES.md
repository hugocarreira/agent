# Latest Discoveries (May 2026)

## Sources Analyzed
- Peter Steinberger: "Just Talk To It", "Shipping at Inference-Speed"
- Simon Willison: "Vibe Engineering" (LLMs amplify senior engineering practices)
- Armin Ronacher: "Plan Mode Is Just a Prompt"
- Anthropic: "Claude Code Best Practices" (context management)

## Key Insights Applied

### 1. Just Talk To It (Peter Steinberger)
Short prompts + screenshots. Natural conversation. No elaborate setups.
```
✅ "add login button"
✅ [screenshot] "fix padding"
✅ "look at ../other-project and do same"

❌ "I need you to carefully analyze..."
```

### 2. Context Management = Critical
Context window fills fast. Performance degrades.

Strategies:
- `/clear` between unrelated tasks
- `/compact` to summarize (keeps decisions, frees space)
- `/btw` for quick questions (doesn't enter history)
- After 2 corrections, /clear and restart with better prompt

### 3. "Give Agent a Way to Verify" = Highest Leverage
Tests, screenshots, expected outputs.
Agent can check itself = dramatically better performance.

```
Before: "make dashboard look better"
After:  "[screenshot] implement this. Take screenshot,
        compare to original, list differences, fix them"
```

### 4. AGENTS.md Must Be SHORT
If too long → agent ignores it.
Ruthlessly prune. Only include what agent can't infer from code.

✅ Include: bash commands, code style diffs, test commands, gotchas
❌ Exclude: standard conventions, API docs, file descriptions

### 5. CLI Tools > MCPs
MCPs pollute context (GitHub MCP = 23k tokens!).
Use: gh, vercel, psql, bun

### 6. "Fresh Eyes" Magic Phrase
When agent finishes work:
```
"Review your own code with fresh eyes. Look for edge cases,
missing tests, performance issues, security concerns."
```
Seems to reset bias toward code it just wrote.

### 7. Blast Radius Thinking
Before starting, ask yourself:
- How many files?
- How long?
- Can I context-switch if needed?

### 8. Iterate Fast, Refactor Smart
80% building, 20% refactoring.
Refactor when prompts get slow.

### 9. Cross-Reference Projects
```
"look at ../other-project and do the same"
"check how we solved X in ../vibetunnel"
```
Model is excellent at pattern reuse.

### 10. Course Correct Early
- ESC: stop mid-action
- After 2 corrections: /clear + better prompt
- Don't let bad context accumulate

## What We Simplified

### Removed (Over-Engineering):
❌ Parallel agents (multi-terminal workflows)
❌ Architect/Implementer pattern (two sessions)
❌ Subagents for investigation
❌ YOLO mode in Codespaces
❌ Git worktrees (unless user asks)
❌ Multi-agent orchestration
❌ Complex delegation patterns

### Kept (What Actually Works):
✅ Just talk to it (short prompts + screenshots)
✅ Test in same context (finds bugs!)
✅ Give agent way to verify (highest leverage)
✅ Context management (/clear, /compact, /btw)
✅ Cross-reference projects (pattern reuse)
✅ CLIs > MCPs (less pollution)
✅ RTK everywhere (token efficiency)
✅ "Fresh eyes" self-review
✅ Course correct early (ESC, /clear)
✅ One terminal, one agent (simplicity)

## Philosophy

**Single agent, single terminal, just ship.**

Peter Steinberger doesn't use multi-terminal workflows.
He uses 1 agent, talks naturally, iterates fast, ships.

We follow that.

## Stats
- AGENTS.md: 495 lines → 365 lines (simplified!)
- Skills: 32 (unchanged)
- Philosophy: "Just Talk To It" + essential best practices
- Stacks unified: 3 (opencode, codex, claude)
- Complexity: MUCH lower
