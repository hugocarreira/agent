# Harness Principles

`agentrc` is built around a small set of repo-local harness principles that should hold across Codex, Claude, OpenCode, and similar coding agents.

## Short Instructions Beat Giant Manuals

- Keep root `AGENTS.md` map-like.
- Put durable detail in versioned docs and templates.
- Do not spend the context window on a 1,000-line instruction file.

## Repo Files Are the System of Record

- Important context should live in the repository, not only in chat history.
- Project scope belongs in `FEATURES.yaml`.
- Session continuity belongs in `PROGRESS.md`.

## State Must Survive Session Boundaries

- Long tasks span many prompts and often many agent sessions.
- Every substantial session should leave a restart point.
- The next agent should not have to rediscover what just happened.

## Verification Is a Harness Primitive

- Good prompts are not enough.
- Agents work best when success can be checked externally.
- Preferred signals:
  - tests
  - screenshots
  - benchmarks
  - expected output diffs
  - concrete manual acceptance checks

## Karpathy Rules Apply Everywhere

- Think before coding.
- Simplicity first.
- Surgical changes.
- Goal-driven execution.

These are not style preferences. They are anti-slop constraints.

## Substantial vs Trivial Sessions

Trivial edits:
- rename
- comment cleanup
- tiny copy fix

Substantial work:
- feature implementation
- multi-file refactor
- architectural change
- bug fix with diagnosis and verification

Only substantial sessions should update `PROGRESS.md`. `FEATURES.yaml` changes only when state, scope, or acceptance changes.

## Skills Are Secondary

- The harness should improve behavior by default.
- Skills are optional deeper workflows.
- If a project depends on a skill to behave sensibly, the base harness is too weak.
