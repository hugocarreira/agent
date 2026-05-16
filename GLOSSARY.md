# AI Dev Glossary - Modern Terms & Practices

## Key Terms

### Vibe Coding
**Definition:** Fast, loose, irresponsible AI-driven development. No attention to how code works.

**Characteristics:**
- Accept everything AI produces
- No review, no understanding
- Ship without testing
- "It works on my machine" mentality

**Result:** Technical debt, security issues, unmaintainable code.

---

### Vibe Engineering (Simon Willison)
**Definition:** Senior engineers using AI while staying accountable. Amplifies expertise.

**Characteristics:**
- Strong fundamentals first
- AI amplifies existing practices
- Final manual review always
- Understand what agent does
- Rock-solid test suite
- Good documentation
- Effective automation

**Key insight:** "AI tools amplify existing senior engineering expertise."

**Not the same as:** Vibe coding (irresponsible)

---

### Agentic Engineering
**Definition:** Building software with agents (LLMs that can read files, run commands, make HTTP requests).

**Same as:** Vibe engineering (when done responsibly)

**Components:**
- Agent (not chatbot!)
- Tools (file access, terminal, HTTP)
- Verification methods
- Harness engineering

---

### Harness Engineering (Mitchell Hashimoto)
**Definition:** Engineering solutions to prevent agent mistakes permanently.

**Forms:**
1. **AGENTS.md improvements** (implicit prompting)
   - Document wrong commands → agent stops using them
   - Document correct APIs → agent uses right ones
   - Each line based on bad behavior

2. **Programmed tools** (explicit verification)
   - Screenshot scripts
   - Filtered test runners
   - Benchmark tools
   - Diff helpers

**Philosophy:** Fix mistakes once, prevent forever.

**Example:** Agent keeps using wrong test command → add to AGENTS.md → never happens again.

---

### Non-Trivial Vibing (Mitchell Hashimoto)
**Definition:** Using AI on real, shipping features (not toy projects).

**Characteristics:**
- Production code
- Real users
- Non-trivial complexity
- Final manual review
- Harness engineering
- Understanding retained

**Example:** Mitchell's Ghostty update feature ($15.98 tokens, 8 hours, shipping to thousands).

**Not the same as:** Toy demos, throwaway prototypes.

---

### AI Engineer (Swyx)
**Definition:** Software engineer specializing in AI applications, wielding the emerging stack.

**Stack:**
- Foundation models (APIs or open source)
- Chaining tools (LangChain, LlamaIndex)
- Vector databases (Pinecone, Weaviate)
- Evaluation frameworks
- Deployment platforms

**Compensation:** $300k-$900k at top companies

**Background:** Usually software engineering (not ML/PhD)

**Products:** Writing apps, learning tools, natural language interfaces

**Not the same as:** ML Engineer (trains models, needs PhD often)

---

### ML Engineer
**Definition:** Trains, fine-tunes, deploys machine learning models.

**Focus:**
- Model training
- Feature engineering
- Data pipelines
- Model serving
- A/B testing
- Feature stores

**Products:** Fraud detection, recommendation systems, anomaly detection

**Background:** Usually CS/Math PhD or equivalent

**Not the same as:** AI Engineer (uses models, ships products)

---

### Software 1.0
**Definition:** Classical programming with human-written code.

**Characteristics:**
- Explicit logic
- Deterministic
- Debuggable
- Predictable

**Examples:** Python, JavaScript, Go, Rust

---

### Software 2.0 (Andrej Karpathy)
**Definition:** Neural networks that approximate logic instead of explicit code.

**Characteristics:**
- Learned from data
- Non-deterministic
- Black box
- Probabilistic

**Examples:** Image recognition, speech-to-text, playing Go

---

### Software 3.0
**Definition:** Software 1.0 (code) + Software 2.0 (ML) + natural language.

**Characteristics:**
- Human code orchestrates LLMs
- LLMs generate code
- Natural language as programming language
- Hybrid human-AI authorship

**Examples:** GitHub Copilot, ChatGPT Code Interpreter, Claude Code

---

### Agent vs Chatbot

**Chatbot:**
- Only text in, text out
- No external tools
- Can't verify work
- Copy-paste workflow

**Example:** ChatGPT web interface

**Agent:**
- Read files
- Execute programs
- Make HTTP requests
- **Can verify own work**
- Loop until success

**Example:** Claude Code, OpenAI Codex, Cursor with agentic mode

---

### Code Core vs LLM Core

**Code Core (LLM Shell):**
- Human code is primary
- LLM fills in parts
- Code orchestrates LLM

**Example:** LangChain apps

**LLM Core (Code Shell):**
- LLM is primary
- Code is tool LLM uses
- LLM orchestrates code

**Example:** ChatGPT plugins, Notion AI

---

### Prompt Engineering
**Definition:** Crafting text inputs to get desired outputs from LLMs.

**Techniques:**
- Clear instructions
- Examples (few-shot)
- Role prompting
- Chain of thought
- XML structuring
- System prompts

**Status:** Overhyped AND here to stay

**Evolution:** From elaborate prompts → shorter prompts (better models)

---

### Fill-in-the-Blank Pattern (Mitchell)
**Definition:** Create scaffold with incomplete functions/TODOs, agent completes.

**Process:**
1. Write function signatures
2. Add TODO comments
3. Add type hints
4. Ask agent to complete

**Why it works:**
- Clearer intent
- Constrained scope
- Better type inference
- Less hallucination

**Example:**
```python
def process_user_data(user_id: str) -> UserData:
    """Process user data and return structured result.
    
    TODO: 
    - Fetch from database
    - Validate required fields
    - Transform to UserData model
    """
    pass  # Agent fills this in
```

---

### Anti-Slop Session (Mitchell)
**Definition:** Cleanup pass after agent work to prevent low-quality code.

**Activities:**
- Move functions to better locations
- Add documentation
- Rename for clarity
- Restructure for maintainability
- Delete dead code

**Purpose:**
- Forces understanding
- Improves future agent sessions
- Maintains code quality

**When:** After every significant agent session

---

### Slop Zone
**Definition:** When agent enters loop of making dumb mistakes repeatedly.

**Signs:**
- Same error 3+ times
- Increasingly worse "fixes"
- Thrashing between approaches
- Token spend accelerating

**Response:**
- Stop immediately
- /clear
- Back out
- Educate yourself manually
- Come back with better understanding

---

### YOLO Mode (Simon Willison)
**Definition:** Auto-approve all agent actions without review.

**When safe:**
- Docker containers (restricted)
- GitHub Codespaces (not your machine)
- Sandboxed environments
- Network allowlisting

**When dangerous:**
- Your local machine
- Production systems
- Sensitive data present

**Philosophy:** "An AI agent is an LLM wrecking its environment in a loop."

---

### Oracle (Amp-specific)
**Definition:** Read-only subagent using slower, higher-cost model for thinking.

**When to consult:**
- Planning phase
- Complex decisions
- Architecture choices
- Unknown unknowns

**Cost:** Higher tokens, slower response

**Benefit:** Better reasoning, fewer mistakes

---

### Context Window Management
**Definition:** Actively managing LLM context to maintain performance.

**Why critical:** Performance degrades as context fills.

**Techniques:**
- /clear (reset completely)
- /compact (summarize, keep key info)
- /btw (quick questions, no history)
- Subagents (separate context)
- After 2 corrections → /clear

**Anti-patterns:**
- Let context fill up
- Keep retrying in polluted context
- Never /clear

---

## Workflow Patterns

### Explore → Plan → Code → Commit
**When:** Vague requirements or uncertain approach

**Process:**
1. **Explore:** Research, ask questions, gather info
2. **Plan:** Write spec to docs/plans/*.md
3. **Code:** Implement according to plan
4. **Commit:** Atomic, conventional commits

**Alternative:** Just code (when requirements clear)

---

### Reproduce Your Own Work (Mitchell)
**Definition:** Do work manually, then force agent to reproduce identical results.

**Purpose:** Build expertise, discover what works from first principles

**Process:**
1. Do task manually (full quality)
2. In separate session, agent does same
3. Compare, iterate until identical
4. Note what prompts/patterns worked

**Pain:** Excruciating (doing work twice)
**Gain:** Deep understanding of agent capabilities

---

### End-of-Day Agents (Mitchell)
**Definition:** Last 30min of day, kick off background agents.

**Tasks:**
- Deep research (survey libraries, pros/cons)
- Parallel vague ideas (illuminate unknowns)
- Issue/PR triage (reports for next morning)

**Benefit:** "Warm start" next morning

**Timing:** Complete in <30min (not overnight loops)

---

### Outsource the Slam Dunks (Mitchell)
**Definition:** High-confidence tasks → agent (background) while you work on something else.

**Requirements:**
- Very confident agent will succeed
- Can check during natural breaks
- **Desktop notifications OFF**

**Your job:** Work on tasks you love, manually

**Agent's job:** Tasks you don't love, high confidence

**Balance:** Skill formation tradeoff (deliberate choice)

---

### Always Have an Agent Running (Mitchell)
**Goal:** Agent running 10-20%+ of workday

**Current reality:** One agent (not multiple)

**Criteria:** Only when truly helpful task exists

**Not about:** Running agents for sake of running agents

---

### Hitting a Wall (Mitchell)
**Definition:** Agent failing repeatedly, can't figure it out.

**Response:**
1. Make 2-3 hail mary attempts (vague → specific)
2. If still failing:
   - Stop
   - Back out  
   - Educate yourself manually
   - Pivot to different approach
3. Don't bash head against wall

**Key insight:** AI isn't always the solution. Sometimes it's a liability.

---

## Cost & Economics

### Token Economics
- Mitchell's non-trivial feature: $15.98
- Less than coffee during same timeframe
- Not about faster/slower
- About parallel work

### Compensation
- AI Engineer: $300k-$900k
- Prompt Engineer: $300k (Anthropic)
- Software Engineer (AI): $900k (OpenAI)

### Market Size
- 5,000 LLM researchers globally
- 50M software engineers globally
- 10:1 ML Engineer jobs vs AI Engineer (today)
- Predicted inversion in 5 years

---

## Philosophy

### "AI is Agile" (Mitchell)
**Meaning:** Fire, ready, aim workflow vs traditional ML waterfall.

**Traditional ML:**
1. Collect data
2. Train model
3. Evaluate
4. Deploy
5. Discover product/market fit issues

**AI-First:**
1. Prompt LLM
2. Build/validate product
3. Get specific data to finetune (maybe)

**Result:** 1,000-10,000x cheaper validation

---

### "AI Amplifies Existing Expertise" (Simon)
**Meaning:** AI makes experts more expert. Doesn't fix juniors.

**What amplifies:**
- Automated testing → agents fly
- Good documentation → agents understand
- Version control → agents iterate
- Automation → agents leverage

**What doesn't amplify:**
- Lack of tests → agents flounder
- Poor docs → agents hallucinate
- Bad git → agents create mess
- No automation → agents can't verify

---

### "When Shipping AI Products, You Want Engineers Not Researchers" (Swyx)
**Meaning:** PhDs not needed for AI products.

**AI products built by:** Software engineers (no ML background)

**AI research done by:** PhDs (5,000 people)

**Scale difference:** 10,000x more engineers than researchers

---

### "Hottest New Programming Language is English" (Andrej Karpathy)
**Meaning:** Natural language as programming interface.

**Evolution:**
- Software 1.0: Human code
- Software 2.0: Neural networks
- Software 3.0: Natural language + code + ML

---

## Anti-Patterns

### Orchestrators
**What:** Multiple specialized agents with coordinator

**Why bad:**
- Complexity theater
- Context fragmentation
- Coordination overhead
- Rare actual value

**Alternative:** Just talk to one agent

---

### Plan Mode Theater
**What:** Elaborate planning phase before every task

**Why bad:**
- Over-planning
- Time waste on simple tasks
- "Plan mode is just a prompt"

**Alternative:**
- Plan only when YOU don't know what you want
- For clear tasks: just do it
- Write to docs/*.md if needed (no ceremony)

---

### RAG Systems (for agentic coding)
**What:** Vector database retrieval for context

**Why bad (for coding):**
- Agent can just read files
- Adds complexity
- Context pollution
- Rare improvement over simple file reads

**When useful:** Knowledge bases, chat over docs (not coding)

---

### Multi-Terminal Theater
**What:** Multiple terminal sessions, complex handoffs

**Why bad:**
- Coordination overhead
- Context switching
- Rare actual value
- Complexity for complexity's sake

**Alternative:** One terminal, one agent (Peter/Mitchell style)

---

### Elaborate Prompts
**What:** Long, detailed, formal prompts

**Why bad:**
- Models got better (shorter prompts work)
- More tokens = more cost
- Harder to iterate
- "Just talk to it"

**Alternative:**
- "add login button"
- "[screenshot] fix padding"
- 1-2 sentences + image

---

### Trust Without Verification
**What:** Ship agent code without review/testing

**Why bad:**
- Security issues
- Subtle bugs
- Technical debt
- Unmaintainable code

**Alternative:**
- Final manual review (always)
- Agent runs tests
- Cleanup sessions
- Understanding retained

---

## Success Factors

### Give Agent Way to Verify (HIGHEST LEVERAGE!)
**Forms:**
- Tests (run and pass)
- Screenshots (compare)
- Benchmarks (improve metric)
- Expected outputs (diff)
- CLI results (known values)

**Impact:** Dramatically better performance

**Why:** Agent fixes own mistakes, prevents regressions

---

### Rock-Solid Test Suite
**Impact:** Agents fly

**Why:** Can verify changes don't break things

**Process:**
- Write tests in same context
- Agent runs tests
- Failures → agent fixes
- Green → ship

---

### Context Management
**Impact:** Performance quality

**Why:** Context fills → performance degrades

**Techniques:**
- /clear often
- /compact when needed
- /btw for quick questions
- After 2 corrections → /clear

---

### Harness Engineering
**Impact:** Prevent repeated mistakes

**Process:**
1. Agent makes mistake
2. Update AGENTS.md or add tool
3. Mistake never happens again

**Accumulation:** Over time, agent gets better and better

---

### Final Manual Review
**Impact:** Code quality, understanding

**Process:**
- Read every line
- Understand what it does
- Verify correctness
- Check security
- Polish if needed

**Never skip.**

---

Based on research from:
- Mitchell Hashimoto (Ghostty)
- Swyx (Latent Space)
- Simon Willison (Datasette)
- Andrej Karpathy (ex-OpenAI/Tesla)
- Peter Steinberger (PSPDFKit)
- Anthropic (Claude team)
