# TODO

- Enhance error handling and fallbacks inside `internal/translator/` when OpenAI incompatible schemas are detected (e.g. specialized response schema mapping).
- Refactor the front-end code inside `ui/src/` to support localized currency estimates rather than fixed token calculations (some API newcomers do dynamic pricing).
- Implement a single text file (e.g. `VERSION.md`) that drives all build artifacts globally without hard-coding variables inside main.go.
- Connect more complex multi-modal API translation layers (Audio to text).
- Re-verify cross-platform building via `go generate make generate-ui` without node.js installed on the system building it (requires CI/CD artifact separation).
