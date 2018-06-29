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
	Args: cobra.ExactArgs(2), 
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tag")
		fmt.Println(args)
	}, 
}

func init() {
	rootCmd.AddCommand(tagCmd)
	tagCmd.SetUsageTemplate(tagUsage)
}