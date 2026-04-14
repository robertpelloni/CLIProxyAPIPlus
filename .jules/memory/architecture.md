# CLIProxyAPI Plus - Architecture, Patterns, and Decisions

## High-Level Overview
CLIProxyAPI Plus is a robust, Go-based proxy server designed to provide a unified, zero-configuration interface for leading AI models (OpenAI, Anthropic Claude, Google Gemini, GitHub Copilot/Codex). It bridges the gap between disparate provider APIs, offering seamless translation, load balancing, quota management, and token routing.

## Core Architecture
The project architecture is heavily modularized, utilizing the **Gin HTTP framework** for routing and middleware. 

1.  **Entrypoint (`cmd/server/main.go`)**: Parses flags, loads configurations, and launches the Gin server or specific CLI/auth modes.
2.  **API Layer (`internal/api/`)**: Defines Gin routes, middleware (auth, cors, usage tracking), and the core HTTP Server struct.
3.  **Thinking Pipeline (`internal/thinking/`)**: This is the heart of the normalization engine. It takes disparate provider payloads, normalizes them into a canonical `ThinkingConfig` representation, validates them, and translates them back out into provider-specific formats via `ProviderApplier`.
4.  **Runtime Executors (`internal/runtime/executor/`)**: Handlers that actually establish the connections and execute the requests to upstream providers (e.g., handling WebSockets for Codex).
5.  **Translators (`internal/translator/`)**: Dedicated packages for translating protocol A to protocol B (e.g., OpenAI to Claude, Codex to Gemini).
6.  **Web UI Submodule (`ui/`)**: A React/Vite single-page application ([Cli-Proxy-API-Management-Center](https://github.com/robertpelloni/Cli-Proxy-API-Management-Center)) that acts as the Management Center. It interacts with the Go backend via `/v0/management` endpoints.

## Recent Architectural Decisions & Patterns

### 1. Native UI Embedding (`go:embed`)
Instead of relying on the host file system or a fragile background auto-updater pulling from GitHub to serve the Management Web UI, the project now natively embeds the `ui/dist/index.html` file directly into the compiled Go binary using `go:embed`.
*   **Pattern**: The `serveManagementControlPanel` handler in `server.go` first checks the embedded byte slice. If present, it serves it directly from memory with a `text/html` content type.
*   **Fallback**: If the embedded asset isn't found (e.g., built incorrectly), it falls back to the legacy file system and auto-updater behavior.

### 2. Seamless UI Routing
The server router has been updated so that the management UI is instantly accessible at multiple intuitive entry points:
*   `/management.html` (Legacy explicit path)
*   `/` (Root path, for zero-friction access when opening the port directly in a browser)
*   `NoRoute` (Fallback, capturing any unhandled `GET` requests and surfacing the UI, which handles its own frontend routing)

### 3. Code Conventions and State Management
*   **Strict Documentation**: The project mandates extensive internal and user-facing documentation (`VISION.md`, `MEMORY.md`, `DEPLOY.md`, `TODO.md`, `CHANGELOG.md`).
*   **Logging**: `logrus` is used exclusively for structured logging. Panics and fatal exits are strictly discouraged in favor of returning errors up the stack.
*   **Submodules**: The `ui` folder is actively maintained as a git submodule to cleanly decouple the complex React frontend lifecycle from the Go backend, while still compiling together seamlessly.

## Future Roadmap Focus (TODO.md / VISION.md)
*   **UI/UX Improvements**: Extensive tooltips, descriptive labels, and clearer layout paradigms within the embedded React app to guide novice users through complex proxy routing settings.
*   **Robustness**: Fortifying the translator packages to intelligently handle non-standard upstream HTTP errors without panicking.
*   **Provider Parity**: Continuous research and seamless implementation of new inference providers into the canonical thinking pipeline.