package get

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Getter defines the interface for the get subcommand.
type Getter interface {
	GetCmd() *cobra.Command
}

// getImpl implements the Getter interface.
type getImpl struct {
	getCmd *cobra.Command
}

var osExit = os.Exit

// RunGet is the callback type that doorman plugin implementors must satisfy.
// It takes a secret key/path and returns the secret value.
type RunGet func(key string) (string, error)

// attachFlags adds flags to the get command.
func (s *getImpl) attachFlags() {
	s.getCmd.Flags().StringP("output", "o", "table", "Output format: table, json, or yaml")
}

// get executes the get callback and outputs the result.
func (s *getImpl) get(runGet RunGet, args []string) {
	if len(args) == 0 {
		s.getCmd.PrintErrln("Error: secret key/path argument required")
		osExit(2)
		return
	}

	key := args[0]
	value, err := runGet(key)
	if err != nil {
		s.getCmd.PrintErrln(err)
		osExit(1)
		return
	}

	format, _ := s.getCmd.Flags().GetString("output")
	switch format {
	case "json":
		data, _ := json.Marshal(map[string]string{"key": key, "value": value})
		fmt.Fprintln(s.getCmd.OutOrStdout(), string(data))
	case "yaml":
		fmt.Fprintf(s.getCmd.OutOrStdout(), "key: %s\nvalue: %s\n", key, value)
	default:
		// Plain text — just the value for piping
		fmt.Fprintln(s.getCmd.OutOrStdout(), value)
	}
}

// GetCmd returns the cobra command.
func (s *getImpl) GetCmd() *cobra.Command {
	return s.getCmd
}
