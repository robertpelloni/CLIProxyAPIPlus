# CHANGELOG

## [Unreleased]
- Added Submodule `Cli-Proxy-API-Management-Center` as `ui/`.
- Modified `openai_compat_executor` to explicitly gate `stream_options.include_usage` to avoid crashing specific newcomers like Together AI, Mistral, xAI, Groq, and DeepSeek on chat completion streams.
- Updated `config.example.yaml` to showcase prominent OpenAI-compatible endpoints.
- Embedded `ui_build/management.html` using `//go:embed` inside `internal/managementasset/embed.go` as a static offline fallback when GitHub Releases fetch fails.
- Created `Makefile` for unified `go generate` compilation of the front-end assets alongside the proxy server itself.
