# [Project Name]

**Core Principles:**
- Think Before Coding (don't assume, surface tradeoffs)
- Simplicity First (min code, no speculation)
- Surgical Changes (touch only what's needed)
- Goal-Driven (define success, loop until verified)

---

## Stack

Pick your stack or mix:

### Option 1: Go (Gin) + Next.js
```
Monorepo:
- apps/api: Go 1.25+, Gin, PostgreSQL, GORM
- apps/web: Next.js 16, Tailwind, TypeScript

Commands:
cd apps/api && make run       # API dev server
cd apps/web && npm run dev    # Frontend dev

Tests:
cd apps/api && make test
cd apps/web && npm test
```

### Option 2: Express + Vite/React
```
Monorepo:
- apps/api: Node 24, Express, Sequelize, PostgreSQL
- apps/web: Vite, React 19, Tailwind, Vitest

Commands:
npm run dev                   # Both API + web
npm run dev:api               # API only
npm run dev:web               # Web only

Tests:
npm run test                  # All tests
cd apps/api && npm test       # API tests
cd apps/web && npm test       # Web tests
```

### Option 3: Go API (Clean Architecture)
```
Standalone:
- Language: Go 1.25+
- Router: chi v5
- Structure: cmd/, internal/{entity,usecase,controller,repository,gateway}

Commands:
make run                      # Dev server with air (hot reload)
make test                     # All tests
make test-integration         # Integration tests
make test-e2e                 # E2E tests

Clean Architecture:
- entity: domain models
- usecase: business logic
- controller: HTTP handlers
- repository: data access
- gateway: external services
```

---

## Test & Build

```bash
# Development
rtk make run              # or npm run dev
rtk make test             # Run tests
rtk npm run lint          # Lint

# Build
rtk make build            # or npm run build
```

---

## Conventions

### Go Projects
- Package structure: `cmd/`, `internal/`, `pkg/`
- Functions: `PascalCase` for exported, `camelCase` for private
- Errors: return explicit errors, never panic in handlers
- Database: UUIDs for PKs, `snake_case` naming
- Tests: `*_test.go`, table-driven tests

### TypeScript/React
- Components: `PascalCase` (e.g., `LoginButton.tsx`)
- Functions/hooks: `camelCase` (e.g., `useAuth`)
- No `any` types (use `unknown` if needed)
- Tailwind tokens only (no hardcoded colors/fonts)

### API
- Routes: `/api/v1/...`
- Status codes: 200, 201, 400, 401, 404, 500
- JWT in `Authorization: Bearer <token>`

---

## Git Commits

```bash
# Format: <type>(<scope>): <description>
feat(api): add user registration
fix(web): prevent form double submit
chore(root): update dependencies

# Scopes: api, web, landing, root
# Types: feat, fix, chore, refactor, docs, test
```

Atomic commits (one logical change per commit).

---

## Out of Scope

```
Never modify:
- lib/billing/ (legacy, external team)
- .env files (never commit!)

Ask before:
- Database schema changes
- Breaking API changes
- Deleting files
```

---

## Project-Specific Notes

[Add anything agents keep getting wrong or can't infer from code]

Examples:
- Use bun, not npm
- API base URL: http://localhost:8080/api/v1
- Run migrations before schema changes
- Brand voice: casual PT-BR (see .guidelines/brand.md)

---

**Keep it simple. Add only what's needed. Update when agent keeps making same mistake.**
