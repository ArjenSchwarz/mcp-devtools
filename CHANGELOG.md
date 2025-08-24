# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed
- Modernised Q Developer agent code to use `any` instead of `interface{}` type constraint
- Improved Q Developer agent test code using `slices.Contains()` instead of manual loops
- Standardised error message casing for q Developer references
- Updated README.md with proper emoji for Q Developer agent (🅰️)

### Added
- Q Developer Agent tool for AWS Q Developer CLI integration
- Complete tool implementation with MCP protocol support
- Parameter support for resume, agent, override-model, yolo-mode, trust-tools, and verbose options
- Environment-based configuration and security controls
- Comprehensive test suite for Q Developer agent functionality
- Feature documentation including requirements, design, tasks, and decision log
- Comprehensive unit tests for Q Developer Agent command building logic including:
  - Basic command construction with prompt only
  - Command building with all parameters
  - Verification of --no-interactive flag inclusion
  - Parameter mapping to CLI flags
  - Empty optional parameter handling
  - Command injection prevention testing
- Comprehensive unit test suite for Q Developer agent response processing functionality
- Enhanced response size limit handling with proper line boundary truncation
- Support for multiple line ending formats (Unix, Windows, mixed) in response processing

### Changed
- Updated task tracking for Q Developer Agent implementation, marking command building and helper function tasks as completed
- Improved `ApplyResponseSizeLimit()` function to truncate at line boundaries rather than mid-line
- Enhanced truncation message format to show size limit in appropriate units (B/KB/MB)
- Updated truncation message to include both original and truncated response sizes

### Fixed
- Empty response handling in `ApplyResponseSizeLimit()` function
- Response truncation now preserves complete lines when possible
- Q Developer Agent test suite performance - reduced execution time from 2+ minutes to 1.3 seconds (90%+ improvement)
- Q Developer Agent command building tests now work regardless of CLI installation status
- Test reliability by removing dependency on external Q Developer CLI execution

### Completed
- Task 6: Timeout Management for Q Developer Agent including comprehensive unit tests and Execute() method integration
- Q Developer Agent implementation with comprehensive unit test suite (300+ test cases)
- Tool registration and integration with MCP registry
- Documentation updates for Q Developer Agent in README.md and docs/tools/overview.md
- All Q Developer Agent implementation tasks marked as completed