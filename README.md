# Ai-agent-test

This repository is a small sandbox for testing Codex and AI-agent workflows.

## Codex Project Files

- `AGENTS.md` contains repository-specific instructions Codex should follow.
- `.codex/config.toml` contains conservative project defaults for local Codex runs.

## Try It

From this repository root, ask Codex to summarize the active project guidance:

```bash
codex --ask-for-approval never "Summarize the current project instructions."
```

To test Codex as an MCP server for agent orchestration:

```bash
codex mcp-server
```

Keep secrets such as `OPENAI_API_KEY` in a local `.env` file that is not committed.
