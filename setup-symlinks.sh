#!/bin/bash
# Setup symlinks from all stacks to ~/work/agent

set -e

AGENT_ROOT="$HOME/work/agent"

echo "🔗 Setting up symlinks for all AI coding stacks..."
echo ""

# ============================================
# OpenCode Setup
# ============================================
echo "📦 OpenCode Configuration"

# Backup existing files
if [ -f "$HOME/.config/opencode/AGENTS.md" ]; then
    echo "  ↳ Backing up existing AGENTS.md"
    mv "$HOME/.config/opencode/AGENTS.md" "$HOME/.config/opencode/AGENTS.md.backup.$(date +%s)"
fi

# Create symlinks
echo "  ↳ Linking AGENTS.md"
ln -sf "$AGENT_ROOT/stacks/opencode/AGENTS.md" "$HOME/.config/opencode/AGENTS.md"

echo "  ↳ Linking agents/"
rm -rf "$HOME/.config/opencode/agents"
ln -sf "$AGENT_ROOT/agents" "$HOME/.config/opencode/agents"

echo "  ↳ Linking skills/"
rm -rf "$HOME/.config/opencode/skills"
ln -sf "$AGENT_ROOT/skills" "$HOME/.config/opencode/skills"

echo "  ↳ Linking plugins/"
rm -rf "$HOME/.config/opencode/plugins"
ln -sf "$AGENT_ROOT/plugins" "$HOME/.config/opencode/plugins"

# ============================================
# Codex Setup
# ============================================
echo ""
echo "📦 Codex Configuration"

# Backup existing
if [ -f "$HOME/.codex/AGENTS.md" ]; then
    echo "  ↳ Backing up existing AGENTS.md"
    mv "$HOME/.codex/AGENTS.md" "$HOME/.codex/AGENTS.md.backup.$(date +%s)"
fi

if [ -f "$HOME/.codex/RTK.md" ]; then
    echo "  ↳ Backing up existing RTK.md"
    mv "$HOME/.codex/RTK.md" "$HOME/.codex/RTK.md.backup.$(date +%s)"
fi

# Create symlinks
echo "  ↳ Linking AGENTS.md"
ln -sf "$AGENT_ROOT/stacks/codex/AGENTS.md" "$HOME/.codex/AGENTS.md"

echo "  ↳ Linking RTK.md"
ln -sf "$AGENT_ROOT/RTK.md" "$HOME/.codex/RTK.md"

echo "  ↳ Linking skills/"
rm -rf "$HOME/.codex/skills"
ln -sf "$AGENT_ROOT/skills" "$HOME/.codex/skills"

# ============================================
# Claude Code Setup
# ============================================
echo ""
echo "📦 Claude Code Configuration"

# Backup existing
if [ -f "$HOME/.claude/CLAUDE.md" ]; then
    echo "  ↳ Backing up existing CLAUDE.md"
    mv "$HOME/.claude/CLAUDE.md" "$HOME/.claude/CLAUDE.md.backup.$(date +%s)"
fi

if [ -f "$HOME/.claude/AGENTS.md" ]; then
    echo "  ↳ Backing up existing AGENTS.md"
    mv "$HOME/.claude/AGENTS.md" "$HOME/.claude/AGENTS.md.backup.$(date +%s)"
fi

if [ -f "$HOME/.claude/RTK.md" ]; then
    echo "  ↳ Backing up existing RTK.md"
    mv "$HOME/.claude/RTK.md" "$HOME/.claude/RTK.md.backup.$(date +%s)"
fi

# Create symlinks
echo "  ↳ Linking CLAUDE.md"
ln -sf "$AGENT_ROOT/stacks/claude/CLAUDE.md" "$HOME/.claude/CLAUDE.md"

echo "  ↳ Linking AGENTS.md"
ln -sf "$AGENT_ROOT/stacks/claude/CLAUDE.md" "$HOME/.claude/AGENTS.md"

echo "  ↳ Linking RTK.md"
ln -sf "$AGENT_ROOT/RTK.md" "$HOME/.claude/RTK.md"

echo "  ↳ Linking skills/"
rm -rf "$HOME/.claude/skills"
ln -sf "$AGENT_ROOT/skills" "$HOME/.claude/skills"

# ============================================
# Summary
# ============================================
echo ""
echo "✅ Symlink setup complete!"
echo ""
echo "📍 All stacks now point to:"
echo "   • Agent configs: $AGENT_ROOT/stacks/{opencode,codex,claude}/"
echo "   • Skills: $AGENT_ROOT/skills/"
echo "   • Agents: $AGENT_ROOT/agents/"
echo "   • RTK config: $AGENT_ROOT/RTK.md"
echo ""
echo "🔍 Verify with:"
echo "   ls -la ~/.config/opencode/AGENTS.md"
echo "   ls -la ~/.codex/AGENTS.md"
echo "   ls -la ~/.claude/CLAUDE.md"
echo ""
