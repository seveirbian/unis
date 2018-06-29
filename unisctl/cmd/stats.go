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
	Args: cobra.NoArgs, 
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stats")
	}, 
}

func init() {
	rootCmd.AddCommand(statsCmd)
	statsCmd.SetUsageTemplate(statsUsage)
}