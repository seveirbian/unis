package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"
)

var imagesUsage = `Usage:  unisctl images [OPTIONS]

Options:
  -h, --help   help for images
`

var imagesCmd = &cobra.Command{
	Use:   "images", 
	Short: "List images in remote registry", 
	Long:  "List images in remote registry", 
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("\"unisctl images\" accepts no arguments.")
			return
		}
		fmt.Println("images")
	}, 
}

func init() {
	rootCmd.AddCommand(imagesCmd)
	imagesCmd.SetUsageTemplate(imagesUsage)
}