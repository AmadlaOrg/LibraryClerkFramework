package doorman

import (
	"github.com/AmadlaOrg/LibraryDoormanFramework/internal/command"
	"github.com/AmadlaOrg/LibraryDoormanFramework/internal/command/get"
	"github.com/AmadlaOrg/LibraryFramework/cli"
	"github.com/spf13/cobra"
)

// RunGet is the callback type that doorman plugin implementors must satisfy.
// It takes a secret key/path and returns the secret value.
type RunGet = get.RunGet

// New sets up the doorman plugin CLI application with UNIX plugin protocol support.
//
// The generated CLI includes:
//   - info subcommand: outputs plugin metadata as JSON
//   - get subcommand: retrieves a secret by key/path
//   - version subcommand: prints the version
//
// Parameters:
//   - name: plugin binary name (e.g., "doorman-vault")
//   - title: display name (e.g., "Doorman Vault")
//   - version: semantic version (e.g., "1.0.0")
//   - supports: list of entity type URIs this plugin handles
//   - description: short description of what the plugin does
//   - runGet: callback that retrieves a secret
func New(
	name, title, version string,
	supports []string,
	description string,
	runGet RunGet) {

	meta := command.PluginMeta{
		Name:        name,
		Version:     version,
		Supports:    supports,
		Description: description,
	}

	cli.New(name, title, version, func(c *cobra.Command) {
		command.New(c, meta, runGet)
	})
}
