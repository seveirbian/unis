package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var statsUsage = `Usage:  unisctl stats [OPTIONS]

Options:
  -h, --help            help for stats
`

var statsCmd = &cobra.Command{
	Use:   "stats", 
	Short: "Display the status of all components of unis", 
	Long:  "Display the status of all components of unis", 
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("\"unisctl stats\" accepts no arguments.")
			return
		}
		fmt.Println("stats")
	}, 
}

func init() {
	rootCmd.AddCommand(statsCmd)
	statsCmd.SetUsageTemplate(statsUsage)
}