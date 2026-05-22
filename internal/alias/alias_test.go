package alias

import (
	"testing"
	"golang.org/x/sys/windows/registry"
)

func TestParseAliasInput(t *testing.T) {
	tests := []struct {
		input       string
		wantName    string
		wantCommand string
		wantErr     bool
	}{
		{"g=python main.py", "g", "python main.py", false},
		{"ll=\"dir /p /w\"", "ll", "dir /p /w", false},
		{"g python main.py", "", "", true},
		{"=command", "", "", true},
		{"name=", "", "", true},
		{"  gs  =  \"git status\"  ", "gs", "git status", false},
		{"empty=\"\"", "", "", true},
	}

	for _, tt := range tests {
		name, cmd, err := ParseInput(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("ParseInput(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			continue
		}
		if name != tt.wantName {
			t.Errorf("ParseInput(%q) name = %q, want %q", tt.input, name, tt.wantName)
		}
		if cmd != tt.wantCommand {
			t.Errorf("ParseInput(%q) command = %q, want %q", tt.input, cmd, tt.wantCommand)
		}
	}
}

func TestRegistryOperations(t *testing.T) {
	// Use a test-specific registry path
	originalRegPath := RegPath
	RegPath = `Software\win-alias\test_aliases`
	defer func() {
		// Clean up and restore
		k, err := registry.OpenKey(RegBase, `Software\win-alias`, registry.ALL_ACCESS)
		if err == nil {
			registry.DeleteKey(k, "test_aliases")
			k.Close()
		}
		RegPath = originalRegPath
	}()

	aliasName := "test-g"
	aliasCmd := "git status"

	// Test Save
	err := Save(aliasName, aliasCmd)
	if err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	// Verify Save via Registry direct read
	k, err := registry.OpenKey(RegBase, RegPath, registry.READ)
	if err != nil {
		t.Fatalf("failed to open test registry key: %v", err)
	}
	val, _, err := k.GetStringValue(aliasName)
	k.Close()
	if err != nil || val != aliasCmd {
		t.Errorf("registry value mismatch: got %q, err %v", val, err)
	}

	// Test Delete
	err = Delete(aliasName)
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	// Verify Delete
	k, err = registry.OpenKey(RegBase, RegPath, registry.READ)
	if err != nil {
		t.Fatalf("failed to open test registry key after delete: %v", err)
	}
	_, _, err = k.GetStringValue(aliasName)
	k.Close()
	if err == nil {
		t.Errorf("alias still exists after Delete")
	}
}
