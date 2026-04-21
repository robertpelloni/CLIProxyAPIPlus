<<<<<<< HEAD
# CHANGELOG

## [Unreleased]
- Added Submodule `Cli-Proxy-API-Management-Center` as `ui/`.
- Modified `openai_compat_executor` to explicitly gate `stream_options.include_usage` to avoid crashing specific newcomers like Together AI, Mistral, xAI, Groq, and DeepSeek on chat completion streams.
- Updated `config.example.yaml` to showcase prominent OpenAI-compatible endpoints.
- Embedded `ui_build/management.html` using `//go:embed` inside `internal/managementasset/embed.go` as a static offline fallback when GitHub Releases fetch fails.
- Created `Makefile` for unified `go generate` compilation of the front-end assets alongside the proxy server itself.
=======
# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added
- Added `ui` submodule ([Cli-Proxy-API-Management-Center](https://github.com/robertpelloni/Cli-Proxy-API-Management-Center)) to embed the web UI natively in the proxy binary.
- Added `go:embed` to serve the Management Web UI directly from memory, eliminating reliance on local disk paths or github downloading.
- Updated server routes so that `/`, `/management.html`, and `NoRoute` all seamlessly serve the Management UI.
- Extensive new documentation files added: `VISION.md`, `MEMORY.md`, and `DEPLOY.md` to capture architectural goals, ongoing AI context, and deployment instructions respectively.

### Changed
- Refactored `server.go` router logic to elegantly fallback to serving the UI via embedded bytes instead of strictly file system paths.
>>>>>>> origin/jules-9238706904812453426-8fd51539
