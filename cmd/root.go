package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/simondrake/hclconvert/cmd/json"
	"github.com/simondrake/hclconvert/cmd/yaml"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "hclconvert",
		Short: "converts to and from hcl format",
	}

	rootCmd.AddCommand(json.NewCommand())
	rootCmd.AddCommand(yaml.NewCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
