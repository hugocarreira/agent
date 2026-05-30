# Agent Instructions

`agentrc` is a cross-agent harness repo. Optimize for portability, small context, and durable repo-local state.

## Core Rules
- Think before coding. Surface assumptions and tradeoffs.
- Simplicity first. Minimum code, no speculative flexibility.
- Surgical changes. Touch only what the task requires.
- Goal-driven execution. Define done conditions before substantial work.

## Verification
- Prefer tests, screenshots, benchmarks, or expected-output diffs.
- If a project has no tests, define concrete manual completion criteria.
- Leave the repo in a clearer state than you found it.

## State Updates
- For substantial sessions, update `PROGRESS.md` before finishing.
- Update `FEATURES.yaml` only when feature state, scope, or acceptance changes.
- Skip both files for trivial edits.

## Commands
- Prefix shell commands with `rtk`.
- Prefer CLIs over heavy MCP usage when local tooling can answer the question.

## Project Pack
Default `agentrc init` scaffolds:
- `AGENTS.md`
- `FEATURES.yaml`
- `PROGRESS.md`

## Project AGENTS.md
Each project-level `AGENTS.md` should stay short:
- core rules
- commands
- state file expectations
- out-of-scope boundaries
- facts agents cannot infer from code

## Out of Scope
- No orchestrators.
- No agent theater.
- No giant instruction manuals at repo root.

## See Also
- [SETUP.md](/home/hugoubuntu/work/agentrc/SETUP.md)
- [docs/harness-principles.md](/home/hugoubuntu/work/agentrc/docs/harness-principles.md)
- [skills/README.md](/home/hugoubuntu/work/agentrc/skills/README.md)
