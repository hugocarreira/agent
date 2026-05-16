# Agent Instructions

**Just talk to it.** No orchestrators, no subagents, no theater.

Work style: telegraph; noun-phrases ok; drop grammar; min tokens.

---

## Core Philosophy

### Just Talk To It (Peter Steinberger)
- Short prompts (1-2 sentences) + screenshots
- Natural conversation
- Iterate together, don't over-plan

### Give Agent a Way to Verify (HIGHEST LEVERAGE!)
**This is the #1 thing that makes AI effective.**

Before: "make dashboard better"  
After: "[screenshot] implement this. Take screenshot, compare, list differences, fix"

Verification methods:
- Tests that agent can run
- Screenshots to compare
- Benchmarks with clear criteria
- Expected outputs to diff
- CLI commands with known results

When agent can check itself = dramatically better performance.

### Context is Precious
Performance degrades as context fills.

- `/clear` between unrelated tasks
- `/clear` after 2 corrections (context polluted!)
- `/btw` for quick questions (no history)
- Watch for "slop zone" (repeated mistakes)

---

## What Works

### Automated Testing
- Rock-solid test suite = agents fly
- Write tests in same context (finds bugs automatically!)
- Test-first when possible

### Planning
**Vague request?** Separate planning from execution.

```
"I want to add OAuth. Interview me about it.
Ask one question at a time. When done, write spec to docs/plans/oauth.md"
```

Then: `"Implement the plan, tests first"`

**Clear task?** Just do it. No planning needed.

### Cross-Reference Projects
```bash
"look at ../other-project and do the same"
```
Model reuses patterns well.

### Cleanup Sessions ("Anti-Slop")
After agent work:
- Move functions to better places
- Add documentation  
- Restructure for clarity

Forces understanding. Improves future sessions.

---

## Mitchell's Patterns

### Fill-in-the-Blank
Create scaffold with TODOs → agent completes.

```python
def process_data(user_id: str) -> UserData:
    """Process and return structured result.
    TODO: fetch from db, validate, transform"""
    pass  # agent fills
```

### When Agent Fails
1. Try 2-3 approaches (vague → specific)
2. If still failing:
   - Stop
   - Back out
   - Educate yourself manually
   - Come back with better approach

AI isn't always the solution.

### Harness Engineering
Every mistake → fix permanently:
- Update AGENTS.md (wrong commands/APIs)
- Add tool (screenshot script, test filter)

Fix once, never again.

### End-of-Day Agents
Last 30min: kick off agents for:
- Deep research
- Issue triage
- Parallel vague ideas

"Warm start" next morning.

---

## Workflow

### Prompting
```
✅ "add login button"
✅ [screenshot] "fix padding"
✅ "write tests first, then implement"

❌ "I need you to carefully analyze..."
```

~50% of prompts should include screenshots.

### Course Correct
- ESC to stop mid-action
- After 2 corrections → /clear + better prompt
- Don't let bad context accumulate

### "Fresh Eyes" Review
```
"Review your own code with fresh eyes. Look for:
- Edge cases not covered
- Tests missing
- Performance issues
- Security concerns"
```

Magic phrase that resets agent's bias.

---

## Tools

### RTK - Token Killer
Always prefix shell commands:
```bash
rtk git status
rtk npm test
```
60-90% token savings.

### CLIs > MCPs
MCPs pollute context. GitHub MCP = 23k tokens!

Prefer: `gh`, `vercel`, `psql`, `bun`

---

## Project Setup

### AGENTS.md Template
```markdown
READ ~/work/agent/AGENTS.md BEFORE ANYTHING.

# My Project

## Stack
- TypeScript + Next.js
- PostgreSQL + Prisma

## Commands
rtk bun test
rtk bun run lint

## Out of Scope
- Never touch: lib/billing/
```

Keep it SHORT. Only what agent can't infer from code.

---

## What NOT to Do

❌ Orchestrators  
❌ Plan mode theater  
❌ Multi-terminal workflows  
❌ Elaborate prompts  
❌ Trust without verification  

---

## Summary

1. **Give agent way to verify** (highest leverage!)
2. **Context is precious** (/clear often)
3. **Just talk to it** (short prompts + screenshots)
4. **Test in same context** (finds bugs!)
5. **Final manual review** (always!)

One terminal. One agent. Just ship.

---

_Based on: Peter Steinberger, Mitchell Hashimoto, Simon Willison, Anthropic_
