# Agent Instructions

**Just talk to it.** No orchestrators, no charades, no plan mode theater.

Work style: telegraph; noun-phrases ok; drop grammar; min tokens.

---

## Core Philosophy

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
Before starting, ask yourself:
- How many files will this touch?
- How long should this take?
- Can this run in parallel?

If stuck, ESC + "what's the status"

---

## What Actually Works (Vibe Engineering - Simon Willison)

LLMs **amplify existing senior engineering practices**:

### Automated Testing (Highest Leverage!)
- Tests for all business logic (unit + integration)
- **Write tests in same context** as implementation (finds bugs automatically!)
- Test-first when possible
- Give Claude a way to verify its work = single highest leverage thing

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
- Start with conversation, ask clarifying questions
- Write to `docs/*.md` if complex
- Iterate on plan before building
- Don't over-plan, iterate fast
- **Explore → Plan → Code → Commit** workflow

### Cross-Reference Projects
```bash
"look at ../other-project and do the same"
"check how we solved X in ../vibetunnel"
```
Model is great at reusing patterns.

### CLI-First
- Start with CLI, build UI later
- Easier to test, model can verify output
- Closes the loop faster

---

## Parallel Agents (Simon Willison + Jesse Vincent)

Run 3-8 agents in parallel. Patterns that work:

### 1. Research / Proof of Concepts
Fire off agent to answer: "Can library X work with our stack?"
No intention to keep the code, just want the answer.

### 2. "How Does This Work Again?"
"Explain our signed cookie flow across all files"
Stash the explanation for future context.

### 3. Small Maintenance Tasks
"Fix this deprecation warning by running tests and figuring it out"
Low stakes, outsource the cognitive overhead.

### 4. Architect / Implementer Pattern (Jesse Vincent)

**Architect Session:**
```bash
# Brainstorm
"I've got an idea. Ask me questions one at a time (multiple choice preferred)
to help refine it. Output 200-300 words at a time, asking if it looks right.
Write the design to docs/plans/feature-name.md"

# Plan
"Write comprehensive implementation plan. Assume engineer has zero context.
Bite-sized tasks. DRY. YAGNI. TDD. Frequent commits.
Write to docs/plans/feature-name-plan.md"
```

**Implementer Session (NEW terminal!):**
```bash
claude  # fresh session
"Read docs/plans/feature-name-plan.md. Let me know if you have questions."
"Execute tasks 1-4. DO NOT DEVIATE FROM THE PLAN."

# When done, flip to Architect:
"Implementer says tasks 1-4 done. Review carefully."

# Copy review back to Implementer, fix issues

# CRITICAL: /clear the Implementer after each chunk!
# Double-ESC the Architect to reset without bias
```

**Why this works:**
- Architect maintains design intent
- Implementer has clean context per chunk
- Two "actors" = fresh eyes on reviews
- No context bloat

### 5. Send Out a Scout (Josh Snyder)
Give agent a genuinely difficult task with **no intention of landing code**.
Just to see which files it modifies and how it approaches the problem.
Learn from its mistakes without making them yourself.

### Isolation Techniques
- **Git worktrees**: `cd .worktrees && git worktree add feature-name`
- **/tmp checkouts**: `git clone . /tmp/experiment`
- **Docker containers**: Claude Code with `--dangerously-skip-permissions` in container
- **GitHub Codespaces**: VS Code agent in browser, someone else's computer!

---

## YOLO Mode (The Joy of Auto-Approve)

**YOLO mode = auto-approve everything = CRITICAL for productivity**

"An AI agent is an LLM wrecking its environment in a loop." - Solomon Hykes

### Risks:
1. Bad shell commands delete things
2. Exfiltration (steal source code, secrets)
3. Use your machine as proxy for attacks

### Mitigation:
1. **Docker/containers** (best): Restrict files + network
2. **Someone else's computer**: Codespaces, cloud VMs
3. **Take a risk**: Avoid malicious prompts, hope for the best

Most people choose #3. You should choose #1 or #2.

### Anthropic's Safe YOLO
```bash
# In Docker without internet (or allowlist: api.anthropic.com, github.com, etc)
claude --dangerously-skip-permissions
```

Lock internet to trusted hosts = prevent exfiltration.

---

## Designing Agentic Loops (Simon Willison)

Look for problems with:
- **Clear success criteria** (tests pass, screenshot matches, build succeeds)
- **Trial and error** needed (ugh, gotta try many variations)

Examples:
- **Debugging**: test failing, let agent iterate until green
- **Performance**: SQL slow, try indexes + benchmark
- **Upgrades**: dependency updates, run tests, fix breaking changes
- **Container size**: try different base images, keep tests passing

**Critical**: Solid test suite = force multiplier.

---

## Workflow

### Iterate Fast, Refactor Smart
- **80% building** → **20% refactoring**
- Work on `main` by default (no worktrees unless parallel work)
- Refactor when prompts get slow
- Commit atomically after each feature

### Code Quality
- TypeScript strict mode, no `any` without justification
- Proper error handling (never swallow)
- No secrets in code (env vars only)
- No `console.log` in production

### Prompting
- Short prompts + screenshots (~50% should include images)
- Show don't tell: `[screenshot] "fix padding"`
- Trigger words for hard tasks: "take your time", "comprehensive", "read all related code"
- **"Fresh eyes"** magic phrase when reviewing own work

### Context Management (CRITICAL!)

**Context window is your most important resource.**

- `/clear` between unrelated tasks
- `/compact` to summarize (keeps key decisions, frees space)
- `ESC + ESC` or `/rewind` to restore/summarize from checkpoint
- Use **subagents for investigation** (separate context!)
- `/btw` for quick questions (doesn't enter history)

**When to /clear:**
- Between unrelated tasks
- After 2+ corrections on same issue (context polluted with failed approaches)
- When session feels "slow" or agent seems confused

**Signs of context bloat:**
- Agent ignoring instructions
- Repeating mistakes
- Slow responses
- Irrelevant file reads

### Course Correct Early
- `ESC`: stop mid-action, redirect
- `ESC + ESC` or `/rewind`: restore previous state
- `"undo that"`: revert changes
- **After 2 corrections, /clear and start fresh with better prompt**

---

## Tools & Efficiency

### RTK - Token Killer
**Always prefix shell commands with `rtk`**

```bash
rtk git status
rtk npm test
rtk cargo build
```

60-90% token savings. See `RTK.md` for full reference.

### CLIs > MCPs
**MCPs pollute context.** GitHub MCP = 23k tokens!

Prefer CLIs with good help:
- `gh` (GitHub)
- `vercel` (deploy)
- `psql` (database)
- `bun` (JavaScript runtime)

Only use MCPs when no CLI exists.

### MCP via CLI (Peter Steinberger)
Progressive disclosure, no context cluttering:
```bash
pnpm mcp:list                                    # list available
pnpm mcp:list chrome-devtools --schema          # see schema
pnpm mcp:call playwright.browser_tabs action=list
pnpm mcp:call chrome-devtools.evaluate_script --args '{"function":"() => document.title"}'
```

### Skills Discovery
32 skills available in `skills/` folder:
- Development: `commit-and-push`, `pullrequest`, `test-and-verify`
- Frontend: `frontend-design`, `accessibility`, `core-web-vitals`
- Next.js: `next-best-practices`, `next-upgrade`
- SEO: `seo`, `seo-audit`, `programmatic-seo`
- Product: `product-strategy`, `page-cro`, `signup-flow-cro`

Skills auto-trigger on relevant keywords. List: `ls ~/work/agent/skills/`

---

## Bash Safety

### NEVER (without explicit permission):
- `rm -rf` outside project
- `sudo` anything
- `git push --force` to main
- `npm publish`, `docker push`
- `curl | bash` patterns
- Destructive git operations (`git reset --hard`, `git restore` to old commit)

### ALWAYS prefer:
- Read-only commands first (`ls`, `rtk grep`, `rtk find`)
- Scoped commands (`npm test`, not `npm run all`)
- Dry-run flags (`terraform plan`)

### Git Discipline (Peter Steinberger)
- **NEVER edit `.env`** or environment files
- **Coordinate before deleting** other agents' work
- Moving/renaming files is OK
- **ABSOLUTELY NEVER** destructive git ops without explicit instruction

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

❌ **No orchestrators** - Talk directly, no delegation theater
❌ **No plan mode charades** - Just write to `docs/*.md` if needed
❌ **No RAG systems** - Model can search fine
❌ **No subagent spawning** - Except for investigation (context separation)
❌ **No elaborate prompts** - Short and natural
❌ **No over-planning** - Iterate and ship
❌ **No trust without verification** - Always provide way to verify

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

## Mental Models

### Single Agent
```
User
  ↓ (short prompt + maybe screenshot)
Agent (reads AGENTS.md + RTK.md + project docs/ + code)
  ↓ (understands context)
  ↓ (implements + tests in same session)
  ↓ (atomic commit)
Result
```

### Parallel Agents (Architect/Implementer)
```
Architect Session              Implementer Session
      ↓                              ↓
  Brainstorm                    [NEW terminal]
      ↓                              ↓
  Plan (docs/plans/)            Read plan
      ↓                              ↓
  Review ←─────────────────────  Execute tasks 1-4
      ↓                              ↓
  Sign off                       /clear (fresh context!)
      ↓                              ↓
  Review ←─────────────────────  Execute tasks 5-8
   (ESC+ESC to reset)                ↓
      ↓                           /clear
      ...                           ...
```

No layers. Direct or parallel. No theater.

**The model is smart. Talk to it like a senior engineer.**

---

## Advanced Patterns

### Non-Interactive Mode
```bash
# CI, pre-commit hooks, scripts
claude -p "prompt" --output-format json

# Loop through tasks
for file in $(cat files.txt); do
  claude -p "Migrate $file" --allowedTools "Edit,Bash(git commit *)"
done
```

### Worktrees for Parallel Work
```bash
mkdir .worktrees  # first time
cd .worktrees
git worktree add feature-oauth
cd feature-oauth
npm install
claude  # isolated session!
```

### CodeRabbit Review with Role-Play (Jesse Vincent)
```
A reviewer analyzed this PR. They're external, reading codebase cold.

1) should we hire this reviewer?
2) which issues should be fixed?
3) are their proposed fixes correct?

Add valid fixes to your todo. Tell me about anything to skip.
```

Prevents credulous acceptance of bad suggestions.

---

## Summary

- **Just talk to it** (no charades)
- **Test in same context** (finds bugs automatically!)
- **Give agent a way to verify** (highest leverage!)
- **Cross-reference projects** (reuse patterns)
- **CLIs > MCPs** (less context pollution)
- **RTK everything** (token efficiency)
- **Iterate fast, refactor smart** (80/20)
- **Context is precious** (/clear, /compact, subagents)
- **YOLO mode in containers** (productivity + safety)
- **Parallel agents for big work** (Architect/Implementer pattern)
- **Design agentic loops** (clear success + trial/error)
- **Course correct early** (ESC, /clear after 2 corrections)

**That's it. Keep it simple. Ship fast. Stay in control.**

---

_Based on workflows from:_
- _[Peter Steinberger](https://steipete.me) - "Just Talk To It"_
- _[Simon Willison](https://simonwillison.net) - "Vibe Engineering", Parallel Agents, Agentic Loops_
- _[Armin Ronacher](https://lucumr.pocoo.org) - "Plan Mode Is Just a Prompt"_
- _[Jesse Vincent](https://blog.fsck.com) - Architect/Implementer Pattern_
- _[Anthropic](https://anthropic.com) - Claude Code Best Practices_
