package json

import (
	"fmt"
	"io"
	"os"

	"github.com/hashicorp/hcl/hcl/printer"
	"github.com/hashicorp/hcl/json/parser"
	"github.com/spf13/cobra"
)

func newToHCLCommand() *cobra.Command {
	var filePath string

	stackDiffCmd := &cobra.Command{
		Use: "to_hcl",
		Run: func(cmd *cobra.Command, args []string) {
			f, err := os.Open(filePath)
			if err != nil {
				fmt.Fprintln(os.Stderr, fmt.Errorf("unable to open json file: %w", err))
				os.Exit(1)
			}

			b, err := io.ReadAll(f)
			if err != nil {
				fmt.Fprintln(os.Stderr, fmt.Errorf("unable to read json file: %w", err))
				os.Exit(1)
			}

			ast, err := parser.Parse(b)
			if err != nil {
				fmt.Fprintln(os.Stderr, fmt.Errorf("unable to parse json file: %w", err))
				os.Exit(1)
			}

			if err := printer.Fprint(os.Stdout, ast); err != nil {
				fmt.Fprintln(os.Stderr, fmt.Errorf("unable to print hcl: %w", err))
				os.Exit(1)
			}
		},
	}

	stackDiffCmd.Flags().StringVarP(&filePath, "file_path", "f", "", "the path of the hcl file")

	stackDiffCmd.MarkFlagRequired("file_path")

	return stackDiffCmd
}
