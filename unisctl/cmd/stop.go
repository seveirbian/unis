package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var stopUsage = `Usage:  unisctl stop [OPTIONS] INSTANCE [INSTANCE...]

Options:
  -h, --help            help for stop
`

var stopCmd = &cobra.Command{
	Use:   "stop", 
	Short: "Stop one or more running instances", 
	Long:  "Stop one or more running instances", 
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("\"unisctl stop\" requires at least 1 argument.")
			return
		}
		fmt.Println("stop")
	}, 
}

func init() {
	rootCmd.AddCommand(stopCmd)
	stopCmd.SetUsageTemplate(stopUsage)
}