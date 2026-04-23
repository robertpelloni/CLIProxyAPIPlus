# ROADMAP.md

This document outlines the long-term structural and architectural vision for CLIProxyAPI Plus.

## Q3/Q4 Goals
1. **Universal LLM Parity**: Provide 100% translation feature parity (Function Calling, System Prompts, Image Modalities) between all tier-1 providers: OpenAI, Claude, Google Gemini, and Codex.
2. **Provider Plugins Ecosystem**: Transition the `internal/translator` logic into a dynamically loadable plugin ecosystem where contributors can maintain their own 3rd party vendor APIs without muddying the core codebase.
3. **Advanced Rate Limiting & Cost Guardrails**: Implement token-budgeting and dollar-cost limiting directly into the proxy. The `ui/` Dashboard will feature an interactive pricing/cost chart based on real-time token tracking.
4. **Enhanced Telemetry**: Build out robust request logging interfaces via OpenTelemetry for massive enterprise setups.
