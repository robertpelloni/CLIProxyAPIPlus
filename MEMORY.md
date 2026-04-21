<<<<<<< HEAD
# Memory & Observations

- The project relies heavily on the `internal/translator/` package to ensure requests and responses are strictly conformant between platforms.
- Some OpenAI-compatible providers (like Together AI, Groq, Mistral, xAI) implement the API slightly out-of-spec (specifically crashing or failing when `stream_options: {"include_usage": true}` is sent). These are explicitly ignored in `internal/runtime/executor/openai_compat_executor.go`.
- The Web UI relies on a `management.html` file. Originally, it pulled from a Github Release. Now, the repository includes the UI as a submodule `ui/` and uses `//go:embed` to serve the UI offline as a foolproof fallback.
- Gzip middleware for the Amp proxy modules occasionally failed with EOF due to chunked stream errors, fixed by skipping empty responses and nil checks.
=======
# MEMORY.md

## Ongoing Observations
- The codebase relies heavily on the Gin HTTP framework.
- The `internal/thinking` module is paramount. It normalizes inputs into a canonical format before applying provider-specific payloads.
- The project has now integrated a `ui` submodule pointing to [Cli-Proxy-API-Management-Center](https://github.com/robertpelloni/Cli-Proxy-API-Management-Center).
- `go:embed` is now utilized to bundle the UI natively, meaning we no longer need to depend on `MANAGEMENT_STATIC_PATH` or the github auto-updater for normal UI availability.

## Design Preferences
- Keep functions small and modular.
- Use Logrus for structured logging without leaking secrets.
- Rely on English documentation except for explicitly localized files.
- Version numbers should ideally be stored in a centralized location to prevent divergence. We currently rely on build arguments `VERSION` passed during compilation. We track logical version updates in `CHANGELOG.md`.
>>>>>>> origin/jules-9238706904812453426-8fd51539
