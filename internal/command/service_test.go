package command

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	meta := PluginMeta{
		Name:        "doorman-vault",
		Version:     "1.0.0",
		Supports:    []string{"amadla.org/entity/infrastructure@^v1.0.0"},
		Description: "HashiCorp Vault secrets",
	}
	runGet := func(key string) (string, error) {
		return "secret-value", nil
	}

	New(cmd, meta, runGet)

	var names []string
	for _, sub := range cmd.Commands() {
		names = append(names, sub.Use)
	}
	assert.Contains(t, names, "get [key]")
	assert.Contains(t, names, "info")
}

func TestInfoSubcommand(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	meta := PluginMeta{
		Name:        "doorman-vault",
		Version:     "1.0.0",
		Supports:    []string{"amadla.org/entity/infrastructure@^v1.0.0"},
		Description: "HashiCorp Vault secrets",
	}
	runGet := func(key string) (string, error) {
		return "", nil
	}

	New(cmd, meta, runGet)

	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetArgs([]string{"info"})
	err := cmd.Execute()

	assert.NoError(t, err)
	assert.Contains(t, buf.String(), `"name":"doorman-vault"`)
	assert.Contains(t, buf.String(), `"version":"1.0.0"`)
}
