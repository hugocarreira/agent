# Agent Skills

This directory contains 16 focused skills that complement the default-behavior harness.

These are secondary tools, not the main operating model. The harness should improve any session by default; skills are for deeper, domain-specific workflows when a task needs them.

---

## Core Workflow (5 skills)

| Skill | Purpose | Typical prompts |
|-------|---------|-----------------|
| **ship** | Branch if needed, commit by context, push, create or update PR | "ship", "commit and push", "create PR" |
| **agent-browser** | Browser automation for testing, scraping, screenshots, and form flows | "open website", "fill form", "take screenshot" |
| **test-and-verify** | Write tests and verify implementation in the same context | "write tests", "verify this works" |
| **changelog** | Create or update `CHANGELOG.md` from real changes | "update changelog", "release notes" |
| **find-skills** | Discover optional skills outside the core catalog | "find a skill for X", "is there a skill for..." |

---

## Frontend & App Quality (6 skills)

| Skill | Purpose | Typical prompts |
|-------|---------|-----------------|
| **frontend-design** | Build distinctive, production-grade frontend UI | "build component", "create page" |
| **accessibility** | Audit and improve a11y using WCAG 2.2 patterns | "accessibility", "WCAG", "screen reader" |
| **security-best-practices** | Language/framework-specific secure-by-default guidance | "security review", "secure code" |
| **performance** | Broad web performance optimization, including Core Web Vitals | "performance", "LCP", "INP", "CLS" |
| **next-best-practices** | Next.js architecture, async APIs, and Cache Components | "Next.js best practices", "use cache", "cacheTag" |
| **next-upgrade** | Upgrade Next.js with migration-guide discipline | "upgrade Next.js", "migrate Next.js" |

---

## Growth & Writing (5 skills)

| Skill | Purpose | Typical prompts |
|-------|---------|-----------------|
| **seo** | Technical and on-page SEO improvements | "SEO", "meta tags", "structured data" |
| **landing-page** | Landing page structure, messaging, and conversion guidance | "landing page", "above-the-fold" |
| **product-marketing-context** | Create or maintain the shared product/ICP/positioning doc | "product context", "ICP", "positioning" |
| **content-research-writer** | Research-backed long-form writing and drafting support | "research and write", "outline article" |
| **domain-name-brainstormer** | Brainstorm domain names and availability options | "suggest domain names", "brainstorm domains" |

---

## Reference Documents

The `_references/` directory contains shared knowledge used by some skills:

| Document | Used By |
|----------|---------|
| **ai-writing-detection.md** | seo |
| **cro-experiments.md** | landing-page |
| **marketing-psychology-models.md** | landing-page |
| **monetization-models.md** | Future optional strategy skills |
| **product-strategy-canvas.md** | Future optional strategy skills |

---

## Notes

- `ship` absorbs the old commit/push/PR split.
- `performance` covers Core Web Vitals work.
- `next-best-practices` covers Cache Components guidance.
- Long-tail or niche skills should live outside the core catalog and be discovered with `find-skills`.

---

_Last updated: focused catalog cleanup_
