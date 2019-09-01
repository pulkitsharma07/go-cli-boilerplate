package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
)

// Variables initialized (injected) at build time.
// Refer: build_dist.sh

// Version specifies the git tag
var Version string
// Commit specifies the git commit
var Commit string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the CLI version",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runVersionCmd(args, cmd.OutOrStdout())
	},
}

// Since this command is in the same package as the baseCmd,
// it can be attached on initialization
func init() {
	baseCmd.AddCommand(versionCmd)
}

func runVersionCmd(args []string, out io.Writer) error {
	fmt.Fprintf(out, "Version:%s\tBuild:%s\n", Version, Commit)
	return nil
}
