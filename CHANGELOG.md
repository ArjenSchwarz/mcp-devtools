# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [Unreleased]

### Added
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
