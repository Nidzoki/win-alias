package alias

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/sys/windows/registry"
)

var (
	RegBase    = registry.CURRENT_USER
	RegPath    = `Software\win-alias\aliases`
	AutoRunKey = `Software\Microsoft\Command Processor`
)

// Save stores an alias in the Windows Registry.
func Save(name, command string) error {
	k, _, err := registry.CreateKey(RegBase, RegPath, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer k.Close()

	return k.SetStringValue(name, command)
}

// Delete removes an alias from the Windows Registry.
func Delete(name string) error {
	k, err := registry.OpenKey(RegBase, RegPath, registry.ALL_ACCESS)
	if err != nil {
		return fmt.Errorf("error opening registry: %w", err)
	}
	defer k.Close()

	if err := k.DeleteValue(name); err != nil {
		return fmt.Errorf("error deleting alias '%s': %w", name, err)
	}
	return nil
}

// List prints all aliases stored in the Registry.
func List() error {
	k, err := registry.OpenKey(RegBase, RegPath, registry.READ)
	if err != nil {
		if err == registry.ErrNotExist {
			fmt.Println("No aliases defined.")
			return nil
		}
		return fmt.Errorf("error reading registry: %w", err)
	}
	defer k.Close()

	names, err := k.ReadValueNames(0)
	if err != nil {
		return fmt.Errorf("error reading values: %w", err)
	}

	if len(names) == 0 {
		fmt.Println("No aliases defined.")
		return nil
	}

	fmt.Println("Active Aliases:")
	for _, name := range names {
		val, _, _ := k.GetStringValue(name)
		fmt.Printf("  %s=%s\n", name, val)
	}
	return nil
}

// Load applies all stored aliases to the current session via doskey.
func Load() {
	k, err := registry.OpenKey(RegBase, RegPath, registry.READ)
	if err != nil {
		if err == registry.ErrNotExist {
			return
		}
		fmt.Fprintf(os.Stderr, "Error loading aliases: %v\n", err)
		return
	}
	defer k.Close()

	names, _ := k.ReadValueNames(0)
	for _, name := range names {
		val, _, _ := k.GetStringValue(name)
		Apply(name, val)
	}
}

// Apply registers a single doskey macro.
func Apply(name, command string) {
	cmd := exec.Command("doskey", fmt.Sprintf("%s=%s", name, command))
	cmd.Run()
}

// Setup configures CMD AutoRun to load aliases on every startup.
func Setup() error {
	k, _, err := registry.CreateKey(RegBase, AutoRunKey, registry.ALL_ACCESS)
	if err != nil {
		return fmt.Errorf("error opening AutoRun key: %w", err)
	}
	defer k.Close()

	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("error getting executable path: %w", err)
	}

	command := fmt.Sprintf("\"%s\" --load", exePath)
	
	existing, _, _ := k.GetStringValue("AutoRun")
	if strings.Contains(existing, "--load") {
		fmt.Println("AutoRun already configured.")
		return nil
	}

	newAutoRun := command
	if existing != "" {
		newAutoRun = existing + " & " + command
	}

	return k.SetStringValue("AutoRun", newAutoRun)
}

// Disable removes the alias --load hook from CMD AutoRun.
func Disable() error {
	k, err := registry.OpenKey(RegBase, AutoRunKey, registry.ALL_ACCESS)
	if err != nil {
		if err == registry.ErrNotExist {
			return nil
		}
		return fmt.Errorf("error opening AutoRun key: %w", err)
	}
	defer k.Close()

	existing, _, _ := k.GetStringValue("AutoRun")
	if !strings.Contains(existing, "--load") {
		fmt.Println("AutoRun hook not found.")
		return nil
	}

	// Remove our command. Handle cases with and without ' & '
	exePath, _ := os.Executable()
	cmd := fmt.Sprintf("\"%s\" --load", exePath)
	
	newAutoRun := strings.ReplaceAll(existing, " & "+cmd, "")
	newAutoRun = strings.ReplaceAll(newAutoRun, cmd+" & ", "")
	newAutoRun = strings.ReplaceAll(newAutoRun, cmd, "")
	newAutoRun = strings.TrimSpace(newAutoRun)

	if newAutoRun == "" {
		return k.DeleteValue("AutoRun")
	}
	return k.SetStringValue("AutoRun", newAutoRun)
}
