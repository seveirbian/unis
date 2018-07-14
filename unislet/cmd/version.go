package cmd

import (
	"os"
	"runtime"
	"text/template"

	"github.com/spf13/cobra"
)

const ctlVersionTemplate = `Unislet:
Version:     0.001.0.0
OS/Arch:     {{getOSArch}}
`

var versionUsage = `Usage:  unislet version [OPTIONS]

Options:
  -h, --help            help for version
`

var ctlVersion = template.Must(template.New("ctlVersion").
	Funcs(template.FuncMap{"getOSArch": getOSArch}).
	Parse(ctlVersionTemplate))

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the unislet version",
	Long:  "Show the unislet version",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := ctlVersion.Execute(os.Stdout, nil); err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.SetUsageTemplate(versionUsage)
}

func getOSArch() string {
	goos := runtime.GOOS
	goarch := runtime.GOARCH

	return goos + "/" + goarch
}
