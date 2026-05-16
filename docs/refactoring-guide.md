---
summary: "Systematic approach to refactoring sessions"
read_when: "user asks to refactor, clean up, or improve code quality"
---

# Refactoring Guide

## When to Refactor

Signs it's time:
- Prompts taking longer than expected
- Code duplication visible in stream
- Tests slow or flaky
- Dependencies outdated
- Architecture feels wrong

## Refactoring Session (~20% of time)

### 1. Code Duplication
```bash
# Find duplicates (if tool available)
npx jscpd .
```
- Extract to shared utilities
- Create reusable components
- DRY up similar logic

### 2. Dead Code
```bash  
# Find unused exports (if tool available)
npx knip
```
- Remove unused imports
- Delete commented code
- Remove unused functions/components

### 3. Modern Patterns

**React**:
- Reduce unnecessary `useEffect`
- Use React Compiler patterns
- Consolidate state
- Extract custom hooks

**TypeScript**:
- Improve type safety
- Remove `any` types
- Add proper interfaces

**Go**:
- Use standard library
- Proper error handling
- Consistent naming

### 4. Performance
- Profile slow operations
- Optimize database queries
- Add caching where appropriate
- Lazy load when possible

### 5. Tests
- Fix flaky tests
- Add missing test coverage
- Speed up slow tests
- Remove redundant tests

### 6. Documentation
- Update outdated docs
- Add missing code comments
- Improve README
- Update API docs

## Refactoring Strategy

**Small, Focused Commits**
Each refactor type gets its own commit:
```bash
git commit -m "refactor: remove code duplication in auth module"
git commit -m "refactor: update to modern React patterns"
git commit -m "chore: remove dead code"
```

**Test After Each Change**
- Run tests after each commit
- Verify app still works
- Fix immediately if broken

**Document Tricky Decisions**
Add inline comments explaining:
- Why not the obvious approach
- Performance considerations  
- Bug fixes that aren't obvious

## Tools by Language

### TypeScript/JavaScript
- `eslint` - linting
- `jscpd` - duplication
- `knip` - dead code
- `ast-grep` - structural search

### Go  
- `go vet` - issues
- `staticcheck` - static analysis
- `golangci-lint` - comprehensive

### Python
- `ruff` - fast linter
- `mypy` - type checking
- `bandit` - security

## Common Refactoring Patterns

### Extract Function
Before:
```typescript
function processUser(user) {
  // 50 lines of logic
}
```

After:
```typescript
function processUser(user) {
  const validated = validateUser(user);
  const enriched = enrichUserData(validated);
  return saveUser(enriched);
}
```

### Consolidate Conditionals
Before:
```typescript
if (user.age > 18) return true;
if (user.hasPermission) return true;
if (user.isAdmin) return true;
return false;
```

After:
```typescript
return user.age > 18 || user.hasPermission || user.isAdmin;
```

### Extract Component (React)
Before:
```tsx
function Dashboard() {
  return (
    <div>
      <div className="stats">
        {/* 100 lines of stats UI */}
      </div>
      <div className="charts">
        {/* 100 lines of charts UI */}
      </div>
    </div>
  );
}
```

After:
```tsx
function Dashboard() {
  return (
    <div>
      <StatsPanel />
      <ChartsPanel />
    </div>
  );
}
```

## Anti-Patterns to Avoid

❌ **Don't**: Refactor everything at once
✅ **Do**: Refactor one area at a time

❌ **Don't**: Change behavior during refactor
✅ **Do**: Keep behavior identical, just improve structure

❌ **Don't**: Refactor without tests
✅ **Do**: Add tests first if missing

❌ **Don't**: Over-abstract prematurely
✅ **Do**: Extract when pattern repeats 3+ times

## After Refactoring

- [ ] All tests passing
- [ ] App works identically  
- [ ] Code more readable
- [ ] Performance same or better
- [ ] Docs updated
- [ ] Commits pushed

---

**Remember**: Refactoring is not changing functionality. It's making code better while keeping behavior identical.
