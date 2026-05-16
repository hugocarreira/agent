# Agent Scripts

Instruções compartilhadas, skills e helpers para otimizar desenvolvimento com AI agents.

Inspirado no workflow de [Peter Steinberger](https://steipete.me) para shipping at inference-speed.

## Estrutura

```
~/work/agent/
├── AGENTS.md          # Regras compartilhadas para todos os agents
├── skills/            # Skills reutilizáveis (workflows)
├── scripts/           # Helper scripts portáveis
├── docs/              # Documentação de processos
└── hooks/             # Git hooks e validações
```

## Setup Rápido

### 1. Symlinks Globais

```bash
# Para Codex
ln -sf ~/work/agent/AGENTS.md ~/.codex/AGENTS.md
ln -sf ~/work/agent/skills ~/.codex/skills

# Para Claude Code
ln -sf ~/work/agent/AGENTS.md ~/.claude/AGENTS.md
ln -sf ~/work/agent/AGENTS.md ~/.claude/CLAUDE.md
ln -sf ~/work/agent/skills ~/.claude/skills
```

### 2. Nos Seus Projetos

Adicione no `AGENTS.md` de cada projeto:

```markdown
READ ~/work/agent/AGENTS.md BEFORE ANYTHING (skip if missing).

# [Nome do Projeto] - Regras Específicas

## Stack
- Linguagem: TypeScript/Go/Python
- Framework: Next.js/Gin/FastAPI
- Database: PostgreSQL/MongoDB

## Patterns Preferidos
[suas regras específicas aqui]
```

## Filosofia Core

### Just Talk To It 🗣️
- Prompts curtos (1-2 sentenças) + screenshots
- Sem charadas elaboradas
- Conversação natural com o modelo

### Shipping at Inference-Speed ⚡
- Velocity limitada por inference time, não por coding
- Expectativa: código funciona out-of-the-box
- Watch the stream, don't read every line

### Blast Radius Thinking 💥
Antes de cada task:
- Quantos arquivos vai tocar?
- Quanto tempo deve levar?
- Pode rodar em paralelo?

## Workflow Típico

1. **Start com CLI** → Testa com agent → Build UI depois
2. **Iterate fast** (80%) → **Refactor** (20%)
3. **Test na mesma context** que implementou
4. **Commit atômico** após cada feature

## Tools Essenciais

- **Terminal**: Ghostty / Alacritty
- **Primary Agent**: GPT-5 Codex (high mode)
- **Review**: GPT-5-Pro via ChatGPT
- **CLIs > MCPs**: `gh`, `vercel`, `psql`, `bun`

## Skills vs Scripts

**Skills**: Workflows completos com contexto
- Vão em `skills/<name>/SKILL.md`
- Têm YAML frontmatter
- Agentes descobrem automaticamente

**Scripts**: Helpers portáveis e leves
- Vão em `scripts/`
- Zero dependências quando possível
- Podem ser chamados por skills

## Regras de Ouro

✅ **DO**
- Use screenshots em ~50% dos prompts
- Cross-reference outros projetos (`../outro-projeto`)
- Queue continue messages para long tasks
- Pense em blast radius antes de começar
- Refatore quando prompts ficam lentos

❌ **DON'T**
- Não leia todo o código (watch the stream)
- Não use RAG/subagents complicados
- Não waste context com MCPs pesados
- Não over-engineer os prompts
- Não tema rodar múltiplos agents em paralelo

## Manutenção

Sempre que descobrir um padrão útil:
1. Anote no AGENTS.md ou crie uma skill
2. Commit com mensagem clara
3. Agents futuros vão se beneficiar

---

**Próximo passo**: Edite `AGENTS.md` com suas preferências e comece a usar!
