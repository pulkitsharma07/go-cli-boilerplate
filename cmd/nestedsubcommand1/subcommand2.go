package nestedsubcommand1

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
)

var subCommand2Cmd = &cobra.Command{
	Use:   "subcommand2",
	Short: "Short description for your subcommand2",
	Long:  `.It prints whatever you want 3 times.`,
        Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runSubCommand2(args, cmd.OutOrStdout())
	},
}

// Since this command is in the same package as the nestedSubCommand1,
// it can be attached on initialization
func init() {
	nestedSubCommand1.AddCommand(subCommand2Cmd)
}

func runSubCommand2(args []string, out io.Writer) error {
	fmt.Fprintf(out, "%s %s %s\n", args[0], args[0], args[0])
	return nil
}
