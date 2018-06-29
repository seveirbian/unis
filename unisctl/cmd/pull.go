package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pullUsage = `Usage:  unisctl pull [OPTIONS] NAME[:TAG]

Options:
  -h, --help   help for pull
`

var pullCmd = &cobra.Command{
	Use:   "pull", 
	Short: "Pull an image from registry", 
	Long:  "Pull an image from registry", 
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("\"unisctl pull requires exactly 1 argument.")
			return
		}
		fmt.Println("pull")
	}, 
}

func init() {
	rootCmd.AddCommand(pullCmd)
	pullCmd.SetUsageTemplate(pullUsage)
}