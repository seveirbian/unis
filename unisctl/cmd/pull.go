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
	Args: cobra.ExactArgs(1), 
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pull")
		fmt.Println(args)
	}, 
}

func init() {
	rootCmd.AddCommand(pullCmd)
	pullCmd.SetUsageTemplate(pullUsage)
}