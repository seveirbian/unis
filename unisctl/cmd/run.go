package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runUsage = `Usage:  unisctl run [OPTIONS] IMAGE

Options:
  -h, --help            help for run
  -p, --public          run as a public instance
`

var runPublicFlag bool

var runCmd = &cobra.Command{
	Use:   "run", 
	Short: "Run a container on a edge node", 
	Long:  "Run a container on a edge node", 
	Args: cobra.ExactArgs(1), 
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run")
		fmt.Println(args)
		if runPublicFlag {
			fmt.Println(runPublicFlag)
		}else {
			fmt.Println(runPublicFlag)
		}
	}, 
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.SetUsageTemplate(runUsage)
	runCmd.Flags().BoolVarP(&runPublicFlag, "public", "p", false, "Run as a public instance")
}