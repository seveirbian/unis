package cmd

import (
	"github.com/spf13/cobra"
)

var rmiUsage = `Usage:  unisctl rmi [OPTIONS] IMAGE [IMAGE...]

Options:
  -h, --help            help for rmi
  -p, --public          Remove public images (default private images)
`

var rmiCmd = &cobra.Command{
	Use:   "rmi",
	Short: "reomve one or more images in registry",
	Long:  "remobe one or more images in registry",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// http.Post()
	},
}

func init() {
	rootCmd.AddCommand(rmiCmd)
	rmiCmd.SetUsageTemplate(rmiUsage)
}
