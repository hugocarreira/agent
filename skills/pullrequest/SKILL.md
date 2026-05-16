---
name: pullrequest
description: If currently on main, create a new branch. Otherwise keep the current branch. Commit each modified file separately with meaningful commit messages, push the branch, and open a Pull Request using GitHub CLI and the repository's PULL_REQUEST_TEMPLATE.md.
compatibility: opencode
metadata:
  audience: developers
  tools: git, github-cli
  workflow: branch, commit, pull-request
  tags: ["tools", "git", "workflow"]
---

# Pull Request Skill (GitHub CLI Workflow)

## Purpose

This skill standardizes a clean Git workflow:

- If currently on `main`, create a new branch from it.
- If already on another branch, keep it and proceed.
- Commit each modified file separately with meaningful commit messages.
- One file per commit.
- Push the branch.
- Open a Pull Request using GitHub CLI (`gh`).
- Use the repository’s `PULL_REQUEST_TEMPLATE.md` (when available).

---

## Requirements

- Git installed
- GitHub CLI installed (`gh`)
- Authenticated session:

```bash
gh auth login
