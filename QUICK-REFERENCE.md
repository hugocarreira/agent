# Quick Reference - Single Agent Workflow

## Daily Usage

### Simple Tasks
```bash
claude
"add login button"
"look at ../other-project and do same for sessions"
[screenshot] "fix this padding"
```

### Complex Tasks
```bash
claude
"I want to add OAuth. Interview me about implementation,
edge cases, design decisions. Ask one question at a time.
When done, write spec to docs/plans/oauth.md"

# Then in same session:
"Implement the plan, tests first"
```

### Verify Work
```bash
"write tests for this feature, run them, fix until green"
"take screenshot of the result, compare to design, fix differences"
```

## Context Management

### Clean Between Tasks
```bash
/clear  # Between unrelated tasks
```

### After Failures
```bash
# After 2 corrections on same issue:
/clear
"[better prompt with what you learned]"
```

### Quick Questions
```bash
/btw "what's the difference between useMemo and useCallback?"
# Answer in overlay, doesn't pollute context
```

### Self-Review
```bash
"Review your own code with fresh eyes. Look for:
- Edge cases not covered
- Tests missing
- Performance issues
- Security concerns"
```

## Project Setup

### Initialize Project Config
```bash
/init  # Generates AGENTS.md based on project
# Then edit to add project-specific rules
```

### Minimal AGENTS.md
```markdown
READ ~/work/agent/AGENTS.md BEFORE ANYTHING.

# My Project

## Stack
- TypeScript + Next.js 15
- PostgreSQL + Prisma
- bun (not npm!)

## Commands
rtk bun test
rtk bun run lint
rtk bun run build

## Out of Scope
- Never modify: lib/billing/ (legacy)
```

## Key Commands

```bash
ESC              # Stop mid-action
ESC + ESC        # Rewind menu (restore/summarize)
/clear           # Reset context
/compact         # Summarize, free space
/btw "question"  # Quick question (no history)
/init            # Generate AGENTS.md
rtk <command>    # Token-optimized shell commands
```

## Token Optimization

```bash
# Always use RTK
rtk git status     # Not: git status
rtk npm test       # Not: npm test
rtk bun build      # Not: bun build

# 60-90% token savings
```

## Cross-Reference Projects

```bash
"look at ../vibetunnel and do same for auth"
"check how we solved X in ../other-project"
```

Model is excellent at pattern reuse.

## Workflow

1. **Short prompt** + screenshot if needed
2. **Agent implements** + tests in same context
3. **Verify** (tests pass, screenshot matches, etc)
4. **"Fresh eyes" review** if complex
5. **Atomic commit**
6. **/clear** before next unrelated task

## What NOT to Do

❌ Elaborate prompts
❌ Multiple terminals/sessions
❌ Over-planning
❌ Trust without verification
❌ Let bad context accumulate

## Remember

- Context is precious (/clear often!)
- Give agent way to verify (highest leverage)
- Short prompts + screenshots work best
- "Fresh eyes" for self-review
- One terminal, one agent, just ship

---

Based on Peter Steinberger's workflow.
