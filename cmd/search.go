package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/nao1215/http-status-code/http"

	"github.com/nao1215/http-status-code/internal/print"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [HTTP status code]",
	Short: "Print HTTP status code description and RFC (reference)",
	Long: `search subcommand print HTTP status code description and RFC (reference)
`,
	Example: `  http-status-code search 404`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := search(cmd, args); err != nil {
			print.Err(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

func search(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("you need to specify HTTP status code (e.g. hsc search 404)")
	}

	for _, v := range args {
		h := http.New()
		h = h.Search(v)

		if h.Code == "" {
			print.Warn("non-existent HTTP status code: " + v)
			continue
		}
		fmt.Printf("%s %s (ref.=%s)\n", h.Code, h.Description, h.RFC)
	}
	return nil
}
