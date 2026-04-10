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
