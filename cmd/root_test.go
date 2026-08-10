package cmd

import "testing"

func TestNewRootCmd_HasSubcommands(t *testing.T) {
	root := NewRootCmd()

	for _, name := range []string{"license", "ignore"} {
		found, _, err := root.Find([]string{name})
		if err != nil {
			t.Errorf("Find(%q) error: %v", name, err)
			continue
		}
		if found.Name() != name {
			t.Errorf("Find(%q) resolved to %q", name, found.Name())
		}
	}
}
