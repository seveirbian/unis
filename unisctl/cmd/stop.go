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
	Args: cobra.MinimumNArgs(1), 
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stop")
		fmt.Println(args)
	}, 
}

func init() {
	rootCmd.AddCommand(stopCmd)
	stopCmd.SetUsageTemplate(stopUsage)
}