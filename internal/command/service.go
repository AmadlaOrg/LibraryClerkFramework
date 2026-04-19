package command

import (
	"encoding/json"
	"fmt"

	"github.com/AmadlaOrg/LibraryDoormanFramework/internal/command/get"
	"github.com/spf13/cobra"
)

// PluginMeta holds metadata for the info subcommand.
type PluginMeta struct {
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Supports    []string `json:"supports"`
	Description string   `json:"description"`
}

// New wires up subcommands onto the root cobra command.
func New(cmd *cobra.Command, meta PluginMeta, runGet get.RunGet) {
	getService := get.New(runGet)
	cmd.AddCommand(getService.GetCmd())

	// info subcommand — outputs plugin metadata as JSON per UNIX plugin protocol
	infoCmd := &cobra.Command{
		Use:   "info",
		Short: "Display plugin metadata",
		Run: func(c *cobra.Command, args []string) {
			data, _ := json.Marshal(meta)
			fmt.Fprintln(c.OutOrStdout(), string(data))
		},
	}
	cmd.AddCommand(infoCmd)
}
