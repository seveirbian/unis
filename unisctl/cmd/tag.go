package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tagUsage = `Usage:  unisctl tag [OPTIONS] SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]

Options:
  -h, --help            help for tag
`

var tagCmd = &cobra.Command{
	Use:   "tag", 
	Short: "Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE", 
	Long:  "Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE", 
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("\"unisctl tag\" requires exactly 2 arguments.")
			return
		}
		fmt.Println("tag")
	}, 
}

func init() {
	rootCmd.AddCommand(tagCmd)
	tagCmd.SetUsageTemplate(tagUsage)
}