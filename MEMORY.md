# Memory & Observations

- The project relies heavily on the `internal/translator/` package to ensure requests and responses are strictly conformant between platforms.
- Some OpenAI-compatible providers (like Together AI, Groq, Mistral, xAI) implement the API slightly out-of-spec (specifically crashing or failing when `stream_options: {"include_usage": true}` is sent). These are explicitly ignored in `internal/runtime/executor/openai_compat_executor.go`.
- The Web UI relies on a `management.html` file. Originally, it pulled from a Github Release. Now, the repository includes the UI as a submodule `ui/` and uses `//go:embed` to serve the UI offline as a foolproof fallback.
- Gzip middleware for the Amp proxy modules occasionally failed with EOF due to chunked stream errors, fixed by skipping empty responses and nil checks.
