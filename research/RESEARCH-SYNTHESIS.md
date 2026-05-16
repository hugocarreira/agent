# Research Synthesis - Modern AI Dev Practices (May 2026)

## Sources Deep Dive

### 1. Mitchell Hashimoto (Ghostty creator, ex-HashiCorp)

#### "My AI Adoption Journey" (Feb 2026)

**6-Step Evolution:**

**Step 1: Drop the Chatbot**
- Chatbots (ChatGPT web) are inefficient for coding
- MUST use agents (read files, execute programs, HTTP requests)
- Copy-paste workflow is a waste of time

**Step 2: Reproduce Your Own Work**
- Forced himself to do work twice (manual + agent)
- Excruciating but built expertise
- Discovered from first principles:
  - Break into clear, actionable tasks
  - Split vague requests into planning vs execution
  - Give agent way to verify = it fixes own mistakes

**Step 3: End-of-Day Agents**
- Last 30min of day: kick off 1+ agents
- Deep research sessions (survey libraries, pros/cons, activity)
- Parallel agents trying vague ideas (illuminate unknowns)
- Issue/PR triage (reports next morning for high-value tasks)
- "Warm start" next morning

**Step 4: Outsource the Slam Dunks**
- High-confidence tasks → agent does while you work on something else
- Turn off desktop notifications (context switching is expensive)
- Check during natural breaks
- Focus on tasks you love, delegate tasks you don't

**Step 5: Engineer the Harness**
- "Harness engineering" = prevent agent mistakes permanently
- Two forms:
  1. Better AGENTS.md (wrong commands, wrong APIs)
  2. Programmed tools (screenshots, filtered tests)
- Each bad behavior → fix once, never happens again

**Step 6: Always Have an Agent Running**
- Goal: agent running at all times
- Currently 10-20% of workday (working to improve)
- Not multiple agents (one is enough balance)
- Only run when truly helpful task exists

**Key Insight:** "AI is Agile" - fire, ready, aim workflow vs traditional waterfall ML

#### "Vibing a Non-Trivial Ghostty Feature" (Oct 2025)

Shipped macOS unobtrusive updates feature ($15.98 in tokens, ~8 hours wall clock).

**Process:**

1. **Pre-AI Planning** (human!)
   - Rough idea of backend (Sparkle framework)
   - Vague idea of frontend (titlebar accessory)
   - Enough to get started

2. **First Session: Prototype UI**
   - Only create plan first (not code!)
   - Agent created agreeable plan
   - UI directionally very good (lots of polish issues)
   - Used for inspiration
   - Hitting a wall → back out, educate myself

3. **Cleanup Sessions** (anti-slop)
   - Move methods, add documentation
   - Agents do better with documented code
   - Forces understanding before accepting

4. **Facing The Bug**
   - Multiple sessions trying different approaches
   - Started vague, got specific
   - All failed
   - Pivoted to different approach

5. **Fill-in-the-Blank Pattern**
   - Created scaffold with incomplete functions + TODO comments
   - Agent completes the owl
   - Works very well

6. **Cleanup Again, Big Time**
   - Restructured view model manually (better foundation)
   - Small manual work → sets up agents for success
   - Marathon cleanup sessions

7. **Simulation**
   - Agents great at generating tests/simulations
   - Quality doesn't matter for debug code

8. **Last Mile**
   - Hook everything up
   - Minor improvements

9. **Anything Else?**
   - Always ask "what else might I be missing?"
   - Do this even for manual code
   - Agent highlights real issues

**Patterns:**
- "I broke a bunch of things, please fix my mess"
- AI for inspiration (throw away, redo manually)
- Better organized code → better future sessions
- Final manual review SUPER important

**Philosophy:**
- AI can work while you do other things (breakfast photo!)
- Not about faster/slower, about parallel work
- Works for him personally

### 2. Swyx (AI Engineer Summit founder)

#### "The Rise of the AI Engineer" (Jun 2023)

**Definition:** Software engineers who specialize in AI applications, wielding the emerging stack.

**Why now:**
- Foundation models are few-shot learners
- In-context learning beyond original intent
- People who aren't LLM researchers find capabilities
- Microsoft/Google/Meta cornered research talent → "AI Research as a Service"
- 5,000 LLM researchers vs 50M software engineers
- GPU hoarding (Stability AI 4,000 GPUs)
- Fire, ready, aim (vs traditional ML waterfall)
- Python + JavaScript (TAM expansion 100%+)
- Generative AI vs Classifier ML

**Software 1.0 + 2.0 = 3.0:**
- Code atop intelligence vs intelligent software
- Prompt Engineering overhyped AND here to stay
- Re-emergence of Software 1.0 paradigms
- LangChain >$200M behemoth
- Code Core vs LLM Core applications

**AI Engineers are making:**
- $300k/yr prompt engineering (Anthropic)
- $900k building software (OpenAI)

**Not a single PhD in sight.**

### 3. Simon Willison (Datasette creator)

#### "Vibe Engineering" (Oct 2025)

**Definition:** Skilled engineers using LLMs while staying accountable.

**What LLMs amplify (senior engineering practices):**
- Automated testing (rock-solid test suite = agents fly)
- Planning in advance (iterate on plan, hand to agent)
- Comprehensive documentation (agent uses without reading code)
- Good version control (understand changes, git bisect)
- Effective automation (CI/CD, linting, preview envs)
- Code review culture
- Manual QA skills
- Research skills
- Ability to ship to preview environment

**What LLMs DON'T amplify:**
- Lack of tests
- Poor documentation
- Bad git habits
- No automation
- Orchestration theater

**Key insight:** "AI tools amplify existing expertise."

#### Other Insights:
- Parallel agents (3-8 instances) for research, maintenance, features
- YOLO mode critical for productivity (but dangerous!)
- Designing agentic loops (clear success + trial/error)
- Context window management critical

### 4. Andrej Karpathy (ex-OpenAI, ex-Tesla)

**Key statements:**
- "Hottest new programming language is English"
- "In numbers, significantly more AI Engineers than ML engineers"
- "One can be quite successful without ever training anything"
- Don't start AI Engineering by reading "Attention is All You Need"

**Focus:** Educational content on YouTube (Zero to Hero series)

### 5. Linus Lee (Notion AI researcher)

**Research areas:**
- Knowledge representation
- Creative work aided by machine understanding
- Software interfaces for clearer thinking
- Future of language and writing systems

**Approach:** Build interfaces that expand domain of thoughts we can think

### 6. Anthropic Claude Code Best Practices

**Most important:** Context window fills fast, performance degrades.

**Key practices:**
1. **Give Claude way to verify** (highest leverage!)
   - Tests, screenshots, expected outputs
   - Agent checks itself = dramatically better

2. **Explore → Plan → Code → Commit**
   - Plan mode for separation
   - But don't over-plan

3. **Provide specific context**
   - Reference files, mention constraints
   - Point to example patterns

4. **AGENTS.md must be SHORT**
   - If too long → agent ignores
   - Only what agent can't infer

5. **Context management**
   - /clear between tasks
   - /compact to summarize
   - /btw for quick questions
   - After 2 corrections → /clear + better prompt

6. **Course correct early**
   - ESC to stop
   - Don't let bad context accumulate

7. **Fresh eyes review**
   - "Review with fresh eyes" magic phrase
   - Seems to reset bias

### 7. Hamel Husain (Fast.ai)

(Blog redirect, couldn't fetch - need to search specific posts)

### 8. Latent Space Podcast Insights

**AI Engineer vs ML Engineer:**
- ML: train models, feature stores, fraud detection
- AI: writing apps, learning tools, natural language interfaces

**Fire, ready, aim:**
- Prompt LLM → build/validate → get data to finetune
- 1,000-10,000x cheaper validation
- Agile vs Waterfall

## Synthesis: What Actually Works

### Mindset
1. **Expect one-shot success** (models are that good now)
2. **Just talk to it** (short prompts, natural conversation)
3. **Give agent way to verify** (HIGHEST LEVERAGE!)
4. **Context is precious** (manage aggressively)
5. **AI amplifies existing senior engineering practices**

### Workflow
1. **Single agent, single terminal** (Mitchell's preference)
2. **Iterate fast, refactor smart** (80/20 rule)
3. **Test in same context** (finds bugs automatically)
4. **Course correct early** (ESC, /clear after 2 fails)
5. **Always have an agent running** (parallel work while you do other things)

### Prompting
1. **Short + screenshots** (~50% should include images)
2. **"Fresh eyes" for self-review** (magic phrase)
3. **Fill-in-the-blank pattern** (scaffold + TODO → agent completes)
4. **"I broke things, fix my mess"** (common pattern)
5. **Plan first for vague requests** (don't jump to code)

### Code Quality
1. **Cleanup sessions** (anti-slop, forces understanding)
2. **Documentation** (helps future agents + humans)
3. **Harness engineering** (fix mistakes once, prevent forever)
4. **Final manual review** (SUPER important, never skip)
5. **Better foundations** (good structure → better agent results)

### Context Management
1. **/clear between unrelated tasks**
2. **/compact to summarize** (keeps decisions, frees space)
3. **/btw for quick questions** (no history pollution)
4. **After 2 corrections → /clear** (polluted context)
5. **Subagents for investigation** (separate context) - optional

### What NOT to Do
❌ Orchestrators
❌ Plan mode theater (just write to docs/)
❌ RAG systems
❌ Multi-terminal workflows (unless you want them)
❌ Elaborate prompts
❌ Over-planning
❌ Trust without verification
❌ Let bad context accumulate

## Key Terms Defined

**Vibe Coding:** Fast, loose, irresponsible AI-driven development. No attention to how code works.

**Vibe Engineering:** Senior engineers using AI while staying accountable. Amplifies expertise.

**Agentic Engineering:** Same as vibe engineering. Building software with agents.

**Harness Engineering:** Engineering solutions to prevent agent mistakes permanently.

**Non-Trivial Vibing:** Mitchell's term for using AI on real, shipping features (not toys).

**Agents vs Chatbots:** Agents can read files, execute programs, make HTTP requests. Chatbots can't.

**AI Engineer vs ML Engineer:**
- ML: Train models, research, GPUs
- AI: Ship products using APIs/OSS models

**Code Core vs LLM Core:**
- Code Core: Human code orchestrates LLM (LangChain)
- LLM Core: LLM orchestrates code (ChatGPT plugins)

## Cost/Time Realities

- Mitchell: $15.98 tokens for non-trivial feature (~8 hours)
- Less than coffee during same timeframe
- Not about faster/slower, about parallel work
- Can work while doing other things

## Who's Doing This

**Companies:** Notion, Replit, Vercel, Amplitude, Figma
**Independent:** Simon Willison, Pieter Levels, Riley Goodside
**Compensation:** $300k-$900k at top companies
**Scale:** Millions using AI-built products overnight

## Future

- 10x more AI Engineer jobs than ML Engineer (predicted 5 years)
- No PhD needed
- "When it comes to shipping AI products, you want engineers, not researchers"
- Software 3.0 = 1.0 (code) + 2.0 (ML) + natural language

---

Based on research from:
- Mitchell Hashimoto (Ghostty)
- Swyx (Latent Space)
- Simon Willison (Datasette)
- Andrej Karpathy (ex-OpenAI/Tesla)
- Anthropic (Claude team)
- Linus Lee (Notion)
