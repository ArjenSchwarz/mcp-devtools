# GitHub Copilot Agent Tool Requirements

## Introduction

This feature adds a new MCP tool that integrates GitHub Copilot CLI as an AI agent, similar to the existing q-developer-agent, gemini-agent, codex-agent, and claude-agent tools. The tool will enable AI coding assistants to leverage GitHub Copilot's capabilities for code analysis, generation, and assistance through the Copilot CLI's non-interactive mode.

The tool will follow the established patterns of existing agent tools while supporting Copilot-specific features such as permission management, model selection, session management, directory access control, and MCP server configuration.

## Requirements

### 1. Core Tool Implementation

**User Story:** As a developer using MCP DevTools, I want to invoke GitHub Copilot as a sub-agent, so that I can leverage Copilot's AI capabilities for code analysis and generation tasks.

**Acceptance Criteria:**

1. <a name="1.1"></a>The system SHALL implement a CopilotTool struct that conforms to the tools.Tool interface
2. <a name="1.2"></a>The system SHALL register the tool with the name "copilot-agent" during package initialisation
3. <a name="1.3"></a>The system SHALL require the tool to be explicitly enabled via ENABLE_ADDITIONAL_TOOLS environment variable
4. <a name="1.4"></a>The system SHALL execute the copilot CLI command in non-interactive mode using the --prompt flag
5. <a name="1.5"></a>The system SHALL capture and return both stdout and stderr output from the Copilot CLI
6. <a name="1.6"></a>The system SHALL use the same variable naming conventions as existing agent tools (prompt, override-model, yolo-mode, etc.)
7. <a name="1.7"></a>The system SHALL always pass the --no-color flag to Copilot to reduce output overhead

### 2. Prompt Execution

**User Story:** As an AI agent, I want to send prompts directly to Copilot CLI, so that I can get code assistance without interactive sessions.

**Acceptance Criteria:**

1. <a name="2.1"></a>The system SHALL accept a required "prompt" parameter containing the instruction for Copilot
2. <a name="2.2"></a>The system SHALL validate that the prompt parameter is not empty or whitespace-only
3. <a name="2.3"></a>The system SHALL execute Copilot using the --prompt flag for direct prompt execution
4. <a name="2.4"></a>The system SHALL return an error if the prompt parameter is missing or invalid

### 3. Model Selection

**User Story:** As a user, I want to specify which AI model Copilot should use, so that I can choose between different models based on task requirements.

**Acceptance Criteria:**

1. <a name="3.1"></a>The system SHALL accept an optional "override-model" parameter to specify the AI model
2. <a name="3.2"></a>The system SHALL pass the model selection to Copilot using the --model flag when specified
3. <a name="3.3"></a>The system SHALL pass model values directly to Copilot without validation
4. <a name="3.4"></a>The system SHALL allow Copilot to use its configured default model when override-model is not provided

### 4. Session Management

**User Story:** As a user, I want to continue previous Copilot conversations, so that I can maintain context across multiple interactions.

**Acceptance Criteria:**

1. <a name="4.1"></a>The system SHALL accept an optional "resume" boolean parameter to resume the most recent session
2. <a name="4.2"></a>The system SHALL pass the --continue flag to Copilot when resume is true
3. <a name="4.3"></a>The system SHALL accept an optional "session-id" string parameter to resume a specific session by ID
4. <a name="4.4"></a>The system SHALL pass the --resume flag with the session ID to Copilot when session-id is provided and not empty
5. <a name="4.5"></a>The system SHALL prioritise session-id over resume when both are provided

### 5. Permission Management

**User Story:** As a user, I want to control what tools and operations Copilot can perform, so that I can balance automation with security.

**Acceptance Criteria:**

1. <a name="5.1"></a>The system SHALL accept an optional "yolo-mode" boolean parameter for automatic tool execution
2. <a name="5.2"></a>The system SHALL pass the --allow-all-tools flag to Copilot when yolo-mode is true
3. <a name="5.3"></a>The system SHALL accept an optional "allow-tool" array parameter for specific tool permissions
4. <a name="5.4"></a>The system SHALL pass each allowed tool using --allow-tool flags to Copilot
5. <a name="5.5"></a>The system SHALL accept an optional "deny-tool" array parameter for tool denials
6. <a name="5.6"></a>The system SHALL pass each denied tool using --deny-tool flags to Copilot
7. <a name="5.7"></a>The system SHALL pass permission pattern strings directly to Copilot without validation or parsing

### 6. Directory Access Control

**User Story:** As a user, I want to grant Copilot access to additional directories, so that it can work with files outside the current project.

**Acceptance Criteria:**

1. <a name="6.1"></a>The system SHALL accept an optional "include-directories" array parameter for additional directory access
2. <a name="6.2"></a>The system SHALL pass each additional directory using --add-dir flags to Copilot
3. <a name="6.3"></a>The system SHALL pass directory paths directly to Copilot without validation or boundary restrictions

### 7. MCP Server Configuration

**User Story:** As a user, I want to disable specific MCP servers during Copilot execution, so that I can avoid conflicts or unwanted tool availability.

**Acceptance Criteria:**

1. <a name="7.1"></a>The system SHALL accept an optional "disable-mcp-server" array parameter for disabling MCP servers
2. <a name="7.2"></a>The system SHALL pass each disabled server using --disable-mcp-server flags to Copilot
3. <a name="7.3"></a>The system SHALL support disabling multiple MCP servers in a single invocation

### 8. Output Filtering

**User Story:** As a user, I want clean, filtered output from Copilot without unnecessary overhead, so that I can focus on the AI's response.

**Acceptance Criteria:**

1. <a name="8.1"></a>The system SHALL filter out usage statistics sections starting with "Total usage est" from the output
2. <a name="8.2"></a>The system SHALL filter out any other Copilot-specific metadata or startup messages that are not relevant to the AI response

### 9. Timeout and Response Size Management

**User Story:** As a user, I want to control execution timeouts and response size limits, so that I can handle long-running operations and large outputs appropriately.

**Acceptance Criteria:**

1. <a name="9.1"></a>The system SHALL use a default timeout of 180 seconds (3 minutes) for Copilot execution
2. <a name="9.2"></a>The system SHALL read timeout configuration from AGENT_TIMEOUT environment variable
3. <a name="9.3"></a>The system SHALL return partial output with a timeout message when timeout is exceeded
4. <a name="9.4"></a>The system SHALL use a default maximum response size of 2MB
5. <a name="9.5"></a>The system SHALL read response size limit from AGENT_MAX_RESPONSE_SIZE environment variable
6. <a name="9.6"></a>The system SHALL truncate responses exceeding the size limit at a line boundary when possible
7. <a name="9.7"></a>The system SHALL append a truncation message indicating original and truncated sizes

### 10. Error Handling

**User Story:** As a user, I want clear error messages when Copilot execution fails, so that I can diagnose and resolve issues quickly.

**Acceptance Criteria:**

1. <a name="10.1"></a>The system SHALL detect when the copilot CLI is not installed by checking for "command not found" or "executable file not found" in error output
2. <a name="10.2"></a>The system SHALL return a descriptive error message when CLI is not found: "copilot CLI not found. Please install Copilot CLI and ensure it's available in your PATH"
3. <a name="10.3"></a>The system SHALL detect authentication failures by checking for "not authenticated" or "authentication" in stderr
4. <a name="10.4"></a>The system SHALL return a descriptive error message when authentication fails including the stderr output
5. <a name="10.5"></a>The system SHALL include stderr output in error messages when available for debugging
6. <a name="10.6"></a>The system SHALL handle context deadline exceeded errors gracefully with partial output

### 11. MCP Tool Definition

**User Story:** As an AI coding assistant, I want clear tool descriptions and parameter definitions, so that I can use the Copilot agent tool effectively.

**Acceptance Criteria:**

1. <a name="11.1"></a>The system SHALL provide a clear tool description indicating it provides GitHub Copilot CLI integration
2. <a name="11.2"></a>The system SHALL mark the prompt parameter as required in the tool definition
3. <a name="11.3"></a>The system SHALL provide descriptions for all optional parameters
4. <a name="11.4"></a>The system SHALL include MCP hint annotations: read-only=false, destructive=true, idempotent=false, open-world=true
5. <a name="11.5"></a>The system SHALL document expected parameter types and valid values in descriptions

### 12. Extended Help Information

**User Story:** As a user, I want detailed usage examples and troubleshooting guidance, so that I can understand how to use the Copilot agent effectively.

**Acceptance Criteria:**

1. <a name="12.1"></a>The system SHALL implement the ProvideExtendedInfo method returning ExtendedHelp structure
2. <a name="12.2"></a>The system SHALL provide at least 5 realistic usage examples with expected results
3. <a name="12.3"></a>The system SHALL document common usage patterns in the CommonPatterns field
4. <a name="12.4"></a>The system SHALL provide troubleshooting tips for common error scenarios
5. <a name="12.5"></a>The system SHALL document parameter details explaining each parameter's purpose
6. <a name="12.6"></a>The system SHALL provide WhenToUse and WhenNotToUse guidance

### 13. Testing Requirements

**User Story:** As a developer, I want basic tests for the Copilot agent tool, so that I can catch obvious regressions.

**Acceptance Criteria:**

1. <a name="13.1"></a>The system SHALL include unit tests in tests/tools/copilot_agent_test.go
2. <a name="13.2"></a>The system SHALL test parameter validation for the required prompt parameter
3. <a name="13.3"></a>The system SHALL test timeout behaviour and response size limiting
4. <a name="13.4"></a>The system SHALL test command argument construction for various parameter combinations

### 14. Documentation Requirements

**User Story:** As a user, I want clear documentation about the Copilot agent tool, so that I can understand its capabilities and how to use it.

**Acceptance Criteria:**

1. <a name="14.1"></a>The system SHALL include tool documentation in docs/tools/copilot-agent.md
2. <a name="14.2"></a>The system SHALL document all parameters with descriptions and examples
3. <a name="14.3"></a>The system SHALL document required environment variables and configuration
4. <a name="14.4"></a>The system SHALL provide usage examples for common scenarios
5. <a name="14.5"></a>The system SHALL document security considerations and permission management
6. <a name="14.6"></a>The system SHALL be listed in the main README.md tools section
7. <a name="14.7"></a>The system SHALL be listed in docs/tools/overview.md

### 15. Backwards Compatibility

**User Story:** As a maintainer, I want the new tool to follow existing patterns, so that it integrates seamlessly with the codebase.

**Acceptance Criteria:**

1. <a name="15.1"></a>The system SHALL follow the same code structure as existing agent tools
2. <a name="15.2"></a>The system SHALL use consistent variable naming conventions across agent tools (resume, session-id, include-directories)
3. <a name="15.3"></a>The system SHALL reuse environment variables (AGENT_TIMEOUT, AGENT_MAX_RESPONSE_SIZE) from other agents
4. <a name="15.4"></a>The system SHALL follow the same registration pattern in init() function
5. <a name="15.5"></a>The system SHALL use the same error handling patterns as other agent tools
6. <a name="15.6"></a>The system SHALL follow the same logging patterns using logrus.Logger

### 16. Security Considerations

**User Story:** As a security-conscious user, I want the Copilot agent tool to be disabled by default, so that it cannot be used without explicit enablement.

**Acceptance Criteria:**

1. <a name="16.1"></a>The system SHALL be disabled by default requiring ENABLE_ADDITIONAL_TOOLS environment variable
2. <a name="16.2"></a>The system SHALL check tool enablement before executing any Copilot commands
3. <a name="16.3"></a>The system SHALL return a clear error message when tool is not enabled: "copilot agent tool is not enabled. Set ENABLE_ADDITIONAL_TOOLS environment variable to include 'copilot-agent'"
4. <a name="16.4"></a>The system SHALL respect file system permissions even in yolo-mode
5. <a name="16.5"></a>The system SHALL document security implications of yolo-mode in extended help
