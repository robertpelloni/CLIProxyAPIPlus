# TODO.md

## Short-term Tasks & Bug Fixes
- [ ] Add tooltips and extensive label descriptions across the UI React components to help new users understand API limits and routing rules.
- [ ] Investigate partially implemented models and ensure they are connected to both backend translators and frontend visualizers.
- [ ] Increase the robustness of `internal/translator/` by adding fallback mechanisms for when upstream APIs respond with non-standard error codes.
- [ ] Refactor the UI auto-updater mechanism. Since the UI is embedded now, evaluate if we still need the runtime auto-updater fetching from Github. If so, provide a clear UI toggle.
- [ ] Conduct a thorough review of the code to find missing/hidden functionality to represent it properly in the Web UI.

## Long-term Plans (Moved to ROADMAP.md if they become structural)
- Systematically research and implement all emerging LLM inference providers.
