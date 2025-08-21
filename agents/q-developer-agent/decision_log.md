# Q Developer Agent Tool - Decision Log

## Feature Name Decision
- **Date**: 2025-08-20
- **Decision**: Use "q-developer-agent" as the feature name
- **Rationale**: Follows the existing naming pattern of other agent tools (gemini-agent, claude-agent) and clearly identifies the tool's purpose
- **Proposed by**: Assistant
- **Approved by**: User

## CLI Flag Correction
- **Date**: 2025-08-20
- **Decision**: Use `--no-interactive` flag instead of `--non-interactive`
- **Rationale**: Corrected based on actual Q Developer CLI help output
- **Proposed by**: User
- **Status**: Implemented

## Context Reference Support
- **Date**: 2025-08-20
- **Decision**: Q Developer does NOT support @ syntax for file/directory references
- **Rationale**: User confirmed it's not supported at this time
- **Proposed by**: User
- **Status**: Documented in requirements

## Authentication Handling
- **Date**: 2025-08-20
- **Decision**: Authentication is out of scope - users responsible for ensuring Q Developer is authenticated before use
- **Rationale**: Follows separation of concerns - tool focuses on MCP integration, not auth management
- **Proposed by**: User
- **Status**: Documented in requirements

## Available Models
- **Date**: 2025-08-20
- **Decision**: Support claude-3.5-sonnet, claude-3.7-sonnet, and claude-sonnet-4 (default)
- **Rationale**: These are the available models in Q Developer CLI
- **Proposed by**: User
- **Status**: Documented in requirements

## Model Parameter Naming
- **Date**: 2025-08-20
- **Decision**: Use `override-model` parameter name instead of `model`
- **Rationale**: Maintain consistency across all agent tools - repo owner decided to standardize on `override-model`
- **Proposed by**: User (after consultation with repo owner)
- **Status**: Updated in requirements

## Trust Tools Configuration
- **Date**: 2025-08-20
- **Decision**: Use `yolo-mode` parameter name (matching Claude/Gemini convention) for `--trust-all-tools` flag
- **Rationale**: Maintain consistency with existing agent tools
- **Proposed by**: User
- **Status**: Updated in requirements

## Output Format
- **Date**: 2025-08-20
- **Decision**: Q Developer returns plain text with no special formatting required
- **Rationale**: User confirmed output format
- **Proposed by**: User
- **Status**: Documented in requirements

## CLI Installation Verification
- **Date**: 2025-08-20
- **Decision**: Do NOT verify Q Developer CLI installation (matching Claude/Gemini behaviour)
- **Rationale**: Analysis showed existing agent tools don't verify installation - they rely on command execution failure
- **Research**: Assistant verified Claude/Gemini implementation patterns
- **Status**: Documented in requirements

## Working Directory Behaviour
- **Date**: 2025-08-20
- **Decision**: Execute Q Developer in current working directory without changing it
- **Rationale**: Matches existing Claude/Gemini agent behaviour
- **Research**: Assistant verified existing implementation patterns
- **Status**: Documented in requirements

## Session Management
- **Date**: 2025-08-21
- **Decision**: Q Developer uses directory-based session management
- **Rationale**: The `--resume` flag "Resumes the previous conversation from this directory" as confirmed by documentation
- **Implementation**: The tool will support the `resume` parameter which maps to `--resume` flag, enabling continuation of previous conversations stored in the current directory
- **Proposed by**: User (via documentation)
- **Status**: Documented in requirements