# Agent Skills

This directory contains 26 high-quality, specialized skills that extend agent capabilities for common development, design, and marketing tasks.

Skills auto-trigger based on keywords in your prompts. You can also invoke them explicitly or discover new ones with `/skills`.

---

## Dev Workflow (6 skills)

| Skill | Purpose | Triggers |
|-------|---------|----------|
| **ship** | Complete shipping workflow: branch (if main), commit files separately, push, PR | "ship", "ship this", "ship it", "ready to ship" |
| **agent-browser** | Browser automation CLI for testing, scraping, form-filling | "open website", "fill form", "take screenshot", "scrape data" |
| **commit-and-push** | Commit changes and push to remote | "commit and push", "commit changes" |
| **pullrequest** | Create PR using GitHub CLI | "create PR", "pull request", "open PR" |
| **test-and-verify** | Write tests and verify implementation works | "write tests", "test and verify" |
| **simplify** | Simplify code for clarity without changing functionality | "simplify", "refine code" |

---

## Frontend & Design (7 skills)

| Skill | Purpose | Triggers |
|-------|---------|----------|
| **frontend-design** | Create distinctive, production-grade frontend interfaces | "build component", "create page", "frontend design" |
| **accessibility** | Audit and improve web accessibility (WCAG 2.2) | "accessibility", "a11y audit", "WCAG compliance", "screen reader" |
| **web-design-guidelines** | Review UI code for Web Interface Guidelines compliance | "review UI", "check design", "audit design" |
| **vercel-react-view-transitions** | React View Transitions API implementation | "view transitions", "page transitions", "animate route", "shared element" |
| **best-practices** | Modern web dev best practices (security, compatibility) | "best practices", "security audit", "code quality review" |
| **security-best-practices** | Language/framework-specific security reviews | "security review", "security best practices", "secure code" |
| **next-best-practices** | Next.js patterns (RSC, async, caching, metadata) | "Next.js best practices", "Next patterns", "RSC" |

---

## Next.js (2 skills)

| Skill | Purpose | Triggers |
|-------|---------|----------|
| **next-cache-components** | Next.js 16 Cache Components (PPR, use cache directive) | "cache components", "PPR", "use cache", "cacheLife" |
| **next-upgrade** | Upgrade Next.js following official migration guides | "upgrade Next.js", "migrate Next.js", "Next.js version" |

---

## SEO & Marketing (6 skills)

| Skill | Purpose | Triggers |
|-------|---------|----------|
| **seo** | Technical SEO optimization (meta tags, structured data, sitemaps) | "SEO", "optimize for search", "meta tags", "structured data" |
| **programmatic-seo** | Build SEO pages at scale using templates | "programmatic SEO", "pages at scale", "template pages", "pSEO" |
| **competitor-alternatives** | Create competitor comparison pages | "alternative page", "vs page", "competitor comparison" |
| **product-marketing-context** | Create/maintain product marketing context document | "product context", "marketing context", "positioning", "ICP" |
| **landing-page** | Landing page structure and conversion strategy | "landing page", "above-the-fold", "conversion page" |
| **signup-flow-cro** | Optimize signup/registration flows | "signup conversion", "registration flow", "signup form" |

---

## Growth & CRO (1 skill)

| Skill | Purpose | Triggers |
|-------|---------|----------|
| **onboarding-cro** | Post-signup onboarding optimization | "onboarding flow", "user activation", "first-run experience" |

---

## Performance (2 skills)

| Skill | Purpose | Triggers |
|-------|---------|----------|
| **core-web-vitals** | Optimize LCP, INP, CLS for page experience | "Core Web Vitals", "LCP", "INP", "CLS", "page experience" |
| **performance** | Web performance optimization (loading, runtime, resources) | "performance", "speed up", "optimize loading", "page speed" |

---

## Utilities (2 skills)

| Skill | Purpose | Triggers |
|-------|---------|----------|
| **changelog** | Create or update CHANGELOG.md following Keep a Changelog format | "changelog", "update changelog", "release notes", "prepare release" |
| **find-skills** | Discover and install agent skills from ecosystem | "find skill", "is there a skill", "skill for X" |

---

## Reference Documents

The `_references/` directory contains 5 knowledge base documents used by multiple skills:

| Document | Size | Used By |
|----------|------|---------|
| **ai-writing-detection.md** | 6.4K | seo (AI writing patterns to avoid) |
| **cro-experiments.md** | 7.7K | landing-page, onboarding-cro (CRO experiment library) |
| **marketing-psychology-models.md** | 19.4K | landing-page, signup-flow-cro (50+ mental models) |
| **monetization-models.md** | 2.7K | Product strategy (7 monetization frameworks) |
| **product-strategy-canvas.md** | 2.3K | Product planning (9-section canvas) |

---

## How Skills Work

### Auto-Trigger
Skills automatically activate when you use trigger keywords:
```
✅ "add accessibility to this page"  → triggers accessibility skill
✅ "create PR for this feature"      → triggers pullrequest skill
✅ "optimize Core Web Vitals"        → triggers core-web-vitals skill
```

### Discover Skills
```
/skills                  # List available skills
/skills <keyword>        # Search for skills
```

Or use the find-skills skill:
```
"is there a skill for browser automation?"
"find a skill for SEO"
```

### Explicit Invocation
You can explicitly invoke a skill by mentioning it:
```
"use the frontend-design skill to build a hero section"
"run accessibility skill on this component"
```

---

## Skill Quality

All skills in this directory have been:
- ✅ Reviewed for quality and usefulness
- ✅ Cleaned of generic/bloat content
- ✅ Enhanced with specific, actionable content
- ✅ Cross-referenced with related skills
- ✅ Integrated with reference documents

**No bloat. Only value.**

---

## Adding Skills

To add a new skill:
1. Create directory: `skills/skill-name/`
2. Add `SKILL.md` with skill definition
3. Follow existing skill structure (triggers, purpose, workflow)
4. Keep it focused and actionable
5. Extract generic content to `_references/` if needed

---

_Last updated: Skills cleanup (32 → 24), bloat removed, references extracted_
