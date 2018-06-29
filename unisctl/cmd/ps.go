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

var allInstancesFlag bool

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "List containers",
	Long:  "List containers",
	Args: cobra.NoArgs, 
	Run: func(cmd *cobra.Command, args []string) {
		if allInstancesFlag {
			fmt.Println(allInstancesFlag)
		}else {
			fmt.Println(allInstancesFlag)
		}
	},
}

func init() {
	rootCmd.AddCommand(psCmd)
	psCmd.SetUsageTemplate(psUsage)
	psCmd.Flags().BoolVarP(&allInstancesFlag, "all", "a", false, "Show all instances")
}