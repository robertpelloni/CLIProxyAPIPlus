# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [6.1.0] - 2026-04-22
### Added
- **Provider Plugins Ecosystem**: Dynamically load Go `.so` plugins at runtime via `LoadPlugins()` without muddying the core codebase.
- Added `plugin-dir` config setting and UI visualizer input to point to `.so` modules.
- Implemented robust multi-modal Audio-to-Text API translation layer for OpenAI <-> Claude.
- Introduced `latency` and `fill-first` routing options across backend configuration and the UI Management Center.
- Added DeepSeek Coder V2 registration.
- Added `ui` submodule ([Cli-Proxy-API-Management-Center](https://github.com/robertpelloni/Cli-Proxy-API-Management-Center)) to embed the web UI natively in the proxy binary.
- Added `go:embed` to serve the Management Web UI directly from memory, eliminating reliance on local disk paths or github downloading.
- Updated server routes so that `/`, `/management.html`, and `NoRoute` all seamlessly serve the Management UI.
- Extensive new documentation files added: `VISION.md`, `MEMORY.md`, and `DEPLOY.md` to capture architectural goals, ongoing AI context, and deployment instructions respectively.

### Changed
- Refactored `server.go` router logic to elegantly fallback to serving the UI via embedded bytes instead of strictly file system paths.
