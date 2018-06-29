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
var pushPublicFlag bool

var pushCmd = &cobra.Command{
	Use:   "push", 
	Short: "Push an image to registry", 
	Long:  "Push an image to registry", 
	Args: cobra.ExactArgs(1), 
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("push")
		fmt.Println(args)
		fmt.Println(cfgFile)
		if pushPublicFlag {
			fmt.Println(pushPublicFlag)
		}else {
			fmt.Println(pushPublicFlag)
		}
	}, 
}

func init() {
	rootCmd.AddCommand(pushCmd)
	pushCmd.SetUsageTemplate(pushUsage)
	pushCmd.Flags().StringVarP(&cfgFile, "configure-file", "f", "", "image's configure file (required)")
	pushCmd.MarkFlagRequired("configure-file")
	pushCmd.Flags().BoolVarP(&pushPublicFlag, "public", "p", false, "Push a public image")
}