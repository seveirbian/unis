package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"
)

var psUsage = `Usage:  unisctl ps [OPTIONS]

Options:
  -a, --all    Show all containers (default shows just running)
  -h, --help   help for ps
`

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "List containers",
	Long:  "List containers",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("\"unisctl ps\" accepts no arguments.")
			return
		}
		fmt.Println("ps")
	},
}

func init() {
	rootCmd.AddCommand(psCmd)
	psCmd.SetUsageTemplate(psUsage)
}