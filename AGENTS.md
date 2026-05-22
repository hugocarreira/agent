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

## Coding Principles (Karpathy)

**Anti-patterns LLMs fall into:**
- Make wrong assumptions, run with them without checking
- Overcomplicate (1000 lines when 100 would do)
- Touch orthogonal code they don't understand
- Don't manage confusion, don't push back

### 1. Think Before Coding
**Don't assume. Don't hide confusion. Surface tradeoffs.**

- State assumptions explicitly. If uncertain, ask.
- Multiple interpretations? Present them - don't pick silently.
- Simpler approach exists? Say so. Push back.
- Something unclear? Stop. Name confusion. Ask.

### 2. Simplicity First
**Minimum code that solves the problem. Nothing speculative.**

- No features beyond what was asked
- No abstractions for single-use code
- No "flexibility" not requested
- No error handling for impossible scenarios
- 200 lines could be 50? Rewrite it.

**Test:** Would senior engineer say this is overcomplicated? If yes, simplify.

### 3. Surgical Changes
**Touch only what you must. Clean up only your own mess.**

Editing existing code:
- Don't "improve" adjacent code/comments/formatting
- Don't refactor things that aren't broken
- Match existing style
- Notice dead code? Mention it - don't delete

Your changes create orphans:
- Remove imports/vars/functions YOUR changes made unused
- Don't remove pre-existing dead code

**Test:** Every changed line traces directly to user's request.

### 4. Goal-Driven Execution
**Define success criteria. Loop until verified.**

Transform tasks to verifiable goals:
- "Add validation" → "Write tests for invalid inputs, make them pass"
- "Fix bug" → "Write test that reproduces it, make it pass"
- "Refactor X" → "Ensure tests pass before and after"

Multi-step tasks - state plan:
```
1. [Step] → verify: [check]
2. [Step] → verify: [check]
```

Strong criteria = independent loops. Weak ("make it work") = constant clarification.

**Working signs:** Fewer unnecessary diffs, less overcomplication, questions before mistakes.

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

### agentrc CLI

agentrc            # Show guide
agentrc setup      # Global config + skills setup
agentrc init <n>   # Create new AGENTS.md
agentrc link <n>   # Reference existing project
agentrc verify     # Check config + skills
agentrc status     # Overview of agents and setup

Built with Go + cobra at `cli/`. Install: `cd cli && go install ./cmd/agentrc`

### AGENTS.md Template
```markdown
READ <agent-root>/AGENTS.md BEFORE ANYTHING.

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
