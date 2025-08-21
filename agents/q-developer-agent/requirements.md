# Q Developer Agent Tool Requirements

## Introduction

The Q Developer Agent tool provides integration with AWS Q Developer CLI through the MCP (Model Context Protocol) server. This tool enables AI agents to leverage Q Developer's capabilities for code analysis, generation, and assistance, similar to how the existing Gemini and Claude agent tools integrate with their respective CLI tools.

## Requirements

### 1. Core CLI Integration

**User Story:** As a developer, I want to invoke Q Developer through the MCP server, so that I can leverage AWS Q's AI capabilities for code assistance and analysis.

**Acceptance Criteria:**
1.1. The system SHALL execute the Q Developer CLI using the `q chat` command with appropriate arguments
1.2. The system SHALL always include the `--no-interactive` flag to ensure Q Developer runs without expecting user input
1.3. The system SHALL capture and return the complete output from Q Developer's response
1.4. The system SHALL handle CLI execution errors gracefully and return meaningful error messages
1.5. The system SHALL apply configurable timeout limits to prevent hanging operations
1.6. The system SHALL execute Q Developer in the current working directory without changing it
1.7. The system SHALL NOT verify Q Developer CLI installation (matching Claude/Gemini behaviour)

### 2. Prompt Management

**User Story:** As a developer, I want to send detailed prompts to Q Developer, so that I can get specific assistance with my code tasks.

**Acceptance Criteria:**
2.1. The system SHALL accept a required "prompt" parameter containing the instruction for Q Developer
2.2. The system SHALL pass the prompt as the first argument to the q chat command
2.3. The system SHALL validate that prompts are non-empty before execution
2.4. The system SHALL NOT support file/directory context references (Q Developer does not support @ syntax)

### 3. Configuration Options

**User Story:** As a developer, I want to configure Q Developer's behaviour through tool parameters, so that I can customise its operation for different use cases.

**Acceptance Criteria:**
3.1. The system SHALL support a `resume` parameter that maps to the `--resume` flag to continue previous conversations from the current directory
3.2. The system SHALL support an `agent` parameter that maps to the `--agent` flag to specify context profiles
3.3. The system SHALL support an `override-model` parameter that maps to the `--model` flag with available models: claude-3.5-sonnet, claude-3.7-sonnet, and claude-sonnet-4 (default)
3.4. The system SHALL support a `yolo-mode` parameter that maps to the `--trust-all-tools` flag (matching naming convention from Claude/Gemini agents)
3.5. The system SHALL support a `trust-tools` parameter that maps to the `--trust-tools` flag to specify trusted tools
3.6. The system SHALL support a `verbose` parameter that maps to the `--verbose` flag for detailed logging
3.7. The system SHALL always set `--no-interactive` flag regardless of user parameters

### 4. Security and Access Control

**User Story:** As a system administrator, I want to control access to the Q Developer agent tool, so that I can manage security and resource usage.

**Acceptance Criteria:**
4.1. The system SHALL check if the q-developer-agent tool is explicitly enabled via environment variables
4.2. The system SHALL require the tool to be included in ENABLE_ADDITIONAL_TOOLS environment variable
4.3. The system SHALL respect file system permissions when Q Developer attempts to read or modify files
4.4. The system SHALL provide clear error messages when the tool is not enabled
4.5. The system SHALL NOT handle AWS authentication - users are responsible for Q Developer authentication
4.6. The system SHALL NOT include authentication error handling beyond basic error reporting

### 5. Response Management

**User Story:** As a developer, I want to receive complete but manageable responses from Q Developer, so that I can process the output effectively.

**Acceptance Criteria:**
5.1. The system SHALL implement a configurable maximum response size limit (default 2MB)
5.2. The system SHALL truncate responses that exceed the size limit at the last line break within 100 characters of the limit
5.3. The system SHALL append a truncation message indicating original and truncated sizes when responses are truncated
5.4. The system SHALL allow the response size limit to be configured via AGENT_MAX_RESPONSE_SIZE environment variable
5.5. The system SHALL log warnings when responses are truncated
5.6. The system SHALL return plain text responses without additional formatting

### 6. Timeout Handling

**User Story:** As a developer, I want operations to complete within reasonable time limits, so that I don't experience hanging operations.

**Acceptance Criteria:**
6.1. The system SHALL implement a configurable timeout for Q Developer operations (default 180 seconds)
6.2. The system SHALL allow timeout configuration via AGENT_TIMEOUT environment variable
6.3. The system SHALL gracefully handle timeout situations and return partial output if available
6.4. The system SHALL append a timeout notification message when operations exceed the time limit
6.5. The system SHALL properly clean up resources when timeouts occur

### 7. Tool Registration and Discovery

**User Story:** As an MCP client, I want to discover and use the Q Developer agent tool, so that I can integrate it into my workflow.

**Acceptance Criteria:**
7.1. The system SHALL register the tool with the MCP registry during initialisation
7.2. The system SHALL provide a complete tool definition including name, description, and parameters
7.3. The system SHALL implement the standard Tool interface with Definition() and Execute() methods
7.4. The system SHALL provide extended help information with examples and troubleshooting tips
7.5. The system SHALL follow the naming convention "q-developer-agent" for consistency

### 8. Error Handling and Logging

**User Story:** As a developer, I want clear error messages and logging, so that I can troubleshoot issues effectively.

**Acceptance Criteria:**
8.1. The system SHALL log tool execution with appropriate log levels (Info, Warning, Error)
8.2. The system SHALL capture both stdout and stderr from Q Developer CLI
8.3. The system SHALL provide meaningful error messages for common failure scenarios
8.4. The system SHALL avoid logging to stdout in stdio mode to prevent breaking the MCP protocol
8.5. The system SHALL include command details in debug logs for troubleshooting

