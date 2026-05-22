package alias

import (
	"fmt"
	"strings"
)

// ParseInput parses a string in the format name="command" or name=command.
func ParseInput(input string) (string, string, error) {
	parts := strings.SplitN(input, "=", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid format. Use: name=\"command\"")
	}

	name := strings.TrimSpace(parts[0])
	command := strings.TrimSpace(parts[1])

	// Remove surrounding quotes
	if len(command) >= 2 && command[0] == '"' && command[len(command)-1] == '"' {
		command = command[1 : len(command)-1]
	}

	if name == "" {
		return "", "", fmt.Errorf("name cannot be empty")
	}
	if command == "" {
		return "", "", fmt.Errorf("command cannot be empty")
	}

	return name, command, nil
}
