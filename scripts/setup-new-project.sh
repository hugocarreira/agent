#!/bin/bash
# Setup AGENTS.md for a NEW project (no existing AGENTS.md)

set -e

WORKDIR="$HOME/work"
AGENT_ROOT="$HOME/work/agent"

if [ -z "$1" ]; then
    echo "Usage: ./setup-new-project.sh <project-name>"
    echo ""
    echo "Example: ./setup-new-project.sh my-app"
    echo "Creates: $WORKDIR/my-app/AGENTS.md from template"
    exit 1
fi

PROJECT_NAME="$1"
PROJECT_PATH="$WORKDIR/$PROJECT_NAME"

# Check if project exists
if [ ! -d "$PROJECT_PATH" ]; then
    echo "❌ Error: Project '$PROJECT_NAME' not found at $PROJECT_PATH"
    echo ""
    echo "Create it first:"
    echo "  mkdir -p $PROJECT_PATH"
    exit 1
fi

# Check if AGENTS.md already exists
if [ -f "$PROJECT_PATH/AGENTS.md" ]; then
    echo "⚠️  Warning: AGENTS.md already exists in $PROJECT_NAME"
    echo ""
    echo "Use setup-existing-project.sh instead, or remove existing file:"
    echo "  rm $PROJECT_PATH/AGENTS.md"
    exit 1
fi

# Copy template
echo "📝 Creating AGENTS.md for new project: $PROJECT_NAME"
cp "$AGENT_ROOT/AGENTS.template.md" "$PROJECT_PATH/AGENTS.md"

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
echo "Then start coding:"
echo "  claude"
