# CLIProxyAPI Plus - Architecture, Patterns, and Decisions

## High-Level Overview
CLIProxyAPI Plus acts as a universal proxy intercepting calls to leading inference platforms (OpenAI, Anthropic, Gemini, Codex) and transparently normalizing, routing, and translating these payloads via a robust canonical `ThinkingConfig` engine. 

## Architectural Modules
1. **Frontend UI (`ui/`)**: A Single-Page React Application (Cli-Proxy-API-Management-Center) heavily leveraging `Zustand` for state management and `react-router-dom`. It is kept as a git submodule to cleanly separate the complex JavaScript ecosystem from the main Go codebase.
2. **Go Backend (`cmd/server` & `internal/`)**: Uses `gin-gonic/gin` for highly performant HTTP multiplexing.
3. **Translators & Executors**: The codebase employs a strict distinction between *Translating* (converting specific JSON dialects like Claude to Canonical) and *Executing* (sending the finalized canonical JSON payload to upstream APIs like Gemini and managing the connection lifecycle).

## Patterns and Implementations
### UI Embedding (`go:embed`)
Instead of serving HTML files from disk paths which can break in Docker or remote deployment scenarios, the project builds the React UI into a single HTML file via Vite, and uses `go:embed` inside `internal/managementasset/ui.go`. 
*   **Result**: The server guarantees that the management portal is always shipped natively alongside the proxy logic.

### Routing Elegance
`internal/api/server.go` leverages the Gin framework to bind the embedded UI logic to three critical endpoints simultaneously:
1.  `/` - To greet any user dropping directly onto the port with the management panel.
2.  `/management.html` - Legacy support.
3.  `s.engine.NoRoute` - To gracefully trap 404s and redirect them to the frontend router where React router can appropriately handle them.

### Auto-Updater Decoupling
Because `go:embed` ensures the UI is always present, the legacy dynamic GitHub `StartAutoUpdater` polling loop in `internal/managementasset/updater.go` was cleverly bypassed if `len(GetIndexHTML()) > 0`. This significantly reduces startup network noise and background polling for self-hosted instances.

## Design Decisions
*   **User Empathy (UI Labels)**: The user wants to ensure the UI is extremely clear. Descriptions are wrapped using `i18n` translation hooks like `t('quota_management.description', 'Manage and monitor...')` to provide immediate English fallback context if translation strings aren't mapped yet.
*   **Strict Testing**: The backend relies heavily on `go test ./...`. Breaking any translator interface will immediately fail the suite. The UI utilizes strict ESLint and TypeScript checking to prevent type regressions before the static HTML is generated.