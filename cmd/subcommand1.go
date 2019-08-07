package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
)

var subCommand1Cmd = &cobra.Command{
	Use:   "subcommand1",
	Short: "Short description for your subcommand1",
	Long:  `It prints CATS CATS CATS`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runSubCommand1(args, cmd.OutOrStdout())
	},
}

// Since this command is in the same package as the baseCmd,
// it can be attached on initialization
func init() {
	baseCmd.AddCommand(subCommand1Cmd)
}

func runSubCommand1(args []string, out io.Writer) error {
	fmt.Fprintf(out, "CATS CATS CATS\n")
	return nil
}
