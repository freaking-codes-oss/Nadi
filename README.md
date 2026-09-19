# Nadi

> Nadi: an auditable AI operator for your Android terminal.

Nadi is a Termux-native Go agent designed around explicit approval, resumability, and auditability. High-impact behavior is never enabled autonomously.

## Current status

The `feature/nadi-foundation` branch contains the first runtime slice:

- Cobra CLI with `doctor`, `inspect`, and one-shot task intake
- SQLite task and audit persistence
- Approval manager and deny-by-default policy engine
- Extensible tool registry and controlled shell tool
- OpenAI-compatible provider client with retry handling
- Secret redaction, resumable session state, and Termux installer

Tool execution is intentionally not wired into autonomous model loops yet. This keeps the safety boundary explicit while the agent runtime is developed.

## Build and test

```sh
go build -o nadi ./cmd/nadi
go test ./...
```

## Termux install

```sh
pkg install golang
./install.sh
nadi doctor
```

Never put API keys in command arguments or shell history. Configure providers through environment variables or the forthcoming TOML configuration layer.
