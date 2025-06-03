package yaml

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "yaml",
		Short: "yaml related commands",
	}

	cmd.AddCommand(newFromHCLCommand())

	return cmd
}
