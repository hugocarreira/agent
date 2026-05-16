# Agent Instructions

Work style: telegraph; noun-phrases ok; drop grammar; min tokens.
Codex CLI output: avoid Markdown tables by default; they render poorly. Use bullets or `key: value` lines. Tables only when explicitly requested.

## Core Principles

- Workspace: `~/work`. Missing repo: ask before cloning.
- "Make a note" here => terse `AGENTS.md` edit.
- Ship = changelog, commit, push when asked.
- Less is more. Don't over-engineer.

## Git Workflow

- Work on `main` by default unless project requires branches
- Atomic commits: one feature/fix per commit
- Commit message: Conventional Commits format (`feat|fix|refactor|build|ci|chore|docs|style|perf|test`)
- Only push when explicitly asked
- No `git worktree` unless asked - prefer simple main-based workflow

### Commit Helper

Use `git add [files] && git commit -m "message"` pattern.
Keep commits focused and atomic.

## Project Structure Preferences

### Language/Stack Defaults
- **TypeScript**: web stuff, full-stack apps
- **Go**: CLIs, APIs, performance-critical services  
- **Python**: data processing, ML, scripts when requested

### Standard Patterns
- Start with CLI, build UI later (easier to test)
- Keep deps minimal and well-maintained
- Check recent releases/commits/adoption before adding deps
- Use repo's package manager (don't swap without approval)

## Prompting Style

### What Works Best
- Short prompts (1-2 sentences) + screenshots
- Show don't tell: drag screenshot + "fix padding"
- Cross-reference: "look at ../other-project and do same"
- For big tasks: "give me options before making changes"

### Context Management  
- Read repo docs before coding
- Update docs + changelog for user-facing changes
- Add inline comments for tricky/bug-prone logic
- Keep docs in `docs/` folder with frontmatter

### Testing
- "After feature is done, write tests" (in same context!)
- Tests in same context almost always find bugs
- Not every UI tweak needs tests, but features/fixes do

## Code Quality

### When to Refactor
- When prompts start taking too long
- When you see ugly code in the stream
- Dedicate ~20% of time to refactoring
- Tools: `eslint`, code duplication checkers, dead code removers

### Inline Comments
- Brief notes for tricky logic
- Bug-prone areas
- Non-obvious decisions
- Previously buggy code

## Runtime Patterns

### CLIs > MCPs
Prefer CLIs with good help menus:
- `gh` for GitHub
- `vercel` for deployments  
- `psql` for database
- Project-specific CLIs when available

Only use MCPs when CLI doesn't exist and workflow is complex.

### Browser Testing
- Manual click-through preferred for UI work
- Automation only when repetitive
- Keep dev server running during iteration

## Common Tasks

### "ship" or "commit and push"
1. Update changelog if user-facing
2. Commit changed files with clear message
3. Push to remote
4. Confirm pushed successfully

### "refactor"
1. Check for code duplication
2. Remove dead code
3. Update to modern patterns
4. Add missing tests
5. Improve docs
6. Commit each type of change separately

### "review"
1. Read through recent changes
2. Check for bugs, edge cases, performance issues
3. Suggest improvements
4. Don't auto-fix unless asked

## Safety

- Never run destructive git ops without consent: `reset --hard`, `clean`, `restore`, force push
- Never expose secrets/keys
- Ask before major architectural changes
- Stop and ask if task scope grows significantly

## Workflow Philosophy

**Iterate Fast, Refactor Smart**
- Build → Test → Iterate → Ship
- Don't aim for perfect first try
- Improve incrementally
- Technical debt is ok if you pay it back regularly

**Blast Radius Awareness**  
Before starting, consider:
- How many files will this touch?
- How long should this take?
- Can this run in parallel with other work?
- If stuck, reassess approach

**Just Talk To It**
- Natural conversation > elaborate prompts
- Show with screenshots when possible
- Ask for options when uncertain
- Steer as you watch the stream

---

Last updated: 2026-02-14
