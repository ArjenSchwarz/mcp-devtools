---
references:
    - specs/copilot-agent-tool/requirements.md
    - specs/copilot-agent-tool/design.md
    - specs/copilot-agent-tool/decision_log.md
---
# GitHub Copilot Agent Tool - Implementation Tasks

## Foundation

- [x] 1. Create package structure and tool skeleton
  - Create internal/tools/copilotagent/ directory
  - Create copilot.go with CopilotTool struct
  - Implement tools.Tool interface (Definition and Execute methods)
  - Add package-level constants (DefaultTimeout, DefaultMaxResponseSize, environment variable names)
  - Requirements: [1.1](requirements.md#1.1), [1.2](requirements.md#1.2), [1.6](requirements.md#1.6), [1.7](requirements.md#1.7)
  - References: internal/tools/qdeveloperagent/qdeveloper.go, internal/tools/tools.go
  - [x] 1.1. Implement tool registration
    - Add init() function with registry.Register(&CopilotTool{})
    - Add import to internal/imports/tools.go
    - Requirements: [1.2](requirements.md#1.2), [15.4](requirements.md#15.4)
    - References: internal/registry/registry.go, internal/imports/tools.go
  - [x] 1.2. Create MCP tool definition
    - Implement Definition() method returning mcp.Tool
    - Define all parameters (prompt required, 8 optional parameters)
    - Add parameter descriptions and type specifications
    - Include MCP hint annotations (read-only=false, destructive=true, idempotent=false, open-world=true)
    - Requirements: [2.1](requirements.md#2.1), [3.1](requirements.md#3.1), [4.1](requirements.md#4.1), [4.3](requirements.md#4.3), [5.1](requirements.md#5.1), [5.3](requirements.md#5.3), [5.5](requirements.md#5.5), [6.1](requirements.md#6.1), [7.1](requirements.md#7.1), [11.1](requirements.md#11.1), [11.2](requirements.md#11.2), [11.3](requirements.md#11.3), [11.4](requirements.md#11.4), [11.5](requirements.md#11.5)
    - References: internal/tools/qdeveloperagent/qdeveloper.go

## Core Implementation

- [ ] 2. Implement core execution logic
  - Implement Execute() method with tool enablement check
  - Add prompt parameter validation (non-empty check)
  - Parse all optional parameters from args map
  - Handle timeout configuration from environment
  - Return appropriate errors for validation failures
  - Requirements: [1.3](requirements.md#1.3), [1.4](requirements.md#1.4), [2.1](requirements.md#2.1), [2.2](requirements.md#2.2), [2.4](requirements.md#2.4), [9.1](requirements.md#9.1), [9.2](requirements.md#9.2), [10.2](requirements.md#10.2), [16.3](requirements.md#16.3)
  - References: internal/tools/enablement.go
  - [ ] 2.1. Implement runCopilot helper method
    - Create runCopilot() method with context, timeout, prompt, and args parameters
    - Build command arguments array inline (copilot -p prompt --no-color)
    - Add model selection (--model flag)
    - Add session management (--continue for resume, --resume for session-id with priority)
    - Add permission management (--allow-all-tools for yolo-mode)
    - Add array parameters (--allow-tool, --deny-tool, --add-dir, --disable-mcp-server)
    - Execute command with exec.CommandContext
    - Capture stdout and stderr
    - Requirements: [1.4](requirements.md#1.4), [1.7](requirements.md#1.7), [2.3](requirements.md#2.3), [3.2](requirements.md#3.2), [3.3](requirements.md#3.3), [4.2](requirements.md#4.2), [4.4](requirements.md#4.4), [4.5](requirements.md#4.5), [5.2](requirements.md#5.2), [5.4](requirements.md#5.4), [5.6](requirements.md#5.6), [5.7](requirements.md#5.7), [6.2](requirements.md#6.2), [6.3](requirements.md#6.3), [7.2](requirements.md#7.2), [7.3](requirements.md#7.3)
    - References: internal/tools/qdeveloperagent/qdeveloper.go
  - [ ] 2.2. Implement error handling
    - Check for context.DeadlineExceeded and return partial output with timeout message
    - Detect CLI not found (check stderr for command not found or executable file not found)
    - Detect authentication failures (check stderr for not authenticated or authentication)
    - Return descriptive errors with stderr included
    - Handle all error scenarios from requirements
    - Requirements: [9.3](requirements.md#9.3), [10.1](requirements.md#10.1), [10.2](requirements.md#10.2), [10.3](requirements.md#10.3), [10.4](requirements.md#10.4), [10.5](requirements.md#10.5), [10.6](requirements.md#10.6)
    - References: internal/tools/codexagent/codex.go

## Output Handling

- [ ] 3. Implement output processing
  - Create filterOutput() method
  - Filter progress indicators (●, ✓, ✗, ↪ characters)
  - Filter command execution traces (lines starting with $)
  - Stop processing at Total usage est
  - Collapse multiple consecutive empty lines
  - Return cleaned output
  - Requirements: [8.1](requirements.md#8.1), [8.2](requirements.md#8.2)
  - References: internal/tools/geminiagent/gemini.go
  - [ ] 3.1. Implement response size limiting
    - Create GetMaxResponseSize() method reading from AGENT_MAX_RESPONSE_SIZE
    - Create ApplyResponseSizeLimit() method
    - Truncate at line boundary within last 100 chars when possible
    - Append truncation message with original and truncated sizes
    - Apply size limit before returning result
    - Requirements: [9.4](requirements.md#9.4), [9.5](requirements.md#9.5), [9.6](requirements.md#9.6), [9.7](requirements.md#9.7)
    - References: internal/tools/geminiagent/gemini.go
  - [ ] 3.2. Implement timeout helper
    - Create GetTimeout() method
    - Read AGENT_TIMEOUT environment variable
    - Return configured value or DefaultTimeout (180 seconds)
    - Validate timeout is positive integer
    - Requirements: [9.1](requirements.md#9.1), [9.2](requirements.md#9.2)
    - References: internal/tools/qdeveloperagent/qdeveloper.go

## Testing

- [ ] 4. Create unit tests
  - Create tests/tools/copilot_agent_test.go
  - Test Definition() method returns valid tool definition
  - Test missing prompt parameter returns error
  - Test empty/whitespace prompt returns error
  - Test GetTimeout() with and without environment variable
  - Test GetMaxResponseSize() with and without environment variable
  - Test ApplyResponseSizeLimit() truncates large output
  - Test filterOutput() removes progress indicators, commands, and usage stats
  - Test command construction for all parameter combinations (basic, model, resume, session-id, yolo-mode, arrays)
  - Requirements: [13.1](requirements.md#13.1), [13.2](requirements.md#13.2), [13.3](requirements.md#13.3), [13.4](requirements.md#13.4)
  - References: tests/tools/qdeveloper_test.go, tests/tools/geminiagent_test.go

## Documentation & Polish

- [ ] 5. Implement extended help
  - Implement ProvideExtendedInfo() method
  - Add 5+ realistic usage examples with arguments and expected results
  - Document common usage patterns
  - Add troubleshooting tips (CLI not found, authentication, tool not enabled)
  - Document parameter details for all parameters
  - Add WhenToUse and WhenNotToUse guidance
  - Requirements: [12.1](requirements.md#12.1), [12.2](requirements.md#12.2), [12.3](requirements.md#12.3), [12.4](requirements.md#12.4), [12.5](requirements.md#12.5), [12.6](requirements.md#12.6)
  - References: internal/tools/qdeveloperagent/qdeveloper.go, internal/tools/claudeagent/claude.go
  - [ ] 5.1. Create user documentation
    - Create docs/tools/copilot-agent.md
    - Document tool purpose and capabilities
    - Document all parameters with descriptions and examples
    - Document ENABLE_ADDITIONAL_TOOLS requirement
    - Document environment variables (AGENT_TIMEOUT, AGENT_MAX_RESPONSE_SIZE)
    - Provide usage examples for common scenarios
    - Document security considerations and permission management
    - Add authentication setup instructions (gh auth login)
    - Requirements: [14.1](requirements.md#14.1), [14.2](requirements.md#14.2), [14.3](requirements.md#14.3), [14.4](requirements.md#14.4), [14.5](requirements.md#14.5)
    - References: docs/tools/q-developer-agent.md, docs/tools/claude-agent.md
  - [ ] 5.2. Update README and overview documentation
    - Add copilot-agent to Agents table in README.md (around line 141)
    - Add copilot-agent to docs/tools/overview.md agents section
    - Add copilot-agent to ENABLE_ADDITIONAL_TOOLS examples in overview.md
    - Requirements: [14.6](requirements.md#14.6), [14.7](requirements.md#14.7)
    - References: README.md, docs/tools/overview.md
  - [ ] 5.3. Update tool enablement documentation
    - Add copilot-agent to supported tool names list in internal/tools/enablement.go comments
    - Requirements: [16.3](requirements.md#16.3)
    - References: internal/tools/enablement.go

## Quality Assurance

- [ ] 6. Run linting and tests
  - Run make lint to check code quality
  - Run make test to verify all tests pass
  - Run go mod tidy to clean dependencies
  - Fix any linting or formatting issues
  - Ensure test coverage meets standards
  - Requirements: [15.1](requirements.md#15.1), [15.2](requirements.md#15.2), [15.3](requirements.md#15.3), [15.5](requirements.md#15.5), [15.6](requirements.md#15.6)
  - References: Makefile

## Integration

- [ ] 7. Build and verify integration
  - Run make build to compile binary
  - Verify copilot-agent tool appears in tool list
  - Test tool registration works correctly
  - Verify tool is disabled by default
  - Test enabling with ENABLE_ADDITIONAL_TOOLS=copilot-agent
  - Requirements: [1.2](requirements.md#1.2), [1.3](requirements.md#1.3), [16.1](requirements.md#16.1), [16.2](requirements.md#16.2), [16.4](requirements.md#16.4)
  - References: Makefile
