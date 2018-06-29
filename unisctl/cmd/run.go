package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runUsage = `Usage:  unisctl run [OPTIONS] IMAGE

Options:
  -h, --help            help for run
`

var runCmd = &cobra.Command{
	Use:   "run", 
	Short: "Run a container on a edge node", 
	Long:  "Run a container on a edge node", 
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("\"unisctl run\" requires exactly 1 argument.")
			return
		}
		fmt.Println("run")
	}, 
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.SetUsageTemplate(runUsage)
}