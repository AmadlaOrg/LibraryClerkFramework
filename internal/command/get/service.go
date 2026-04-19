package get

import "github.com/spf13/cobra"

// New creates and returns a new Getter.
func New(runGet RunGet) Getter {
	g := &getImpl{
		getCmd: &cobra.Command{
			Use:   "get [key]",
			Short: "Retrieve a secret by key/path",
			Args:  cobra.MinimumNArgs(0),
		},
	}

	g.attachFlags()

	g.getCmd.Run = func(cmd *cobra.Command, args []string) {
		g.get(runGet, args)
	}

	return g
}
