package yaml

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestDeleteNestedKeyRecursive(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		path     []string
		expected string
	}{
		{
			name: "Delete top-level key",
			input: `
key1: value1
key2: value2
`,
			path: []string{"key1"},
			expected: `key2: value2
`,
		},
		{
			name: "Delete nested key",
			input: `
parent:
  child1:
    grandchild1: value1
    grandchild2: value2
  child2: value3
`,
			path: []string{"parent", "child1", "grandchild1"},
			expected: `parent:
    child1:
        grandchild2: value2
    child2: value3
`,
		},
		{
			name: "Delete non-existent key",
			input: `
key1: value1
key2: value2
`,
			path: []string{"key3"},
			expected: `key1: value1
key2: value2
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse input YAML
			var node yaml.Node
			err := yaml.Unmarshal([]byte(tt.input), &node)
			if err != nil {
				t.Fatalf("Failed to parse input YAML: %v", err)
			}

			// Call the function
			result := DeleteNestedKeyRecursive(&node, tt.path)

			// Convert result back to byte array
			output, err := yaml.Marshal(result)
			if err != nil {
				t.Fatalf("Failed to marshal result: %v", err)
			}

			// Compare with expected output
			if string(output) != tt.expected {
				t.Errorf("\nExpected:\n%s\nGot:\n%s", tt.expected, string(output))
			}
		})
	}
}
