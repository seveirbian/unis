package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"
)

var imagesUsage = `Usage:  unisctl images [OPTIONS]

Options:
  -a, --all    Show all images (default private images)
  -h, --help   help for images
`

var allImagesFlag bool

var imagesCmd = &cobra.Command{
	Use:   "images", 
	Short: "List images in remote registry", 
	Long:  "List images in remote registry", 
	Args: cobra.NoArgs, 
	Run: func(cmd *cobra.Command, args []string) {
		if allImagesFlag {
			fmt.Println(allImagesFlag)
		}else {
			fmt.Println(allImagesFlag)
		}
	}, 
}

func init() {
	rootCmd.AddCommand(imagesCmd)
	imagesCmd.SetUsageTemplate(imagesUsage)
	imagesCmd.Flags().BoolVarP(&allImagesFlag, "all", "a", false, "Show all images")
}