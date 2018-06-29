package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"
)

var connectUsage = `Usage:  unisctl connect [OPTIONS] [SERVER]

Options:
  -h, --help   help for connect
`

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to unis-apiserver",
	Long: "Connect to unis-apiserver",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("\"unisctl connect\" requires exactly 1 argument.")
			return
		}
		fmt.Println(args)
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.SetUsageTemplate(connectUsage)
}