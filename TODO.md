<<<<<<< HEAD
# TODO

- Enhance error handling and fallbacks inside `internal/translator/` when OpenAI incompatible schemas are detected (e.g. specialized response schema mapping).
- Refactor the front-end code inside `ui/src/` to support localized currency estimates rather than fixed token calculations (some API newcomers do dynamic pricing).
- Implement a single text file (e.g. `VERSION.md`) that drives all build artifacts globally without hard-coding variables inside main.go.
- Connect more complex multi-modal API translation layers (Audio to text).
- Re-verify cross-platform building via `go generate make generate-ui` without node.js installed on the system building it (requires CI/CD artifact separation).
=======
# TODO.md

## Short-term Tasks & Bug Fixes
- [x] Add tooltips and extensive label descriptions across the UI React components to help new users understand API limits and routing rules.
- [x] Investigate partially implemented models and ensure they are connected to both backend translators and frontend visualizers.
- [x] Increase the robustness of `internal/translator/` by adding fallback mechanisms for when upstream APIs respond with non-standard error codes.
- [x] Refactor the UI auto-updater mechanism. Since the UI is embedded now, evaluate if we still need the runtime auto-updater fetching from Github. If so, provide a clear UI toggle.
- [x] Conduct a thorough review of the code to find missing/hidden functionality to represent it properly in the Web UI.

## Long-term Plans (Moved to ROADMAP.md if they become structural)
- Systematically research and implement all emerging LLM inference providers.
>>>>>>> origin/jules-9238706904812453426-8fd51539
