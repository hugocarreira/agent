# Como Usar Este Setup

## 🎯 O Que É Isso?

Um ambiente centralizado para 3 harnesses de AI (opencode, codex, claude):
- Regras globais em `~/work/agent/AGENTS.md`
- 32 skills compartilhadas
- Token optimization (RTK)
- Templates para novos projetos

**Resultado:** Todos os harnesses seguem as mesmas regras, compartilham skills, funcionam igual.

---

## ⚡ Setup Inicial (Uma Vez Só)

### 1. Criar Symlinks (se ainda não fez)

```bash
cd ~/work/agent
./setup-symlinks.sh
```

**O que isso faz:**
```
~/.config/opencode/AGENTS.md  → ~/work/agent/stacks/opencode/AGENTS.md
~/.config/opencode/skills/    → ~/work/agent/skills/
~/.codex/AGENTS.md            → ~/work/agent/stacks/codex/AGENTS.md
~/.codex/skills/              → ~/work/agent/skills/
~/.claude/CLAUDE.md           → ~/work/agent/stacks/claude/CLAUDE.md
~/.claude/skills/             → ~/work/agent/skills/
```

### 2. Verificar

```bash
# Ver se symlinks foram criados
ls -la ~/.config/opencode/AGENTS.md
ls -la ~/.codex/AGENTS.md
ls -la ~/.claude/CLAUDE.md

# Todos devem apontar pra ~/work/agent/stacks/...
```

**Pronto!** Setup global completo. Agora configure cada projeto.

---

## 📁 Setup Por Projeto

### Opção 1: Projeto Novo

```bash
cd ~/work/meu-projeto-novo
cp ~/work/agent/AGENTS.template.md ./AGENTS.md

# Edita AGENTS.md:
# 1. Escolhe Option 1, 2 ou 3 (apaga o resto)
# 2. Adiciona comandos específicos
# 3. Adiciona "Out of Scope" (o que NÃO mexer)
# 4. MANTÉM CURTO! (só o que agent não consegue inferir do código)
```

### Opção 2: Projeto Existente

```bash
cd ~/work/meu-projeto-existente

# Se já tem AGENTS.md/CLAUDE.md:
# - Lê ~/work/agent/AGENTS.md (regras globais)
# - Adiciona linha no topo: "READ ~/work/agent/AGENTS.md BEFORE ANYTHING."
# - Remove duplicações (global já cobre)

# Se não tem:
cp ~/work/agent/AGENTS.template.md ./AGENTS.md
# Edita com stack + comandos específicos
```

---

## 🚀 Workflow Diário

### Como o Claude Lê os Arquivos

Quando você roda `claude` em qualquer projeto:

```
1. Claude lê: ~/.claude/CLAUDE.md
   ↓ (que é symlink pra ~/work/agent/stacks/claude/CLAUDE.md)
   ↓ (que tem: "READ ~/work/agent/AGENTS.md")

2. Claude lê: ~/work/agent/AGENTS.md (regras globais)
   ↓ (core rules, RTK, skills, etc)

3. Claude lê: ~/work/agent/RTK.md (token optimization)
   ↓

4. Claude lê: ./AGENTS.md (regras do projeto, se existir)
   ↓ (stack, comandos, out of scope)

5. Claude lê: seu código
```

**Ordem de prioridade:** Projeto > Global > Defaults

### Exemplo Real

```bash
cd ~/work/meu-app
claude

# Você:
"add login button"

# Claude entende:
# - AGENTS.md global: "Just talk to it, short prompts, give way to verify"
# - RTK.md: "Prefix shell commands with rtk"
# - ./AGENTS.md: "Stack: Next.js 16, Commands: npm run dev, Never touch: lib/billing/"
# - Código: vê que já existe Button component, reutiliza padrões

# Claude faz:
# 1. Cria LoginButton.tsx seguindo padrões existentes
# 2. Sugere teste
# 3. Roda: rtk npm test
# 4. Commit atômico
```

---

## 📝 O Que Colocar no AGENTS.md do Projeto

### ✅ Colocar (Agent NÃO consegue inferir)

```markdown
## Stack
- Next.js 16, PostgreSQL, Tailwind
- Use bun (não npm!)

## Commands
rtk bun test
rtk bun run dev

## Out of Scope
- Never touch: lib/billing/ (legacy, external team)
- Never modify: .env files

## Project-Specific
- API base: http://localhost:8080/api/v1
- Run migrations before schema changes: make migrate-up
- Brand voice: casual PT-BR (see .guidelines/brand.md)
```

### ❌ NÃO Colocar (Agent infere do código)

```markdown
❌ "Components go in src/components/" (óbvio do código)
❌ "Use TypeScript" (vê nos imports)
❌ "Functions should be small" (regra global já cobre)
❌ "Never commit secrets" (regra global já cobre)
❌ Documentação completa da API (link é suficiente)
```

**Regra:** Se agent consegue ver no código, não escreve. Se agent erra sempre, adiciona.

---

## 🔧 Comandos Úteis

### Atualizar Global (Afeta Todos os Projetos)

```bash
cd ~/work/agent
# Edita AGENTS.md, RTK.md, etc
rtk git add -A
rtk git commit -m "feat: add new global rule"

# Mudança reflete em todos os 3 harnesses automaticamente (symlinks!)
```

### Atualizar Projeto Específico

```bash
cd ~/work/meu-projeto
# Edita ./AGENTS.md (só este projeto)
rtk git add AGENTS.md
rtk git commit -m "docs: update project rules"
```

### Testar Se Tá Funcionando

```bash
cd ~/work/meu-projeto
claude

# Pergunta algo que só tá no AGENTS.md do projeto:
"what's the API base URL?"

# Se responder certo → tá lendo
# Se responder errado → não tá lendo (verifica symlinks)
```

---

## 🎓 Dicas de Uso

### 1. Mantenha AGENTS.md do Projeto CURTO

```bash
# Ruim (500 linhas)
- Explica tudo em detalhes
- Documenta cada arquivo
- Repete regras globais

# Bom (50-100 linhas)
- Stack + comandos
- Out of scope
- Gotchas específicos
```

### 2. Use RTK em Tudo

```bash
✅ rtk git status
✅ rtk npm test
✅ rtk make run

❌ git status (sem rtk)
```

60-90% menos tokens = contexto dura mais.

### 3. /clear Entre Tarefas

```bash
# Tarefa 1
"add login button"
# [agent faz]

/clear  # Limpa contexto

# Tarefa 2 (não relacionada)
"fix API rate limiting"
```

### 4. Cross-Reference Projetos

```bash
cd ~/work/projeto-novo
claude

"look at ../projeto-antigo/src/auth and implement same pattern here"
```

Agent reutiliza padrões bem.

---

## 🐛 Troubleshooting

### Agent Não Lê AGENTS.md do Projeto

```bash
# Verifica se existe
ls -la ./AGENTS.md

# Se não existe, cria:
cp ~/work/agent/AGENTS.template.md ./AGENTS.md
```

### Agent Ignora Regras Globais

```bash
# Verifica symlinks
ls -la ~/.claude/CLAUDE.md
ls -la ~/.config/opencode/AGENTS.md

# Se não existem, roda setup:
cd ~/work/agent
./setup-symlinks.sh
```

### AGENTS.md Muito Longo, Agent Ignora

```bash
# Encurta!
# Remove:
# - Coisas óbvias do código
# - Duplicações do global
# - Documentação extensa

# Meta: 50-100 linhas
```

### Skills Não Aparecem

```bash
# Verifica symlink
ls -la ~/.claude/skills/

# Deve apontar pra ~/work/agent/skills/
# Se não, roda setup novamente:
cd ~/work/agent
./setup-symlinks.sh
```

---

## 📊 Resumo

### Setup (Uma Vez)
1. `cd ~/work/agent && ./setup-symlinks.sh`
2. Verifica symlinks criados

### Por Projeto (Primeira Vez)
1. `cp ~/work/agent/AGENTS.template.md ./AGENTS.md`
2. Escolhe stack (Option 1, 2 ou 3)
3. Adiciona comandos + out of scope
4. Mantém curto (50-100 linhas)

### Uso Diário
1. `cd ~/work/projeto`
2. `claude`
3. Conversa natural: "add feature X"
4. Agent lê: global → RTK → projeto → código
5. Agent implementa seguindo todas as regras

**Simples assim!** 🎯

---

## 🎬 Exemplo Completo (Do Zero)

```bash
# 1. Setup global (uma vez)
cd ~/work/agent
./setup-symlinks.sh
# ✅ Symlinks criados

# 2. Novo projeto
cd ~/work/novo-app
cp ~/work/agent/AGENTS.template.md ./AGENTS.md

# 3. Edita AGENTS.md
# Escolhe Option 2 (Express + React)
# Adiciona comandos: npm run dev, npm test
# Adiciona out of scope: Never touch src/legacy/

# 4. Usa
claude

"add user registration with email validation"

# Agent:
# ✅ Lê global (just talk, RTK, etc)
# ✅ Lê projeto (Express + React, npm run dev)
# ✅ Implementa
# ✅ Escreve testes
# ✅ Roda: rtk npm test
# ✅ Commit atômico

/clear

"look at ../outro-projeto/auth and do the same"

# Agent:
# ✅ Lê outro projeto
# ✅ Replica padrão
# ✅ Adapta pro contexto atual
```

**Pronto pra usar!** 🚀
