package main

import (
	"fmt"
	"os"
	"strings"

	"win-alias/internal/alias"
)

func main() {
	if len(os.Args) < 2 {
		if err := alias.List(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	arg := os.Args[1]

	switch arg {
	case "--setup":
		if err := alias.Setup(); err != nil {
			fmt.Printf("Error during setup: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("AutoRun configured. Aliases will persist across CMD sessions.")
	case "--disable":
		if err := alias.Disable(); err != nil {
			fmt.Printf("Error during disable: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("AutoRun persistence disabled.")
	case "--load":
		alias.Load()
	case "unalias":
		if len(os.Args) < 3 {
			fmt.Println("Usage: unalias <name>")
			os.Exit(1)
		}
		if err := alias.Delete(os.Args[2]); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Alias '%s' removed.\n", os.Args[2])
	default:
		// Check if it's an assignment: name="command"
		if strings.Contains(arg, "=") {
			handleAssignment(strings.Join(os.Args[1:], " "))
		} else {
			printUsage()
		}
	}
}

func handleAssignment(input string) {
	name, command, err := alias.ParseInput(input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if err := alias.Save(name, command); err != nil {
		fmt.Printf("Error saving alias: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Alias saved: %s=%s\n", name, command)
	// Apply immediately to current session
	alias.Apply(name, command)
}

func printUsage() {
	fmt.Println("win-alias: Linux-like aliases for Windows CMD")
	fmt.Println("\nUsage:")
	fmt.Println("  alias name=\"command\"   Add/Update alias")
	fmt.Println("  alias                  List all aliases")
	fmt.Println("  unalias name           Remove an alias")
	fmt.Println("  alias --setup          Configure persistence (AutoRun)")
	fmt.Println("  alias --disable        Disable persistence (AutoRun)")
	fmt.Println("  alias --load           Apply aliases to current session")
}
