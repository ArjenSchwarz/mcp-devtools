package tools_test

import (
	"os"
	"strings"
	"testing"

	"github.com/sammcj/mcp-devtools/internal/tools/qdeveloperagent"
	"github.com/sammcj/mcp-devtools/tests/testutils"
	"github.com/stretchr/testify/assert"
)

// Basic tests following the pattern of other agent tools (geminiagent_test.go, claudeagent_test.go)

func TestQDeveloperTool_Definition(t *testing.T) {
	tool := &qdeveloperagent.QDeveloperTool{}
	def := tool.Definition()

	assert.NotNil(t, def)
	assert.Equal(t, "q-developer-agent", def.GetName())
}

func TestQDeveloperTool_Definition_ParameterSchema(t *testing.T) {
	tool := &qdeveloperagent.QDeveloperTool{}
	def := tool.Definition()

	// Test basic definition properties
	testutils.AssertEqual(t, "q-developer-agent", def.Name)
	testutils.AssertNotNil(t, def.Description)

	// Test that description contains key phrases
	desc := def.Description
	if !testutils.Contains(desc, "AWS Q Developer CLI") {
		t.Errorf("Expected description to contain 'AWS Q Developer CLI', got: %s", desc)
	}

	// Test input schema exists
	testutils.AssertNotNil(t, def.InputSchema)

	// Test that input schema has required properties
	schema := def.InputSchema
	testutils.AssertNotNil(t, schema.Properties)

	// Verify required prompt parameter exists
	promptProp, hasPrompt := schema.Properties["prompt"]
	testutils.AssertTrue(t, hasPrompt)
	testutils.AssertNotNil(t, promptProp)

	// Verify prompt is in required array
	testutils.AssertNotNil(t, schema.Required)
	promptRequired := false
	for _, required := range schema.Required {
		if required == "prompt" {
			promptRequired = true
			break
		}
	}
	testutils.AssertTrue(t, promptRequired)
}

func TestQDeveloperTool_Definition_OptionalParameters(t *testing.T) {
	tool := &qdeveloperagent.QDeveloperTool{}
	def := tool.Definition()
	schema := def.InputSchema

	// Test optional parameters exist
	optionalParams := []string{"resume", "agent", "override-model", "yolo-mode", "trust-tools", "verbose"}

	for _, param := range optionalParams {
		prop, exists := schema.Properties[param]
		if !exists {
			t.Errorf("Expected optional parameter '%s' to exist in schema", param)
			continue
		}
		testutils.AssertNotNil(t, prop)
	}

	// Verify none of the optional parameters are in required array
	for _, param := range optionalParams {
		for _, required := range schema.Required {
			if required == param {
				t.Errorf("Optional parameter '%s' should not be in required array", param)
			}
		}
	}
}

func TestQDeveloperTool_Definition_ParameterNamingConventions(t *testing.T) {
	tool := &qdeveloperagent.QDeveloperTool{}
	def := tool.Definition()
	schema := def.InputSchema

	// Test that we use consistent naming conventions
	expectedParams := map[string]bool{
		"prompt":         true, // required
		"resume":         true, // optional boolean
		"agent":          true, // optional string
		"override-model": true, // follows decision log standardization
		"yolo-mode":      true, // matches Claude/Gemini convention
		"trust-tools":    true, // optional string
		"verbose":        true, // optional boolean
	}

	// Verify we have exactly these parameters (no more, no less)
	for param := range schema.Properties {
		_, expected := expectedParams[param]
		if !expected {
			t.Errorf("Unexpected parameter found: %s", param)
		}
	}

	for expectedParam := range expectedParams {
		_, exists := schema.Properties[expectedParam]
		if !exists {
			t.Errorf("Expected parameter missing: %s", expectedParam)
		}
	}
}

// Configuration tests (these are fast and don't execute CLI)

func TestQDeveloperTool_TimeoutConfiguration_DefaultTimeout(t *testing.T) {
	// Save original environment variable
	originalValue := os.Getenv("AGENT_TIMEOUT")
	defer func() {
		if originalValue == "" {
			_ = os.Unsetenv("AGENT_TIMEOUT")
		} else {
			_ = os.Setenv("AGENT_TIMEOUT", originalValue)
		}
	}()

	// Clear environment variable to test default behaviour
	_ = os.Unsetenv("AGENT_TIMEOUT")

	tool := &qdeveloperagent.QDeveloperTool{}
	timeout := tool.GetTimeout()

	testutils.AssertEqual(t, qdeveloperagent.DefaultTimeout, timeout)
}

func TestQDeveloperTool_TimeoutConfiguration_CustomTimeout(t *testing.T) {
	// Save original environment variable
	originalValue := os.Getenv("AGENT_TIMEOUT")
	defer func() {
		if originalValue == "" {
			_ = os.Unsetenv("AGENT_TIMEOUT")
		} else {
			_ = os.Setenv("AGENT_TIMEOUT", originalValue)
		}
	}()

	// Set custom timeout
	customTimeout := "300"
	_ = os.Setenv("AGENT_TIMEOUT", customTimeout)

	tool := &qdeveloperagent.QDeveloperTool{}
	timeout := tool.GetTimeout()

	testutils.AssertEqual(t, 300, timeout)
}

func TestQDeveloperTool_TimeoutConfiguration_InvalidTimeout(t *testing.T) {
	// Save original environment variable
	originalValue := os.Getenv("AGENT_TIMEOUT")
	defer func() {
		if originalValue == "" {
			_ = os.Unsetenv("AGENT_TIMEOUT")
		} else {
			_ = os.Setenv("AGENT_TIMEOUT", originalValue)
		}
	}()

	// Set invalid timeout value
	_ = os.Setenv("AGENT_TIMEOUT", "not-a-number")

	tool := &qdeveloperagent.QDeveloperTool{}
	timeout := tool.GetTimeout()

	// Should fall back to default when invalid value is provided
	testutils.AssertEqual(t, qdeveloperagent.DefaultTimeout, timeout)
}

func TestQDeveloperTool_ResponseSizeLimit_DefaultLimit(t *testing.T) {
	// Save original environment variable
	originalValue := os.Getenv("AGENT_MAX_RESPONSE_SIZE")
	defer func() {
		if originalValue == "" {
			_ = os.Unsetenv("AGENT_MAX_RESPONSE_SIZE")
		} else {
			_ = os.Setenv("AGENT_MAX_RESPONSE_SIZE", originalValue)
		}
	}()

	// Clear environment variable to test default behaviour
	_ = os.Unsetenv("AGENT_MAX_RESPONSE_SIZE")

	tool := &qdeveloperagent.QDeveloperTool{}
	logger := testutils.CreateTestLogger()

	// Test with small output (should not be truncated)
	smallOutput := "This is a small response that should not be truncated."
	result := tool.ApplyResponseSizeLimit(smallOutput, logger)
	testutils.AssertEqual(t, smallOutput, result)

	// Test with large output (should be truncated)
	largeOutput := strings.Repeat("Q", 3*1024*1024) // 3MB
	result = tool.ApplyResponseSizeLimit(largeOutput, logger)

	// Should be truncated to default 2MB limit
	testutils.AssertTrue(t, len(result) < len(largeOutput))
	testutils.AssertTrue(t, strings.Contains(result, "[RESPONSE TRUNCATED"))
	testutils.AssertTrue(t, strings.Contains(result, "exceeded 2.0MB limit"))
}

func TestQDeveloperTool_ResponseSizeLimit_CustomLimit(t *testing.T) {
	// Save original environment variable
	originalValue := os.Getenv("AGENT_MAX_RESPONSE_SIZE")
	defer func() {
		if originalValue == "" {
			_ = os.Unsetenv("AGENT_MAX_RESPONSE_SIZE")
		} else {
			_ = os.Setenv("AGENT_MAX_RESPONSE_SIZE", originalValue)
		}
	}()

	// Set custom limit (1MB = 1048576 bytes)
	_ = os.Setenv("AGENT_MAX_RESPONSE_SIZE", "1048576")

	tool := &qdeveloperagent.QDeveloperTool{}
	logger := testutils.CreateTestLogger()

	// Test with output larger than custom limit
	largeOutput := strings.Repeat("Q", 1500000) // 1.5MB
	result := tool.ApplyResponseSizeLimit(largeOutput, logger)

	// Should be truncated to custom 1MB limit
	testutils.AssertTrue(t, len(result) < len(largeOutput))
	testutils.AssertTrue(t, strings.Contains(result, "[RESPONSE TRUNCATED"))
	testutils.AssertTrue(t, strings.Contains(result, "exceeded 1.0MB limit"))
}

func TestQDeveloperTool_Constants(t *testing.T) {
	// Test that constants are exported and have expected values
	testutils.AssertEqual(t, "AGENT_MAX_RESPONSE_SIZE", qdeveloperagent.AgentMaxResponseSizeEnvVar)
	testutils.AssertEqual(t, "AGENT_TIMEOUT", qdeveloperagent.AgentTimeoutEnvVar)
	testutils.AssertEqual(t, 2*1024*1024, qdeveloperagent.DefaultMaxResponseSize)
	testutils.AssertEqual(t, 180, qdeveloperagent.DefaultTimeout)
}

// Note: No Execute() tests - these would be slow as they execute real CLI.
// Integration tests can be added separately if needed for CI with proper CLI setup.
