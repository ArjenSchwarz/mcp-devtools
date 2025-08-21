package tools_test

import (
	"os"
	"strings"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/sammcj/mcp-devtools/internal/tools/qdeveloperagent"
	"github.com/sammcj/mcp-devtools/tests/testutils"
	"github.com/stretchr/testify/assert"
)

func TestQDeveloperTool_Definition(t *testing.T) {
	tool := &qdeveloperagent.QDeveloperTool{}
	def := tool.Definition()

	assert.NotNil(t, def)
	assert.Equal(t, "q-developer-agent", def.GetName())
}

func TestQDeveloperTool_InitialStructure(t *testing.T) {
	tool := &qdeveloperagent.QDeveloperTool{}

	// Test that the tool struct can be instantiated
	assert.NotNil(t, tool)

	// Test that it implements the Tool interface by having required methods
	def := tool.Definition()
	assert.NotNil(t, def)
}

// Task 2.1: Unit tests for parameter validation and schema correctness
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

func TestQDeveloperTool_Definition_ParameterTypes(t *testing.T) {
	tool := &qdeveloperagent.QDeveloperTool{}
	def := tool.Definition()
	schema := def.InputSchema

	// Test string parameters
	stringParams := []string{"prompt", "agent", "override-model", "trust-tools"}
	for _, param := range stringParams {
		prop, exists := schema.Properties[param]
		testutils.AssertTrue(t, exists)

		// Check type (this is implementation-specific for mcp-go)
		propMap, ok := prop.(map[string]interface{})
		if ok {
			propType, hasType := propMap["type"]
			if hasType {
				testutils.AssertEqual(t, "string", propType)
			}
		}
	}

	// Test boolean parameters
	boolParams := []string{"resume", "yolo-mode", "verbose"}
	for _, param := range boolParams {
		prop, exists := schema.Properties[param]
		testutils.AssertTrue(t, exists)

		// Check type (this is implementation-specific for mcp-go)
		propMap, ok := prop.(map[string]interface{})
		if ok {
			propType, hasType := propMap["type"]
			if hasType {
				testutils.AssertEqual(t, "boolean", propType)
			}
		}
	}
}

func TestQDeveloperTool_Definition_ParameterDescriptions(t *testing.T) {
	tool := &qdeveloperagent.QDeveloperTool{}
	def := tool.Definition()
	schema := def.InputSchema

	// Test that all parameters have descriptions
	requiredDescriptions := map[string]string{
		"prompt":         "prompt",
		"resume":         "previous conversation",
		"agent":          "Context profile",
		"override-model": "model",
		"yolo-mode":      "trust-all-tools",
		"trust-tools":    "tools to trust",
		"verbose":        "verbose",
	}

	for param, expectedKeyword := range requiredDescriptions {
		prop, exists := schema.Properties[param]
		testutils.AssertTrue(t, exists)

		// Check description exists and contains expected keyword
		propMap, ok := prop.(map[string]interface{})
		if ok {
			description, hasDesc := propMap["description"]
			testutils.AssertTrue(t, hasDesc)
			descStr, isString := description.(string)
			testutils.AssertTrue(t, isString)
			testutils.AssertTrue(t, len(descStr) > 0)

			// Verify description contains expected keyword
			if !testutils.Contains(descStr, expectedKeyword) {
				t.Errorf("Expected parameter '%s' description to contain '%s', got: %s", param, expectedKeyword, descStr)
			}
		}
	}
}

func TestQDeveloperTool_Definition_ModelEnumValues(t *testing.T) {
	tool := &qdeveloperagent.QDeveloperTool{}
	def := tool.Definition()
	schema := def.InputSchema

	// Test that override-model parameter description mentions available models
	prop, exists := schema.Properties["override-model"]
	testutils.AssertTrue(t, exists)

	propMap, ok := prop.(map[string]interface{})
	testutils.AssertTrue(t, ok)

	description, hasDesc := propMap["description"]
	testutils.AssertTrue(t, hasDesc)
	descStr, isString := description.(string)
	testutils.AssertTrue(t, isString)

	// Check that all expected models are mentioned in description
	expectedModels := []string{"claude-3.5-sonnet", "claude-3.7-sonnet", "claude-sonnet-4"}
	for _, model := range expectedModels {
		if !testutils.Contains(descStr, model) {
			t.Errorf("Expected override-model description to mention model '%s', got: %s", model, descStr)
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

func TestQDeveloperTool_Definition_DefaultValues(t *testing.T) {
	tool := &qdeveloperagent.QDeveloperTool{}
	def := tool.Definition()
	schema := def.InputSchema

	// Test that boolean parameters have default values
	booleanDefaults := map[string]bool{
		"resume":    false,
		"yolo-mode": false,
		"verbose":   false,
	}

	for param, expectedDefault := range booleanDefaults {
		prop, exists := schema.Properties[param]
		testutils.AssertTrue(t, exists)

		propMap, ok := prop.(map[string]interface{})
		if ok {
			// Check if default is specified (implementation-specific)
			if defaultVal, hasDefault := propMap["default"]; hasDefault {
				testutils.AssertEqual(t, expectedDefault, defaultVal)
			}
		}
	}
}

// Task 3.1: Unit tests for environment variable handling

func TestQDeveloperTool_EnablementCheck_ToolNotEnabled(t *testing.T) {
	// Save original environment variable
	originalValue := os.Getenv("ENABLE_ADDITIONAL_TOOLS")
	defer func() {
		if originalValue == "" {
			_ = os.Unsetenv("ENABLE_ADDITIONAL_TOOLS")
		} else {
			_ = os.Setenv("ENABLE_ADDITIONAL_TOOLS", originalValue)
		}
	}()

	// Clear environment variable to test disabled tool
	_ = os.Unsetenv("ENABLE_ADDITIONAL_TOOLS")

	tool := &qdeveloperagent.QDeveloperTool{}
	ctx := testutils.CreateTestContext()
	logger := testutils.CreateTestLogger()
	cache := testutils.CreateTestCache()

	args := map[string]interface{}{
		"prompt": "Test prompt",
	}

	result, err := tool.Execute(ctx, logger, cache, args)

	testutils.AssertError(t, err)
	testutils.AssertErrorContains(t, err, "Q Developer agent tool is not enabled")
	testutils.AssertErrorContains(t, err, "ENABLE_ADDITIONAL_TOOLS")
	testutils.AssertErrorContains(t, err, "q-developer-agent")
	testutils.AssertEqual(t, (*mcp.CallToolResult)(nil), result)
}

func TestQDeveloperTool_EnablementCheck_ToolEnabled(t *testing.T) {
	// Save original environment variable
	originalValue := os.Getenv("ENABLE_ADDITIONAL_TOOLS")
	defer func() {
		if originalValue == "" {
			_ = os.Unsetenv("ENABLE_ADDITIONAL_TOOLS")
		} else {
			_ = os.Setenv("ENABLE_ADDITIONAL_TOOLS", originalValue)
		}
	}()

	// Enable the tool
	_ = os.Setenv("ENABLE_ADDITIONAL_TOOLS", "q-developer-agent")

	tool := &qdeveloperagent.QDeveloperTool{}

	// Test that the tool is now enabled (should not return enablement error)
	// Note: This test would need CLI mocking for full execution, but we can verify enablement check passes
	ctx := testutils.CreateTestContext()
	logger := testutils.CreateTestLogger()
	cache := testutils.CreateTestCache()

	args := map[string]interface{}{
		"prompt": "Test prompt",
	}

	// Execute will still fail due to missing Q CLI, but should not fail on enablement check
	result, err := tool.Execute(ctx, logger, cache, args)

	// Should get a CLI not found error, not an enablement error
	if err != nil {
		testutils.AssertTrue(t, !strings.Contains(err.Error(), "tool is not enabled"))
		// We expect CLI execution error instead
		testutils.AssertTrue(t, strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "failed"))
	}
	// Result may be nil due to CLI error, that's expected in this test environment
	_ = result
}

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

func TestQDeveloperTool_TimeoutConfiguration_ZeroTimeout(t *testing.T) {
	// Save original environment variable
	originalValue := os.Getenv("AGENT_TIMEOUT")
	defer func() {
		if originalValue == "" {
			_ = os.Unsetenv("AGENT_TIMEOUT")
		} else {
			_ = os.Setenv("AGENT_TIMEOUT", originalValue)
		}
	}()

	// Set zero timeout value
	_ = os.Setenv("AGENT_TIMEOUT", "0")

	tool := &qdeveloperagent.QDeveloperTool{}
	timeout := tool.GetTimeout()

	// Should fall back to default when zero value is provided
	testutils.AssertEqual(t, qdeveloperagent.DefaultTimeout, timeout)
}

func TestQDeveloperTool_TimeoutConfiguration_NegativeTimeout(t *testing.T) {
	// Save original environment variable
	originalValue := os.Getenv("AGENT_TIMEOUT")
	defer func() {
		if originalValue == "" {
			_ = os.Unsetenv("AGENT_TIMEOUT")
		} else {
			_ = os.Setenv("AGENT_TIMEOUT", originalValue)
		}
	}()

	// Set negative timeout value
	_ = os.Setenv("AGENT_TIMEOUT", "-60")

	tool := &qdeveloperagent.QDeveloperTool{}
	timeout := tool.GetTimeout()

	// Should fall back to default when negative value is provided
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

func TestQDeveloperTool_ResponseSizeLimit_InvalidEnvironmentVariable(t *testing.T) {
	// Save original environment variable
	originalValue := os.Getenv("AGENT_MAX_RESPONSE_SIZE")
	defer func() {
		if originalValue == "" {
			_ = os.Unsetenv("AGENT_MAX_RESPONSE_SIZE")
		} else {
			_ = os.Setenv("AGENT_MAX_RESPONSE_SIZE", originalValue)
		}
	}()

	// Set invalid environment variable value
	_ = os.Setenv("AGENT_MAX_RESPONSE_SIZE", "invalid")

	tool := &qdeveloperagent.QDeveloperTool{}

	// Should fall back to default when invalid value is provided
	maxSize := tool.GetMaxResponseSize()
	testutils.AssertEqual(t, qdeveloperagent.DefaultMaxResponseSize, maxSize)
}

func TestQDeveloperTool_ResponseSizeLimit_ZeroEnvironmentVariable(t *testing.T) {
	// Save original environment variable
	originalValue := os.Getenv("AGENT_MAX_RESPONSE_SIZE")
	defer func() {
		if originalValue == "" {
			_ = os.Unsetenv("AGENT_MAX_RESPONSE_SIZE")
		} else {
			_ = os.Setenv("AGENT_MAX_RESPONSE_SIZE", originalValue)
		}
	}()

	// Set zero environment variable value
	_ = os.Setenv("AGENT_MAX_RESPONSE_SIZE", "0")

	tool := &qdeveloperagent.QDeveloperTool{}

	// Should fall back to default when zero value is provided
	maxSize := tool.GetMaxResponseSize()
	testutils.AssertEqual(t, qdeveloperagent.DefaultMaxResponseSize, maxSize)
}

func TestQDeveloperTool_EnvironmentConstants(t *testing.T) {
	// Test that constants are exported and have expected values
	testutils.AssertEqual(t, "AGENT_MAX_RESPONSE_SIZE", qdeveloperagent.AgentMaxResponseSizeEnvVar)
	testutils.AssertEqual(t, "AGENT_TIMEOUT", qdeveloperagent.AgentTimeoutEnvVar)
	testutils.AssertEqual(t, 2*1024*1024, qdeveloperagent.DefaultMaxResponseSize)
	testutils.AssertEqual(t, 180, qdeveloperagent.DefaultTimeout)
}
