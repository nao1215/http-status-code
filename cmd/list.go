package cmd

import (
	"os"

	"github.com/nao1215/hsc/http"
	"github.com/nao1215/hsc/internal/print"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List up HTTP status code, description, and RFC (reference)",
	Long: `list subcommand List up HTTP status code, description, and RFC (reference)
`,
	Example: `  hsc list`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := list(cmd, args); err != nil {
			print.Err(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list(cmd *cobra.Command, args []string) error {
	h := http.New()

	tableData := [][]string{}
	for _, v := range h.All() {
		tableData = append(tableData, []string{v.Code, v.Description, v.RFC})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"HTTP status code", "Description", "RFC"})
	table.SetAutoWrapText(false)
	for _, v := range tableData {
		table.Append(v)
	}
	table.Render()

	return nil
}
