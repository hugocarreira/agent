---
summary: "Checklist for shipping features to production"
read_when: "user asks to ship, deploy, or release"
---

# Shipping Checklist

## Pre-Ship

- [ ] Feature works as expected locally
- [ ] Tests written and passing (if applicable)
- [ ] Code reviewed (by AI or human)
- [ ] No console errors/warnings
- [ ] Performance acceptable (no obvious slowdowns)
- [ ] Mobile responsive (if web UI)
- [ ] Accessibility basics (keyboard nav, screen reader)

## Commit & Push

- [ ] Changelog updated (user-facing changes)
- [ ] Atomic commits with clear messages
- [ ] No secrets/keys in code
- [ ] Dependencies updated if needed
- [ ] Pushed to remote

## Deploy

- [ ] CI/CD green
- [ ] Staging tested (if available)
- [ ] Database migrations run (if applicable)
- [ ] Environment variables set
- [ ] Deploy to production

## Post-Deploy

- [ ] Verify in production
- [ ] Check logs for errors
- [ ] Monitor performance
- [ ] Update docs if needed
- [ ] Close related issues/PRs

## Rollback Plan

If something breaks:
1. Check logs immediately
2. Identify the issue
3. Quick fix if trivial
4. Revert deploy if complex
5. Fix in dev, redeploy

---

**Remember**: Shipping broken code is ok if you can fix it fast. Don't let perfect be the enemy of done.
