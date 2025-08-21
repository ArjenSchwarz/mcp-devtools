package tools_test

import (
	"testing"

	"github.com/sammcj/mcp-devtools/internal/tools/qdeveloperagent"
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
