package json

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "json",
		Short: "json related commands",
	}

	cmd.AddCommand(newFromHCLCommand())

	return cmd
}
