package json

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/hashicorp/hcl"
	"github.com/spf13/cobra"
)

func newFromHCLCommand() *cobra.Command {
	var filePath string

	stackDiffCmd := &cobra.Command{
		Use: "from_hcl",
		Run: func(cmd *cobra.Command, args []string) {
			f, err := os.Open(filePath)
			if err != nil {
				fmt.Fprintln(os.Stderr, fmt.Errorf("unable to open hcl file: %w", err))
				os.Exit(1)
			}

			b, err := io.ReadAll(f)
			if err != nil {
				fmt.Fprintln(os.Stderr, fmt.Errorf("unable to read hcl file: %w", err))
				os.Exit(1)
			}

			var a any
			if err := hcl.Unmarshal(b, &a); err != nil {
				fmt.Fprintln(os.Stderr, fmt.Errorf("unable to parse hcl file: %w", err))
				os.Exit(1)
			}

			json, err := json.MarshalIndent(a, "", "	")
			if err != nil {
				fmt.Fprintln(os.Stderr, fmt.Errorf("unable to convert hcl file to json: %w", err))
				os.Exit(1)
			}

			fmt.Fprintln(os.Stdout, string(json))
		},
	}

	stackDiffCmd.Flags().StringVar(&filePath, "file_path", "", "the path of the hcl file")

	stackDiffCmd.MarkFlagRequired("file_path")

	return stackDiffCmd
}
