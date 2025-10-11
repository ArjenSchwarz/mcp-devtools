# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [Unreleased]

### Fixed
- GitHub Copilot Agent Tool output filtering:
  - Revised filtering strategy to extract content after last progress indicator (●, ✓, ✗, ↪)
  - Fixed issue where actual answers on same line as progress indicator were filtered out
  - Updated to handle Unicode progress characters correctly using rune slicing
  - Improved content extraction to capture multi-line responses after last indicator
  - Maintained filtering of command traces ($ prefix) and usage statistics section

### Changed
- GitHub Copilot Agent Tool specification updates:
  - Updated design.md with revised output filtering implementation based on actual Copilot CLI behaviour
  - Added decision log entry documenting filtering strategy revision based on integration testing discoveries
  - Updated tasks.md with bug fix subtasks for output filtering implementation and testing
- GitHub Copilot Agent Tool test coverage:
  - Revised filter tests to match new output extraction behaviour
  - Added test cases for answers on same line as progress indicator
  - Added test cases for multi-line answers after last indicator
  - Updated edge case tests to reflect revised filtering logic

### Added
- Documentation for GitHub Copilot Agent Tool:
  - Created [docs/tools/copilot-agent.md](docs/tools/copilot-agent.md) with usage examples, configuration details, and troubleshooting guide
  - Added copilot-agent to README.md agents table with emoji icon and description
  - Added copilot-agent to [docs/tools/overview.md](docs/tools/overview.md) with configuration examples
  - Updated internal/tools/enablement.go comments to include copilot-agent in supported tools list
- golangci-lint configuration file (.golangci.yml) with version 2 configuration and enabled linters

### Changed
- Exported FilterOutput method in copilot agent tool to enable unit testing
- Updated FilterOutput implementation to use strings.HasPrefix for symbol filtering instead of rune comparison
- Completed GitHub Copilot Agent Tool test coverage with comprehensive unit tests
- Updated specification tasks to mark core implementation and testing as completed

### Fixed
- Fixed golangci-lint configuration compatibility issues

- GitHub Copilot Agent Tool implementation:
  - New `copilot-agent` tool providing access to GitHub Copilot CLI through MCP
  - Support for core Copilot CLI features: model selection, session management (resume and session ID), and permission controls
  - Configurable timeout (AGENT_TIMEOUT) and response size limits (AGENT_MAX_RESPONSE_SIZE)
  - Output filtering to remove Copilot metadata and progress indicators
  - Tool-specific permission management via allow-tool and deny-tool parameters
  - Extended help documentation with examples and troubleshooting guidance
- GitHub Copilot Agent Tool specification documents:
  - Requirements document outlining integration with GitHub Copilot CLI as an MCP tool
  - Design document detailing architecture, component structure, and implementation approach
  - Decision log tracking key design decisions including parameter naming, security model, and output filtering
  - Peer review validation document synthesising external AI feedback from Gemini, Codex, and Q Developer
  - Tasks document providing implementation checklist
