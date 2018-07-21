package cmd

import (
	"github.com/spf13/cobra"
)

var statusUsage = `Usage:  unisctl status [OPTIONS]

Options:
  -h, --help   help for status
`

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "List containers",
	Long:  "List containers",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.SetUsageTemplate(statusUsage)
}
