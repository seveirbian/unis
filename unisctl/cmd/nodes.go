package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"
)

var nodesUsage = `Usage:  unisctl nodes [OPTIONS]

Options:
  -a, --all    Show all nodes (default private nodes)
  -h, --help   help for nodes
`

var allNodesFlag bool

var nodesCmd = &cobra.Command{
	Use:   "nodes", 
	Short: "Display the status of all edge nodes", 
	Long:  "Display the status of all edge nodes", 
	Args: cobra.NoArgs, 
	Run: func(cmd *cobra.Command, args []string) {
		if allNodesFlag {
			fmt.Println(allNodesFlag)
		}else {
			fmt.Println(allNodesFlag)
		}
	}, 
}

func init() {
	rootCmd.AddCommand(nodesCmd)
	nodesCmd.SetUsageTemplate(nodesUsage)
	nodesCmd.Flags().BoolVarP(&allNodesFlag, "all", "a", false, "Show all nodes")
}