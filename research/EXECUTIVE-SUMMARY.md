# Executive Summary - Modern AI Dev Practices

**TL;DR:** AI tools amplify existing senior engineering expertise. The highest leverage action is giving agents a way to verify their work.

---

## Top 10 Insights (Read This First!)

### 1. Give Agent a Way to Verify = HIGHEST LEVERAGE
Tests, screenshots, benchmarks, expected outputs. When agent can check itself = dramatically better performance.

**Before:** "make dashboard better"  
**After:** "[screenshot] implement this. Take screenshot, compare, list differences, fix them"

### 2. AI Amplifies What You're Already Good At
Rock-solid tests → agents fly.  
No tests → agents flounder.

AI doesn't fix what you're bad at.

### 3. Context is Precious, Manage Aggressively
Performance degrades as context fills.

- /clear between unrelated tasks
- /clear after 2 corrections (polluted context!)
- /btw for quick questions (no history)
- Watch for "slop zone" (repeated dumb mistakes)

### 4. Just Talk To It (Peter Steinberger)
Short prompts (1-2 sentences) + screenshots work best.

Models got better → prompts got shorter.

❌ "I need you to carefully analyze..."  
✅ "add login button"  
✅ "[screenshot] fix padding"

### 5. Harness Engineering = Prevent Mistakes Forever (Mitchell)
Every time agent makes mistake:
- Update AGENTS.md (wrong commands/APIs)
- Add tool (screenshot script, filtered tests)

Fix once, never happens again. Accumulates over time.

### 6. Final Manual Review = SUPER Important
**Never skip.** Read every line. Understand what it does.

Not about trust. About understanding and quality.

### 7. Plan Only When YOU Don't Know What You Want
Clear task → just do it.  
Vague request → separate planning from execution.

Planning is about separation when YOU'RE uncertain, not making agent "think harder."

### 8. Cleanup Sessions ("Anti-Slop")
After agent work:
- Move functions to better places
- Add documentation
- Restructure for maintainability

Forces understanding. Improves future sessions.

### 9. When Agent Fails: Back Out, Educate Yourself
After 2-3 attempts, if still failing:
- Stop
- Back out
- Educate yourself manually
- Come back with better approach

AI isn't always the solution. Sometimes it's a liability.

### 10. Non-Trivial Vibing = Real Shipping Features
Not toy demos. Real production code. Thousands of users.

Mitchell's example: $15.98 tokens, ~8 hours, shipping to thousands.

Process: Plan → Fill-in-blank → Cleanup → Simulate → Test → Ship

---

## Three Workflow Styles

### Peter Steinberger: "Just Talk To It"
- Single agent, single terminal
- Short prompts + screenshots
- Cross-reference projects
- "Fresh eyes" self-review
- Expect one-shot success

**Philosophy:** Watch the stream, trust the model, iterate fast.

### Mitchell Hashimoto: "Non-Trivial Vibing"
- Pre-AI planning (human thinks first!)
- Fill-in-the-blank pattern (scaffold + TODOs)
- Cleanup sessions (anti-slop)
- Harness engineering (prevent repeated mistakes)
- Always have agent running (10-20% of workday)

**Philosophy:** Not faster/slower, about parallel work while doing other things.

### Simon Willison: "Vibe Engineering"
- AI amplifies existing expertise
- Rock-solid test suite
- Good documentation
- Effective automation
- Parallel agents for research (3-8 instances)

**Philosophy:** "AI tools amplify existing senior engineering expertise."

---

## Key Terms

**Vibe Coding:** Irresponsible (accept everything, no review)  
**Vibe Engineering:** Responsible (amplify expertise, final review)  
**Agentic Engineering:** Same as vibe engineering  

**Harness Engineering:** Prevent mistakes permanently  
**Non-Trivial Vibing:** Real shipping features (not toys)  

**AI Engineer:** Ships products using models (no PhD needed)  
**ML Engineer:** Trains models (often needs PhD)  

**Agent:** Reads files, runs commands, verifies work  
**Chatbot:** Text in, text out (copy-paste workflow)  

**Fill-in-the-Blank:** Scaffold + TODOs → agent completes  
**Anti-Slop Session:** Cleanup after agent work  
**Slop Zone:** Agent making dumb mistakes repeatedly  

---

## What Works

✅ Give agent way to verify (highest leverage!)  
✅ Rock-solid test suite  
✅ Test in same context (finds bugs automatically)  
✅ Context management (/clear, /compact, /btw)  
✅ Harness engineering (fix once, prevent forever)  
✅ Final manual review (always!)  
✅ Cleanup sessions (forces understanding)  
✅ Fill-in-the-blank pattern (scaffold + TODOs)  
✅ Cross-reference projects ("look at ../other-project")  
✅ "Fresh eyes" self-review (magic phrase)  
✅ CLIs > MCPs (less context pollution)  
✅ RTK everywhere (60-90% token savings)  

---

## What Doesn't Work

❌ Orchestrators (complexity theater)  
❌ Plan mode theater (over-planning)  
❌ RAG systems for coding (agent can just read files)  
❌ Multi-terminal workflows (coordination overhead)  
❌ Elaborate prompts (models got better)  
❌ Trust without verification (ship without review)  
❌ Let context accumulate (performance degrades)  
❌ Bash head against wall (if failing, back out)  

---

## Economics

**Cost:** $15.98 for non-trivial feature (less than coffee)  
**Time:** Not about faster/slower, about parallel work  
**Compensation:** $300k-$900k at top companies  

**Market:** 5,000 LLM researchers vs 50M software engineers  
**Prediction:** 10:1 ML Engineer jobs vs AI Engineer will invert in 5 years  

**Philosophy:** "When shipping AI products, you want engineers, not researchers."

---

## Evolution of AI Dev

**Phase 1:** Chatbots (inefficient, copy-paste)  
**Phase 2:** Agents (read files, run commands)  
**Phase 3:** Harness Engineering (prevent mistakes forever)  

**Prompt evolution:** Elaborate → Short (models got better)  
**Workflow evolution:** Waterfall ML → Agile AI  

**"AI is Agile":** Fire, ready, aim (vs ready, aim, fire)

---

## Mitchell's 6-Step Journey

1. **Drop the Chatbot** (use agents, not chat)
2. **Reproduce Your Own Work** (build expertise)
3. **End-of-Day Agents** (warm start next morning)
4. **Outsource Slam Dunks** (high-confidence tasks)
5. **Engineer the Harness** (prevent mistakes forever)
6. **Always Have Agent Running** (10-20% of workday)

**Each step builds on previous.** Can't skip.

---

## Critical Success Factors

### #1: Give Agent Way to Verify
- Tests (run and pass)
- Screenshots (compare)
- Benchmarks (improve metric)
- Expected outputs (diff)
- CLI results (known values)

**Impact:** Dramatically better performance.

### #2: Context Management
- /clear often (between tasks, after 2 corrections)
- /compact when needed (long sessions)
- /btw for quick questions (no history)
- Watch for "slop zone"

**Impact:** Maintain performance quality.

### #3: Final Manual Review
- Read every line
- Understand what it does
- Verify correctness
- Check security
- Polish if needed

**Impact:** Code quality, understanding, no regressions.

### #4: Harness Engineering
- Update AGENTS.md (wrong commands/APIs)
- Add tools (screenshots, tests)
- Fix once, never again

**Impact:** Accumulates over time, agent gets better.

### #5: Rock-Solid Test Suite
- Write tests in same context
- Agent runs tests
- Failures → agent fixes
- Green → ship

**Impact:** Agents fly with good tests.

---

## When to Use AI

### High Success Tasks
- Fill-in-the-blank (scaffold + TODOs)
- "I broke things, fix my mess"
- Add documentation
- Generate tests/simulations
- UI polish (spacing, colors)
- Rename refactors
- Move functions around
- Delete dead code

### Low Success Tasks
- Architecture decisions
- System design
- Database schema
- Picking dependencies
- Novel algorithms
- Security-critical code (review extra carefully!)

### Perfect for AI
- Inspiration (sometimes throw away)
- Prototyping (directionally good)
- Research (survey landscape)
- Cleanup (anti-slop sessions)
- Verification (run tests, compare screenshots)

---

## Common Failure Modes

### Slop Zone
**Signs:** Same error 3+ times, increasingly worse fixes, token spend accelerating.  
**Response:** Stop, /clear, back out, educate yourself.

### Context Pollution
**Signs:** Session feels "slow," agent seems confused, dumb mistakes.  
**Response:** /clear immediately, start fresh with better prompt.

### Hitting a Wall
**Signs:** Agent failing repeatedly, can't figure it out.  
**Response:** 2-3 attempts → back out → educate yourself → pivot approach.

### Trust Without Verification
**Signs:** Shipping without review, not understanding code.  
**Response:** Always final manual review. Always understand what agent did.

---

## Software 1.0 → 2.0 → 3.0

**Software 1.0:** Human-written code (Python, JavaScript, Go)  
**Software 2.0:** Neural networks (learned from data)  
**Software 3.0:** 1.0 + 2.0 + natural language  

**"Hottest new programming language is English."** - Andrej Karpathy

---

## Compensation Tiers

**AI Engineer (ships products):** $300k-$900k  
**Prompt Engineer:** $300k (Anthropic)  
**ML Engineer (trains models):** $200k-$400k  

**No PhD needed for AI Engineer.**  
**50M potential candidates (all software engineers).**

---

## Where to Start

1. **Read:** AGENTS.md (core rules)
2. **Read:** QUICK-REFERENCE.md (daily workflow)
3. **Read:** GLOSSARY.md (terms + patterns)
4. **Read:** RESEARCH-SYNTHESIS.md (deep details)

**Then:**
1. Use agents (not chatbots)
2. Give agent way to verify
3. Manage context aggressively
4. Final manual review always
5. Build harness over time

**Philosophy:** Just talk to it. One terminal, one agent, just ship.

---

## Key Quotes

> "AI tools amplify existing senior engineering expertise." - Simon Willison

> "When shipping AI products, you want engineers, not researchers." - Swyx

> "The hottest new programming language is English." - Andrej Karpathy

> "Give agent a way to verify = single highest leverage thing." - Anthropic

> "AI is Agile." - Mitchell Hashimoto

> "Plan mode is just a prompt." - Armin Ronacher

> "Not about faster/slower, about parallel work." - Mitchell Hashimoto

> "Fix once, never happens again." - Harness Engineering philosophy

---

## Resources

**Blogs:**
- Peter Steinberger: steipete.me
- Mitchell Hashimoto: mitchellh.com/writing
- Simon Willison: simonwillison.net
- Andrej Karpathy: karpathy.ai

**Podcasts:**
- Latent Space (Swyx + Alessio)
- Dwarkesh (with Karpathy)

**Tools:**
- Claude Code (Anthropic)
- OpenAI Codex
- Amp Code
- Cursor (with agentic mode)

**This repo:**
- ~/work/agent/ (centralized configs)
- 32 skills (shared across stacks)
- RTK.md (token optimization)

---

**Last updated:** May 2026  
**Based on research from:** Peter Steinberger, Mitchell Hashimoto, Simon Willison, Andrej Karpathy, Swyx, Anthropic
