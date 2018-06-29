package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"
)

var nodesUsage = `Usage:  unisctl nodes [OPTIONS]

Options:
  -h, --help   help for nodes
`

var nodesCmd = &cobra.Command{
	Use:   "nodes", 
	Short: "Display the status of all edge nodes", 
	Long:  "Display the status of all edge nodes", 
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("\"unisctl nodes\" accepts no arguments.")
			return
		}
		fmt.Println("nodes")
	}, 
}

func init() {
	rootCmd.AddCommand(nodesCmd)
	nodesCmd.SetUsageTemplate(nodesUsage)
}