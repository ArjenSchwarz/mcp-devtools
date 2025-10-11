# Decision Log - GitHub Copilot Agent Tool

## 2025-10-08: Initial Requirements Decisions

### Decision: No Banner Support
**Context**: Copilot CLI has a `--banner` flag to show startup banner.
**Decision**: Do not support banner flag - not needed for non-interactive agent usage.
**Rationale**: Banner is visual fluff that adds no value in programmatic/agent context.

### Decision: Focus on Agent Execution Only
**Context**: Copilot CLI supports version flag and help topics.
**Decision**: Keep tool strictly focused on agent execution; do not expose version or help topics.
**Rationale**: Tool users can reference Copilot documentation directly; keep tool interface minimal and focused.

### Decision: Accept Permission Patterns as Strings
**Context**: Copilot supports complex permission patterns like `shell(git:*)`.
**Decision**: Accept permission patterns as strings and pass through directly without validation.
**Rationale**: Q Developer agent uses the same approach with `trust-tools` parameter. Consistent with existing patterns and allows flexibility for future Copilot pattern enhancements.

### Decision: Filter Copilot Output
**Context**: Copilot outputs usage statistics and metadata.
**Decision**: Filter out sections starting with "Total usage est" and other non-relevant metadata.
**Rationale**: Keep response clean and focused on AI's actual answer, similar to how Gemini agent filters startup messages.

### Decision: Always Use --no-color Flag
**Context**: Copilot supports colour output control.
**Decision**: Always pass --no-color flag; do not expose as user parameter.
**Rationale**: Colour codes add overhead and complexity in programmatic contexts; no value in agent-to-agent communication.

### Decision: No Session Information in Response
**Context**: Some agents (Claude, Codex) include session info in responses.
**Decision**: Keep output clean with just Copilot's response; no session metadata.
**Rationale**: User preference for clean, minimal output.

### Decision: No Copilot-Specific Environment Variables
**Context**: Copilot has environment variables like COPILOT_ALLOW_ALL.
**Decision**: Do not document or support Copilot-specific environment variables through tool parameters.
**Rationale**: Keep tool interface simple; users can set environment variables directly if needed.

### Decision: Use Standard Agent Environment Variables
**Context**: All agents use AGENT_TIMEOUT and AGENT_MAX_RESPONSE_SIZE.
**Decision**: Reuse existing environment variables for consistency.
**Rationale**: Maintains consistency across all agent tools; users have single configuration point.

## 2025-10-08: Post-Review Decisions

### Decision: Verified Copilot CLI Flags
**Context**: Design-critic agent questioned whether flags actually exist.
**Decision**: User verified all flags exist via `copilot --help` output.
**Rationale**: All proposed flags (`--model`, `--add-dir`, `--disable-mcp-server`, `--log-level`, `--log-dir`, `--screen-reader`, `-p/--prompt`, `--continue`, `--resume`, `--allow-all-tools`, `--allow-tool`, `--deny-tool`, `--no-color`) are confirmed to exist.

### Decision: Use Only yolo-mode Parameter
**Context**: Initial requirements had both yolo-mode and allow-all-tools.
**Decision**: Use only "yolo-mode" as parameter name, which maps to --allow-all-tools CLI flag.
**Rationale**: Consistent with project terminology; simpler interface without dual naming.

### Decision: Authentication Out of Scope
**Context**: Peer-review-validator suggested authentication verification requirements.
**Decision**: Authentication handling is not within scope, consistent with other agents.
**Rationale**: Users are expected to authenticate externally (via `gh auth login`); tool doesn't manage authentication state.

### Decision: Rate Limiting is Acceptable Risk
**Context**: Design-critic raised concerns about premium API abuse.
**Decision**: No rate limiting or usage controls in the tool.
**Rationale**: Acceptable risk for this use case; users manage their own API quotas.

### Decision: No Integration Tests
**Context**: Peer-review-validator recommended integration tests with stubbed CLI.
**Decision**: Only basic unit tests; no integration tests.
**Rationale**: Integration tests would be costly (premium API) and add complexity without sufficient value for this tool.

### Decision: Prompt Handling Matches Other Agents
**Context**: Security concerns about command injection in prompt parameter.
**Decision**: Pass prompt similar to other agents; acceptable risk level.
**Rationale**: Consistent with existing agent implementations; security model is inherited from those patterns.

### Decision: No Directory Path Validation
**Context**: Peer-review-validator suggested path traversal prevention.
**Decision**: Pass add-dir paths directly without validation; intentionally allows access outside project boundaries.
**Rationale**: The feature's purpose is to grant access beyond normal boundaries; validation would defeat the feature.

### Decision: No Process Cleanup Requirements
**Context**: Peer-review-validator suggested process group management for cleanup.
**Decision**: No special cleanup logic; command finishes and exits naturally.
**Rationale**: Copilot CLI is a short-lived command that completes execution; no need for complex process management.

### Decision: No Additional Observability
**Context**: Peer-review-validator suggested structured logging and metrics.
**Decision**: Match observability level of existing agents only.
**Rationale**: Consistency with codebase; no special requirements for this agent over others.

## 2025-10-08: Final Parameter Naming Decisions

### Decision: No Model Validation
**Context**: Initial requirements validated model values against specific list.
**Decision**: Pass model values directly to Copilot without validation.
**Rationale**: Avoid needing tool updates when Copilot adds new models; let Copilot handle validation.

### Decision: Use "resume" Instead of "continue"
**Context**: Copilot uses --continue and --resume flags.
**Decision**: Parameter name is "resume" (boolean), maps to --continue flag. Parameter "session-id" (string) maps to --resume flag.
**Rationale**: Matches naming convention from other agent tools (codex, q-developer use "resume").

### Decision: Use "include-directories" Instead of "add-dir"
**Context**: Copilot CLI flag is --add-dir.
**Decision**: Parameter name is "include-directories", maps to --add-dir flag.
**Rationale**: Matches naming convention from claude-agent tool for consistency.

### Decision: Remove Log-Level and Log-Dir Parameters
**Context**: Initial requirements included log-level and log-dir optional parameters.
**Decision**: Remove these parameters; no other agent tools expose logging configuration.
**Rationale**: Consistency with other agents; users can set environment variables if needed.

### Decision: Remove Screen-Reader Parameter
**Context**: Initial requirements included screen-reader accessibility parameter.
**Decision**: Remove screen-reader parameter.
**Rationale**: Tool output is parsed by AI agents, not displayed to humans; accessibility features not relevant.

### Decision: Match Error Handling Patterns
**Context**: Need consistent error handling across agents.
**Decision**: Detect CLI missing by checking for "command not found" or "executable file not found"; detect auth failures by checking for "not authenticated" or "authentication" in stderr; include stderr in error messages.
**Rationale**: Exact pattern match with q-developer and codex agents for consistency.

## 2025-10-09: Design Phase Decisions

### Decision: Updated Output Filtering Based on Actual Copilot Output
**Context**: Received actual Copilot CLI output showing progress indicators (●, ✓, ✗, ↪), command traces ($ prefix), and usage statistics.
**Decision**: Filter output to remove:
- Progress indicators (Unicode characters ●, ✓, ✗, ↪)
- Command execution traces (lines starting with $)
- Usage statistics (stop processing at "Total usage est")
- Collapse multiple consecutive empty lines
**Rationale**: Based on real Copilot output format, ensures clean AI-to-AI communication without metadata overhead.

### Decision: Documentation Locations
**Context**: Need to document tool in multiple locations like other agents.
**Decision**: Add documentation to:
- docs/tools/copilot-agent.md (detailed user docs)
- README.md agent tools table
- docs/tools/overview.md agents section
- internal/imports/tools.go import list
**Rationale**: Consistent with how other agent tools are documented and registered.

## 2025-10-09: Integration Testing Discoveries

### Discovery: Actual Copilot Output Format Differs from Design Assumptions
**Context**: During integration testing with actual Copilot CLI, discovered that output format differs significantly from initial assumptions.
**Observation**:
- Initial design assumed progress indicators were on separate lines from content
- Actual Copilot output: `● 4` (answer on same line as progress indicator)
- Initial filtering removed entire lines starting with progress indicators, thus removing the answer
**Impact**: Tool returns empty responses for simple queries because actual answers are filtered out.

### Decision: Revise Output Filtering Strategy
**Context**: Current filter removes lines starting with progress indicators, but Copilot places answers after the last progress indicator (sometimes on same line).
**Decision**: Implement revised filtering strategy:
1. Find the last progress indicator (●, ✓, ✗, ↪) in the output
2. Extract content from after that indicator (including content on the same line)
3. Continue extracting until "Total usage est" appears
4. Clean and return the extracted content
**Rationale**:
- Copilot's output pattern shows the actual AI answer follows the last progress indicator
- All content before the last progress indicator is metadata and execution traces
- This approach captures the actual response whilst discarding metadata
**References**: Updated design.md Output Filtering section with revised implementation

### Decision: Add Bug Fix Task for Filtering
**Context**: Need to implement the revised filtering strategy to fix empty response issue.
**Decision**: Add task 8 to tasks.md with subtasks:
- 8.1: Update FilterOutput() implementation
- 8.2: Update filter tests to match new behaviour
- 8.3: Run integration verification with actual Copilot CLI
**Rationale**: Structured approach to fixing the filtering bug whilst maintaining test coverage and verification.
