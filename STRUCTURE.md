# Project Structure & Submodules

## Codebase Overview
CLIProxyAPI Plus operates as a Go backend acting as a local/remote LLM routing proxy, featuring a unified embedded React UI frontend.

### Submodules
1. **`ui` (React Dashboard)**
   - **Repository:** `https://github.com/robertpelloni/Cli-Proxy-API-Management-Center`
   - **Path:** `./ui`
   - **Version:** Latest `main` (Embedded at build-time).
   - **Description:** A single-page application built with React, Vite, and Zustand. It serves as the visual configuration editor for the proxy's `config.yaml`, quota management, and payload filtering. It is compiled to `dist/index.html` and injected via `//go:embed` inside `internal/managementasset/ui.go`.

### Key Directories
- **`cmd/`**: Entrypoints for binaries. `cmd/server/main.go` is the core proxy application.
- **`internal/`**: Core non-exported logic.
  - **`api/`**: The Gin-based HTTP router mapping `/v1/chat/completions` etc.
  - **`auth/`**: Provider-specific login flows (e.g. `claude`, `codex`, `gemini`, `kilo`).
  - **`config/`**: YAML parsing logic and internal state representations (`config.go`).
  - **`translator/`**: Natively maps JSON schemas between divergent APIs (e.g. Anthropic to OpenAI formats).
- **`sdk/`**: Exported utilities, including the dynamic `translator/plugin.go` system for external loading.
- **`examples/`**: Demonstration templates, such as `dummy-plugin/` for illustrating how to build external `.so` translation plugins.
