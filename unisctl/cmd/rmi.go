package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rmiUsage = `Usage:  unisctl rmi [OPTIONS] IMAGE [IMAGE...]

Options:
  -h, --help            help for rmi
`

var rmiCmd = &cobra.Command{
	Use:   "rmi", 
	Short: "reomve one or more images in registry", 
	Long:  "remobe one or more images in registry", 
	Args: cobra.MinimumNArgs(1), 
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rmi")
		fmt.Println(args)
	}, 
}

func init() {
	rootCmd.AddCommand(rmiCmd)
	rmiCmd.SetUsageTemplate(rmiUsage)
}