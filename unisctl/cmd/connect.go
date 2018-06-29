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
	Args: cobra.ExactArgs(1), 
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("connect")
		fmt.Println(args)
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.SetUsageTemplate(connectUsage)
}