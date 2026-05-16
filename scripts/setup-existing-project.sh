#!/bin/bash
# Setup AGENTS.md for an EXISTING project (has AGENTS.md already)

set -e

WORKDIR="$HOME/work"
AGENT_ROOT="$HOME/work/agent"

if [ -z "$1" ]; then
    echo "Usage: ./setup-existing-project.sh <project-name>"
    echo ""
    echo "Example: ./setup-existing-project.sh my-app"
    echo "Adds global reference to existing $WORKDIR/my-app/AGENTS.md"
    exit 1
fi

PROJECT_NAME="$1"
PROJECT_PATH="$WORKDIR/$PROJECT_NAME"
AGENTS_FILE="$PROJECT_PATH/AGENTS.md"

# Check if project exists
if [ ! -d "$PROJECT_PATH" ]; then
    echo "❌ Error: Project '$PROJECT_NAME' not found at $PROJECT_PATH"
    exit 1
fi

# Check if AGENTS.md exists
if [ ! -f "$AGENTS_FILE" ]; then
    echo "❌ Error: AGENTS.md not found in $PROJECT_NAME"
    echo ""
    echo "This project has no AGENTS.md. Use setup-new-project.sh instead:"
    echo "  ./setup-new-project.sh $PROJECT_NAME"
    exit 1
fi

# Check if already has reference
if grep -q "READ ~/work/agent/AGENTS.md" "$AGENTS_FILE"; then
    echo "✅ AGENTS.md already references global config"
    echo ""
    echo "File: $AGENTS_FILE"
    echo "All set! Nothing to do."
    exit 0
fi

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
echo "Then start your AI session (opencode/codex/claude)"
