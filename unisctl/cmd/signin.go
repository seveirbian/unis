package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var signinUsage = `Usage:  unisctl signin [OPTIONS]

Options:
  -h, --help            help for signin
  -p, --password        password
  -u, --username        username
`

var siUsername, siPassword string

var signinCmd = &cobra.Command{
	Use:   "signin", 
	Short: "Sign in", 
	Long:  "Sign in", 
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("\"unisctl signin\" accepts no arguments.")
			return
		}
		fmt.Println("signin")
	}, 
}

func init() {
	rootCmd.AddCommand(signinCmd)
	signinCmd.SetUsageTemplate(signinUsage)
	signinCmd.Flags().StringVarP(&siUsername, "username", "u", "", "username used to sign in")
	signinCmd.Flags().StringVarP(&siPassword, "password", "p", "", "password used to sign in")
	signinCmd.MarkFlagRequired("username")
	signinCmd.MarkFlagRequired("password")
}