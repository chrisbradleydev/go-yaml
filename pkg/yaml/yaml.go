package yaml

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func DeleteNestedKey(data *yaml.Node, path string) *yaml.Node {
	keys := strings.Split(path, ".")
	return DeleteNestedKeyRecursive(data, keys)
}

func DeleteNestedKeyRecursive(node *yaml.Node, keys []string) *yaml.Node {
	if node.Kind == yaml.DocumentNode {
		node.Content[0] = DeleteNestedKeyRecursive(node.Content[0], keys)
		return node
	}

	if node.Kind != yaml.MappingNode || len(keys) == 0 {
		return node
	}

	for i := 0; i < len(node.Content); i += 2 {
		if node.Content[i].Value == keys[0] {
			if len(keys) == 1 {
				// Remove this key-value pair
				node.Content = append(node.Content[:i], node.Content[i+2:]...)
				return node
			}
			// Recurse into the next level
			node.Content[i+1] = DeleteNestedKeyRecursive(node.Content[i+1], keys[1:])
			return node
		}
	}

	// If the key wasn't found at this level, recurse into all values
	for i := 1; i < len(node.Content); i += 2 {
		node.Content[i] = DeleteNestedKeyRecursive(node.Content[i], keys)
	}

	return node
}

func FindAndDelete(filename string, paths []string, values []byte) error {
	var node yaml.Node
	err := yaml.Unmarshal(values, &node)
	if err != nil {
		return err
	}
	for _, path := range paths {
		updatedDefaultValues := DeleteNestedKey(&node, path)
		err = WriteYaml(updatedDefaultValues, filename)
		if err != nil {
			return fmt.Errorf("Error writing YAML file: %v\n", err)
		}
	}
	return nil
}

func WriteYaml(node *yaml.Node, filename string) error {
	// Create or truncate the file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	// Create a new encoder
	encoder := yaml.NewEncoder(file)
	defer encoder.Close()

	// Set the encoder to use 2-space indentation
	encoder.SetIndent(2)

	// Encode the YAML node to the file
	if err := encoder.Encode(node); err != nil {
		return fmt.Errorf("error encoding YAML: %w", err)
	}

	return nil
}
