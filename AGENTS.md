# Agent Instructions

**Just talk to it.** No orchestrators, no subagents, no multi-terminal theater.

Work style: telegraph; noun-phrases ok; drop grammar; min tokens.

---

## Core Philosophy (Peter Steinberger)

### Just Talk To It
- Short prompts (1-2 sentences) + screenshots
- Natural conversation, no elaborate prompts
- "give me options" when uncertain
- Iterate together, don't over-plan

### Shipping at Inference-Speed
- Expect code works out-of-the-box
- Watch the stream, don't read every line
- Trust the model when context is solid

### Blast Radius Thinking
Before starting, mentally ask:
- How many files will this touch?
- How long should this take?
- Can I start this and context-switch if needed?

If stuck, ESC + "what's the status"

---

## What Actually Works (Vibe Engineering - Simon Willison)

**"AI tools amplify existing senior engineering expertise."**

LLMs amplify what you're already good at. They don't fix what you're bad at.

### Give Agent a Way to Verify (HIGHEST LEVERAGE!)

**This is the #1 thing that makes AI effective.**

Before: "make the dashboard better"
After: "[screenshot] implement this. Take screenshot, compare to original, list differences, fix them"

Other verification methods:
- Tests that agent can run (unit, integration, e2e)
- Benchmarks with clear success criteria
- Screenshots to compare
- Expected outputs to diff against
- CLI commands with known results

**When agent can check itself = dramatically better performance.**

### Automated Testing
- Rock-solid test suite = agents fly
- **Write tests in same context** as implementation (finds bugs automatically!)
- Test-first when possible
- No tests = agent flounders

### Documentation
- Keep `docs/` folder with frontmatter (`summary`, `read_when`)
- Update docs for user-facing changes
- Inline comments for tricky/bug-prone logic
- README for onboarding

### Version Control
- Conventional Commits (`feat:`, `fix:`, `refactor:`, `chore:`, `docs:`)
- Atomic commits (one logical change)
- Only push when asked
- No force push without permission
- Use `git bisect` when hunting bugs

### Planning (Simple, Not Theater)
**Explore → Plan → Code → Commit**

**For vague requests: separate planning from execution.**

When unsure, ask agent to ask YOU questions:
```
"I want to add OAuth. Interview me about implementation details,
edge cases, and design decisions. Ask one question at a time.
When done, write spec to docs/plans/oauth.md"
```

Then in same session: `"Implement the plan, tests first"`

**For clear tasks: just do it.** No planning needed.

Planning is about separation when YOU don't know what you want.
Not about making the agent think harder.

### Cross-Reference Projects
```bash
"look at ../other-project and do the same"
"check how we solved X in ../vibetunnel"
```
Model is excellent at reusing patterns.

### CLI-First
- Start with CLI, build UI later
- Easier to test, model can verify output
- Closes the loop faster

---

## Workflow Patterns (Mitchell Hashimoto)

### Mitchell's 6-Step Evolution

**1. Drop the Chatbot**
- Chatbots (ChatGPT web) are inefficient
- Use agents (read files, run commands, HTTP)

**2. Reproduce Your Own Work**
- Do work manually, then force agent to reproduce
- Excruciating but builds expertise
- Discover what works from first principles

**3. End-of-Day Agents**
- Last 30min: kick off 1+ agents
- Deep research, parallel vague ideas, issue triage
- "Warm start" next morning

**4. Outsource the Slam Dunks**
- High-confidence tasks → agent (background)
- You work on something else
- **Turn off notifications** (context switching expensive)
- Check during natural breaks

**5. Engineer the Harness**
- "Harness engineering" = prevent mistakes permanently
- Update AGENTS.md (wrong commands/APIs)
- Add tools (screenshots, filtered tests)
- Fix once, never happens again

**6. Always Have an Agent Running**
- Goal: agent running 10-20%+ of workday
- Not multiple agents (one is enough)
- Only when truly helpful task exists

### Non-Trivial Vibing (Mitchell)

Real shipping features, not toys.

**Process:**
1. **Pre-AI planning** (human thinks first!)
2. **Plan before code** (vague requests only)
3. **Fill-in-the-blank pattern:**
   - Create scaffold with TODOs
   - Agent completes the owl
4. **Cleanup sessions** ("anti-slop")
   - Forces understanding
   - Documents code
   - Better foundations → better results
5. **"I broke things, fix my mess"**
6. **Simulation/testing**
7. **Always ask "anything else?"**
8. **Final manual review** (SUPER important)

**Patterns:**
- Use AI for inspiration (sometimes throw away)
- "Hitting a wall" → back out, educate yourself
- Not about faster/slower, about parallel work
- Can work while doing other things

### Iterate Fast, Refactor Smart
- **80% building** → **20% refactoring**
- Work on `main` by default
- Refactor when prompts get slow
- Commit atomically after each feature

### Code Quality
- TypeScript strict mode, no `any` without justification
- Proper error handling (never swallow)
- No secrets in code (env vars only)
- No `console.log` in production

### Prompting That Works
Short prompts + screenshots (~50% should include images):
```
✅ "add login button"
✅ [screenshot] "fix padding here"
✅ "look at ../other-project auth and do same"
✅ "write tests first, then implement"

❌ "I need you to carefully analyze the system..."
```

Trigger words for hard tasks: "take your time", "comprehensive", "read all related code"

### Context Management (CRITICAL!)

**Context window is your most precious resource.**

Performance degrades as context fills. Manage aggressively.

#### When to /clear
- Between unrelated tasks
- **After 2+ corrections on same issue** (context polluted!)
- When session feels "slow" or agent seems confused
- "Slop zone" (agent making dumb mistakes repeatedly)

#### When to /compact
- Long session on same task
- Want to keep history but free space
- Tells agent: "Focus on API changes" or "Preserve test commands"

#### /rewind for Quick Fixes
- `ESC + ESC`: open rewind menu
- Restore conversation, code, or both
- Or summarize from a checkpoint

#### Quick Questions
`/btw "what's the difference between X and Y?"`
- Answer appears in dismissible overlay
- Never enters conversation history
- Doesn't pollute context

### Course Correct Early
- `ESC`: stop mid-action, redirect immediately
- `"undo that"`: revert changes
- **After 2 corrections, /clear and start fresh with better prompt**
- Don't let bad context accumulate
- "Slop zone" = time to back out and rethink

### When Agent Fails (Mitchell's Pattern)

1. **Make 2-3 hail mary attempts** (start vague, get specific)
2. **If still failing:**
   - Stop
   - Back out
   - Educate yourself manually
   - Come back with better understanding
3. **Don't keep bashing head against wall**
4. **Sometimes pivot to different approach**

AI isn't always the solution. Sometimes it's a liability.

### Fresh Eyes Review
When agent finishes implementation:
```
"Review your own code with fresh eyes. Look for:
- Edge cases not covered
- Tests missing
- Performance issues
- Security concerns"
```

The phrase "fresh eyes" seems to reset agent's bias toward code it just wrote.

---

## Tools & Efficiency

### RTK - Token Killer
**Always prefix shell commands with `rtk`**

```bash
rtk git status
rtk npm test
rtk cargo build
```

60-90% token savings. See `RTK.md`.

### CLIs > MCPs
**MCPs pollute context.** GitHub MCP alone = 23k tokens!

Prefer CLIs:
- `gh` (GitHub)
- `vercel` (deploy)
- `psql` (database)
- `bun` (JavaScript runtime)

Only use MCPs when no CLI exists.

Agent is great at learning CLIs: `"use foo-cli --help to learn it, then do X"`

### Skills Discovery
32 skills in `skills/` folder:
- Development: `commit-and-push`, `pullrequest`, `test-and-verify`
- Frontend: `frontend-design`, `accessibility`, `core-web-vitals`
- Next.js: `next-best-practices`, `next-upgrade`
- SEO: `seo`, `seo-audit`, `programmatic-seo`
- Product: `product-strategy`, `page-cro`, `signup-flow-cro`

Auto-trigger on keywords. List: `ls ~/work/agent/skills/`

---

## Bash Safety

### NEVER (without explicit permission):
- `rm -rf` outside project
- `sudo` anything
- `git push --force` to main
- `npm publish`, `docker push`
- `curl | bash` patterns
- Destructive git ops (`git reset --hard`, `git restore` to old commit)

### ALWAYS prefer:
- Read-only first (`ls`, `rtk grep`, `rtk find`)
- Scoped commands (`npm test`, not `npm run all`)
- Dry-run flags (`terraform plan`)

### Git Discipline
- **NEVER edit `.env`** or environment files
- Moving/renaming files is OK
- **Coordinate before deleting** files (could be from other work)

---

## Project Context

### Read First
1. Project's `AGENTS.md` or `CLAUDE.md` (if exists)
2. Relevant existing code (understand patterns)
3. `docs/` folder (check frontmatter `read_when`)

### AGENTS.md / CLAUDE.md Best Practices

**Run `/init` to generate starter based on your project.**

**Keep it SHORT** (context is precious!):

✅ **Include:**
- Bash commands agent can't guess
- Code style differing from defaults
- Test commands and runners
- Branch naming, PR conventions
- Architectural decisions specific to project
- Environment quirks (required env vars)
- Common gotchas

❌ **Exclude:**
- Things agent figures out from code
- Standard language conventions
- Detailed API docs (link instead)
- Frequently changing info
- Long explanations
- File-by-file descriptions
- Self-evident practices

**If agent ignores rules, file is too long.**
**If agent asks questions answered in file, phrasing is ambiguous.**

Use `@path/to/file` to import other files.

### Per-Project Template
```markdown
READ ~/work/agent/AGENTS.md BEFORE ANYTHING.

# [Project Name]

## Stack
- Language: TypeScript/Go/Python
- Framework: Next.js/Gin/FastAPI
- Database: PostgreSQL/MongoDB
- Package manager: bun/pnpm/npm

## Commands
rtk bun test           # Run tests
rtk bun run lint       # Lint
rtk bun run build      # Build

## Conventions
- API routes in: `app/api/`
- Components in: `components/`
- Use bun (not npm)

## Out of Scope
- Never modify: `lib/billing/` (legacy, don't touch)
```

---

## What NOT to Do

❌ **No orchestrators** - Talk directly, no delegation
❌ **No plan mode theater** - Just write to `docs/*.md` if needed
❌ **No subagents** - Single agent is enough
❌ **No multi-terminal workflows** - One terminal, one session
❌ **No elaborate prompts** - Short and natural
❌ **No over-planning** - Iterate and ship
❌ **No trust without verification** - Always give agent way to verify

---

## Error Recovery

1. **Self-fix** - Read error, fix once (ONE attempt)
2. **Simplify** - Strip to minimal reproduction
3. **Stop and report**:
   ```
   ❌ BLOCKED
   Task: [what I was doing]
   Error: [exact message]
   Tried: [fix attempts]
   Assessment: [what I think is wrong]
   ```
4. **Ask user** - Only after 2 failed attempts

Never retry same fix twice. Never swallow errors.

---

## Advanced: Non-Interactive Mode

For CI, pre-commit hooks, scripts:
```bash
# One-off
claude -p "explain what this project does"

# In scripts
claude -p "list all API endpoints" --output-format json

# Loop through tasks
for file in $(cat files.txt); do
  claude -p "migrate $file to new API" --allowedTools "Edit"
done
```

---

## Mental Model

```
User
  ↓ (short prompt + maybe screenshot)
Agent (reads AGENTS.md + RTK.md + project docs/ + code)
  ↓ (understands context)
  ↓ (implements + tests in same session)
  ↓ (atomic commit)
Result
```

No layers. No delegation. No theater. No multi-terminal.

**The model is smart. Talk to it like a senior engineer.**

One terminal. One agent. Just ship.

---

## Summary

- **Just talk to it** (no charades)
- **Test in same context** (finds bugs automatically!)
- **Give agent a way to verify** (highest leverage!)
- **Cross-reference projects** (reuse patterns)
- **CLIs > MCPs** (less context pollution)
- **RTK everything** (token efficiency)
- **Iterate fast, refactor smart** (80/20)
- **Context is precious** (/clear, /compact, /btw)
- **Course correct early** (ESC, /clear after 2 fails)
- **"Fresh eyes" for self-review** (magic phrase)
- **One terminal, one agent** (keep it simple)

**That's it. Keep it simple. Ship fast. Stay in control.**

---

_Based on Peter Steinberger's workflow: [steipete.me](https://steipete.me)_
_Context management from Anthropic Claude Code Best Practices_
_"Vibe Engineering" concept from Simon Willison_
