#!/bin/bash
# Interactive setup for project AGENTS.md

set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
AGENT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
WORKDIR="$HOME/work"

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "  Project Setup - AGENTS.md"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""

# Ask for project name
read -p "Project name (in ~/work/): " PROJECT_NAME

if [ -z "$PROJECT_NAME" ]; then
    echo "❌ Error: Project name cannot be empty"
    exit 1
fi

PROJECT_PATH="$WORKDIR/$PROJECT_NAME"

# Check if project directory exists
if [ ! -d "$PROJECT_PATH" ]; then
    echo ""
    echo "❌ Error: Project '$PROJECT_NAME' not found at $PROJECT_PATH"
    echo ""
    echo "Create it first:"
    echo "  mkdir -p $PROJECT_PATH"
    exit 1
fi

echo ""
echo "Project: $PROJECT_PATH"
echo ""

# Check if AGENTS.md exists
AGENTS_FILE="$PROJECT_PATH/AGENTS.md"
HAS_AGENTS=false

if [ -f "$AGENTS_FILE" ]; then
    HAS_AGENTS=true
    echo "ℹ️  Found existing AGENTS.md"
else
    echo "ℹ️  No AGENTS.md found"
fi

echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "  Select option:"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "  1) New project (create AGENTS.md from template)"
echo "  2) Existing project (add global reference to existing AGENTS.md)"
echo "  3) Cancel"
echo ""
read -p "Option [1-3]: " OPTION

case $OPTION in
    1)
        # New project
        if [ "$HAS_AGENTS" = true ]; then
            echo ""
            echo "⚠️  Warning: AGENTS.md already exists!"
            read -p "Overwrite? [y/N]: " CONFIRM
            if [[ ! "$CONFIRM" =~ ^[Yy]$ ]]; then
                echo "Cancelled."
                exit 0
            fi
            # Backup existing
            BACKUP_FILE="$AGENTS_FILE.backup.$(date +%s)"
            cp "$AGENTS_FILE" "$BACKUP_FILE"
            echo "📦 Backup created: $BACKUP_FILE"
        fi

        echo ""
        echo "📝 Creating AGENTS.md from template..."
        cp "$AGENT_ROOT/AGENTS.template.md" "$AGENTS_FILE"
        
        echo "✅ Done!"
        echo ""
        echo "Next steps:"
        echo "  cd $PROJECT_PATH"
        echo "  # Edit AGENTS.md:"
        echo "  #   1. Pick stack option (1, 2, or 3) - delete others"
        echo "  #   2. Add project-specific commands"
        echo "  #   3. Add 'Out of Scope' rules"
        echo "  #   4. Keep it SHORT (50-100 lines)"
        echo ""
        echo "Then start your AI"
        ;;
        
    2)
        # Existing project
        if [ "$HAS_AGENTS" = false ]; then
            echo ""
            echo "❌ Error: No AGENTS.md found"
            echo ""
            echo "Use option 1 (New project) instead."
            exit 1
        fi

        # Check if already has reference
        if grep -q "READ ~/work/agent/AGENTS.md" "$AGENTS_FILE"; then
            echo ""
            echo "✅ AGENTS.md already references global config"
            echo ""
            echo "File: $AGENTS_FILE"
            echo "All set! Nothing to do."
            exit 0
        fi

        echo ""
        echo "📝 Adding global reference to existing AGENTS.md..."
        
        # Backup
        BACKUP_FILE="$AGENTS_FILE.backup.$(date +%s)"
        cp "$AGENTS_FILE" "$BACKUP_FILE"
        echo "📦 Backup created: $BACKUP_FILE"

        # Add reference at the top
        TEMP_FILE=$(mktemp)
        echo "READ ~/work/agent/AGENTS.md BEFORE ANYTHING." > "$TEMP_FILE"
        echo "" >> "$TEMP_FILE"
        cat "$AGENTS_FILE" >> "$TEMP_FILE"
        mv "$TEMP_FILE" "$AGENTS_FILE"

        echo "✅ Done!"
        echo ""
        echo "Added global reference to: $AGENTS_FILE"
        echo ""
        echo "Next steps:"
        echo "  cd $PROJECT_PATH"
        echo "  # Review AGENTS.md and remove duplications"
        echo "  # (Global config already covers common rules)"
        echo "  # Keep only project-specific stuff"
        echo ""
        echo "Then start your AI"
        ;;
        
    3)
        echo ""
        echo "Cancelled."
        exit 0
        ;;
        
    *)
        echo ""
        echo "❌ Invalid option: $OPTION"
        exit 1
        ;;
esac
