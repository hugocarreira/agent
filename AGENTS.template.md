# Project Agent Rules
#
# INSTRUCTIONS: Copy this file to your project as .opencode/AGENTS.md
# Fill in the sections marked with [PLACEHOLDER] — delete or leave empty what doesn't apply.
# This file extends the global ~/.config/opencode/AGENTS.md rules.
# Project-level rules take precedence over global rules for this project.

---

## Project Context

### What this project is
<!-- [PLACEHOLDER] One or two sentences describing the product/service -->
<!-- Example: A B2B SaaS platform for construction project management, targeting mid-size contractors. -->

### Tech stack
<!-- [PLACEHOLDER] List the exact technologies in use -->
- **Frontend**: <!-- e.g. Next.js 15, React 19, TypeScript 5.x, Tailwind CSS 4, shadcn/ui -->
- **Backend**: <!-- e.g. Node.js 22, Fastify, Prisma ORM, PostgreSQL 16 -->
- **Auth**: <!-- e.g. Clerk / NextAuth v5 / Supabase Auth -->
- **Database**: <!-- e.g. PostgreSQL on Supabase, Redis on Upstash -->
- **Deployment**: <!-- e.g. Vercel (frontend), Railway (backend), Cloudflare (DNS + CDN) -->
- **Testing**: <!-- e.g. Vitest, Playwright, MSW -->
- **Package manager**: <!-- e.g. pnpm / bun / npm -->
- **Monorepo tool**: <!-- e.g. Turborepo / nx / none -->

### Repository structure
<!-- [PLACEHOLDER] High-level layout of important directories -->
```
/
├── apps/
│   ├── web/          # Next.js frontend
│   └── api/          # Fastify backend
├── packages/
│   ├── ui/           # Shared component library
│   └── db/           # Prisma schema + migrations
└── tooling/          # Shared ESLint, TypeScript configs
```

### Environment
<!-- [PLACEHOLDER] Key environment variables agents must know about (names only, never values) -->
- `DATABASE_URL` — PostgreSQL connection string
- `NEXT_PUBLIC_API_URL` — Backend API base URL
- `AUTH_SECRET` — Auth signing secret

---

## Coding Conventions

### Naming
<!-- [PLACEHOLDER] Project-specific naming rules -->
- Components: PascalCase (`UserCard`, `InvoiceTable`)
- Hooks: camelCase with `use` prefix (`useInvoices`, `useCurrentUser`)
- API routes: kebab-case (`/api/v1/project-members`)
- Database tables: snake_case (`project_members`, `invoice_items`)
- CSS classes: follow Tailwind conventions; custom classes in `kebab-case`

### Import style
<!-- [PLACEHOLDER] -->
- Use `@/` alias for project-root imports
- Group imports: external → internal → relative → styles
- No default exports for utilities; named exports only

### Component patterns
<!-- [PLACEHOLDER] -->
- Server Components by default in Next.js; add `"use client"` only when needed
- Forms use `react-hook-form` + `zod`
- Data fetching uses TanStack Query on the client, `fetch` + `cache()` on the server
- API calls go through `src/lib/api.ts` — never `fetch` directly in components

### Error handling
<!-- [PLACEHOLDER] -->
- API errors use `{ error: { code: string, message: string, details?: unknown } }`
- Client errors use React Error Boundaries for section-level failures
- Log with `pino` (backend) — never `console.log` in production code

---

## Testing Requirements

<!-- [PLACEHOLDER] What must be tested and how -->
- Unit tests for all pure functions and hooks (`vitest`)
- Integration tests for all API endpoints (`supertest` or Fastify inject)
- E2E tests for critical user flows: signup, core feature, billing (`playwright`)
- No PRs merged with failing tests
- Coverage target: >80% for `packages/`, >60% for `apps/`

---

## Agent-Specific Rules for This Project

### All agents must:
- Read `package.json` and `tsconfig.json` before making any assumptions about versions
- Check existing component patterns in `packages/ui/src/` before creating new components
- Use `pnpm` (not npm or yarn) for all package operations
- Run `pnpm typecheck` after any TypeScript changes
- Never modify `packages/db/migrations/` — only add new migration files

### Frontend agent (@frontend-developer):
<!-- [PLACEHOLDER] Frontend-specific overrides -->
- App router only (`app/`) — no pages router
- Use `packages/ui` components before creating new ones
- Dark mode is supported — always test both themes
- i18n: all user-facing strings go through `next-intl` (`t('key')`)

### Backend agent (@backend-developer):
<!-- [PLACEHOLDER] Backend-specific overrides -->
- All database access via Prisma — no raw SQL unless absolutely necessary
- Every new endpoint needs a corresponding OpenAPI schema update in `src/openapi/`
- Rate limiting via `@fastify/rate-limit` — always apply to public endpoints
- Multi-tenant: every query must scope by `organizationId`

### DevOps agent (@devops-engineer):
<!-- [PLACEHOLDER] Infrastructure-specific rules -->
- CI: GitHub Actions workflows in `.github/workflows/`
- Deployments triggered by merge to `main` (prod) or `staging` branch (staging)
- Secrets managed via GitHub Secrets + Doppler
- Never change production Terraform state without a review step

### Designer agent (@designer-ui-ux):
<!-- [PLACEHOLDER] Design-specific rules -->
- Design tokens live in `packages/ui/src/tokens.css`
- Font: Inter (headings), Geist Mono (code)
- Primary color: `#6366F1` (indigo-500), accent: `#F59E0B` (amber-500)
- Figma file: <!-- link to Figma if exists -->

---

## Git & PR Conventions

<!-- [PLACEHOLDER] -->
- Branch naming: `feat/short-description`, `fix/short-description`, `chore/short-description`
- Commit style: Conventional Commits (`feat:`, `fix:`, `chore:`, `docs:`, `refactor:`)
- PRs require: passing CI, at least 1 approval, no unresolved comments
- Changelog auto-generated from conventional commits via `changelogen`
- Use `pullrequest` skill for all PRs — never push directly to `main`

---

## Out of Scope for Agents

<!-- [PLACEHOLDER] Things agents should NOT do in this project without explicit user approval -->
- Modifying billing/payment logic (Stripe webhooks, subscription state)
- Changing database schema without a migration file
- Upgrading major versions of Next.js, React, or Prisma
- Editing `.env.production` or production infrastructure configs
- Merging to `main` without a passing CI build

---

## Useful Commands

<!-- [PLACEHOLDER] Commands agents will need to run frequently -->
```bash
# Install dependencies
pnpm install

# Dev server (all apps)
pnpm dev

# Type check
pnpm typecheck

# Run tests
pnpm test

# Run E2E tests
pnpm test:e2e

# Lint + format
pnpm lint
pnpm format

# Database: generate Prisma client
pnpm db:generate

# Database: run migrations
pnpm db:migrate

# Build all
pnpm build
```
