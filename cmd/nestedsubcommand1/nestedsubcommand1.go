package nestedsubcommand1

import (
	"github.com/spf13/cobra"
)

var nestedSubCommand1 = &cobra.Command{
	Use:   "nestedsubcommand1",
	Short: "Short description for your nestedsubcommand1",
	Long:  `Sample command to show a command which does not have a implementation of its own`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Help()
		return nil
	},
}

// SetParent can be called to join this command (along with its sub-tree) to some other command (baseCmd in this case)
func SetParent(parentCommand *cobra.Command) {
	parentCommand.AddCommand(nestedSubCommand1)
}
