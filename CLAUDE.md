# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Status

This is a new project (`crud-productos`) with no application code yet. The repository has the OpenSpec spec-driven workflow installed but no tech stack, framework, or source files have been created.

## OpenSpec Workflow

This project uses OpenSpec for spec-driven development. All feature work should flow through the OpenSpec change lifecycle.

### CLI Commands

```bash
# List active changes
openspec list --json

# Create a new change
openspec new change "<kebab-case-name>"

# Check change status and artifact graph
openspec status --change "<name>" --json

# Get instructions for generating an artifact
openspec instructions <artifact-id> --change "<name>" --json
```

### Change Lifecycle

1. **Propose** (`/opsx:propose`) — Creates a change directory under `openspec/changes/<name>/` and generates `proposal.md`, `design.md`, and `tasks.md` in sequence.
2. **Explore** (`/opsx:explore`) — Thinking/investigation mode; reads codebase and artifacts freely but never writes application code.
3. **Apply** (`/opsx:apply`) — Implements tasks from `tasks.md` one at a time, marking each `- [ ]` → `- [x]` when done.
4. **Archive** (`/opsx:archive`) — Moves the completed change to `openspec/changes/archive/YYYY-MM-DD-<name>/`.

### Directory Layout

```
openspec/
  config.yaml              # Schema and per-artifact rules
  changes/
    <name>/
      .openspec.yaml       # Change metadata
      proposal.md          # What & why
      design.md            # How
      tasks.md             # Implementation checklist
      specs/               # Delta specs (if any)
    archive/               # Completed changes
  specs/                   # Canonical capability specs (synced from delta specs on archive)
```

### config.yaml

`openspec/config.yaml` is currently minimal (schema: spec-driven, no project context, no custom rules). When the tech stack is decided, add a `context:` block describing it so artifact generation is informed.
