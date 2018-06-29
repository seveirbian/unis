package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pushUsage = `Usage:  unisctl push [OPTIONS] NAME[:TAG]

Options:
  -f, --configure-file  Add configure file with image
  -h, --help            help for push
`

var cfgFile string

var pushCmd = &cobra.Command{
	Use:   "push", 
	Short: "Push an image to registry", 
	Long:  "Push an image to registry", 
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("\"unisctl push\" requires exactly 1 argument.")
			return
		}
		fmt.Println("push")
		fmt.Println(cfgFile)
	}, 
}

func init() {
	rootCmd.AddCommand(pushCmd)
	pushCmd.SetUsageTemplate(pushUsage)
	pushCmd.Flags().StringVarP(&cfgFile, "configure-file", "f", "", "image's configure file (required)")
	pushCmd.MarkFlagRequired("configure-file")
}