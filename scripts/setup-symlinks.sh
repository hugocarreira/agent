#!/bin/bash
# Setup symlinks from all stacks to ~/work/agent

set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
AGENT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

echo "🔗 Setting up AI coding stacks..."
echo ""
echo "Philosophy: Just talk to it. No orchestrators, no charades."
echo ""

# ============================================
# RTK (Rust Token Killer)
# ============================================
if command -v rtk &>/dev/null; then
    echo "✅ RTK already installed ($(rtk --version))"
else
    echo "📦 Installing RTK (Rust Token Killer)..."
    if command -v brew &>/dev/null; then
        brew install rtk-ai/tap/rtk
        echo "  ↳ RTK installed ($(rtk --version))"
    else
        echo "  ⚠️  Homebrew not found — install RTK manually:"
        echo "     brew install rtk-ai/tap/rtk"
        echo "     or: curl -fsSL https://raw.githubusercontent.com/rtk-ai/rtk/refs/heads/master/install.sh | sh"
    fi
fi
echo ""

# ============================================
# OpenCode Setup
# ============================================
echo "📦 OpenCode Configuration"

# Backup existing files
if [ -f "$HOME/.config/opencode/AGENTS.md" ] && [ ! -L "$HOME/.config/opencode/AGENTS.md" ]; then
    echo "  ↳ Backing up existing AGENTS.md"
    mv "$HOME/.config/opencode/AGENTS.md" "$HOME/.config/opencode/AGENTS.md.backup.$(date +%s)"
fi

# Generate AGENTS.md (concatenate core files so opencode reads full content)
echo "  ↳ Generating AGENTS.md"
rm -f "$HOME/.config/opencode/AGENTS.md"
{
    if command -v rtk &>/dev/null; then
        cat "$AGENT_ROOT/RTK.md"
        echo ""
        echo "---"
        echo ""
    fi
    cat "$AGENT_ROOT/AGENTS.md"
} > "$HOME/.config/opencode/AGENTS.md"

echo "  ↳ Linking skills/"
rm -rf "$HOME/.config/opencode/skills"
ln -sf "$AGENT_ROOT/skills" "$HOME/.config/opencode/skills"

echo "  ↳ Linking plugins/"
rm -rf "$HOME/.config/opencode/plugins"
ln -sf "$AGENT_ROOT/plugins" "$HOME/.config/opencode/plugins"

# Remove old agents symlink if exists (we killed orchestrator)
if [ -L "$HOME/.config/opencode/agents" ]; then
    echo "  ↳ Removing old orchestrator agents/"
    rm -f "$HOME/.config/opencode/agents"
fi

# ============================================
# Codex Setup
# ============================================
echo ""
if [ -d "$HOME/.codex" ]; then
    echo "📦 Codex Configuration"

    # Backup existing
    if [ -f "$HOME/.codex/AGENTS.md" ] && [ ! -L "$HOME/.codex/AGENTS.md" ]; then
        echo "  ↳ Backing up existing AGENTS.md"
        mv "$HOME/.codex/AGENTS.md" "$HOME/.codex/AGENTS.md.backup.$(date +%s)"
    fi

    if [ -f "$HOME/.codex/RTK.md" ] && [ ! -L "$HOME/.codex/RTK.md" ]; then
        echo "  ↳ Backing up existing RTK.md"
        mv "$HOME/.codex/RTK.md" "$HOME/.codex/RTK.md.backup.$(date +%s)"
    fi

    # Create symlinks
    echo "  ↳ Linking AGENTS.md"
    rm -f "$HOME/.codex/AGENTS.md"
    ln -sf "$AGENT_ROOT/stacks/codex/AGENTS.md" "$HOME/.codex/AGENTS.md"

    echo "  ↳ Linking RTK.md"
    rm -f "$HOME/.codex/RTK.md"
    ln -sf "$AGENT_ROOT/RTK.md" "$HOME/.codex/RTK.md"

    echo "  ↳ Linking skills/"
    rm -rf "$HOME/.codex/skills"
    ln -sf "$AGENT_ROOT/skills" "$HOME/.codex/skills"
else
    echo "⏭️  Codex — skipped (no ~/.codex found)"
fi

# ============================================
# Claude Code Setup
# ============================================
echo ""
if [ -d "$HOME/.claude" ]; then
    echo "📦 Claude Code Configuration"

    # Backup existing
    if [ -f "$HOME/.claude/CLAUDE.md" ] && [ ! -L "$HOME/.claude/CLAUDE.md" ]; then
        echo "  ↳ Backing up existing CLAUDE.md"
        mv "$HOME/.claude/CLAUDE.md" "$HOME/.claude/CLAUDE.md.backup.$(date +%s)"
    fi

    if [ -f "$HOME/.claude/AGENTS.md" ] && [ ! -L "$HOME/.claude/AGENTS.md" ]; then
        echo "  ↳ Backing up existing AGENTS.md"
        mv "$HOME/.claude/AGENTS.md" "$HOME/.claude/AGENTS.md.backup.$(date +%s)"
    fi

    if [ -f "$HOME/.claude/RTK.md" ] && [ ! -L "$HOME/.claude/RTK.md" ]; then
        echo "  ↳ Backing up existing RTK.md"
        mv "$HOME/.claude/RTK.md" "$HOME/.claude/RTK.md.backup.$(date +%s)"
    fi

    # Create symlinks
    echo "  ↳ Linking CLAUDE.md"
    rm -f "$HOME/.claude/CLAUDE.md"
    ln -sf "$AGENT_ROOT/stacks/claude/CLAUDE.md" "$HOME/.claude/CLAUDE.md"

    echo "  ↳ Linking AGENTS.md (points to CLAUDE.md)"
    rm -f "$HOME/.claude/AGENTS.md"
    ln -sf "$AGENT_ROOT/stacks/claude/CLAUDE.md" "$HOME/.claude/AGENTS.md"

    echo "  ↳ Linking RTK.md"
    rm -f "$HOME/.claude/RTK.md"
    ln -sf "$AGENT_ROOT/RTK.md" "$HOME/.claude/RTK.md"

    echo "  ↳ Linking skills/"
    rm -rf "$HOME/.claude/skills"
    ln -sf "$AGENT_ROOT/skills" "$HOME/.claude/skills"
else
    echo "⏭️  Claude Code — skipped (no ~/.claude found)"
fi

# ============================================
# Summary
# ============================================
echo ""
echo "✅ Symlink setup complete!"
echo ""
echo "📍 All stacks now point to:"
echo "   • Core rules: $AGENT_ROOT/AGENTS.md"
echo "   • Stack configs: $AGENT_ROOT/stacks/{opencode,codex,claude}/"
echo "   • Skills: $AGENT_ROOT/skills/"
echo "   • RTK config: $AGENT_ROOT/RTK.md"
echo ""
echo "🔍 Verify with:"
echo "   ls -la ~/.config/opencode/AGENTS.md"
echo "   ls -la ~/.codex/AGENTS.md"
echo "   ls -la ~/.claude/CLAUDE.md"
echo ""
echo "🎯 Philosophy: Just talk to it. No orchestrators, no charades."
echo ""
