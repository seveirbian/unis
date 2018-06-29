package cmd

import (
	"os"
	"runtime"
	"text/template"

	"github.com/spf13/cobra"
)

const ctlVersionTemplate = `Unisctl:
Version:     0.001.0.0
OS/Arch:     {{getOSArch}}
`

var versionUsage = `Usage:  unisctl version [OPTIONS]

Options:
  -h, --help            help for version
`

var ctlVersion = template.Must(template.New("ctlVersion").
	Funcs(template.FuncMap{"getOSArch": getOSArch}).
	Parse(ctlVersionTemplate))

var versionCmd = &cobra.Command{
	Use:   "version", 
	Short: "Show the unisctl and unisapiserver version", 
	Long:  "Show the unisctl and unisapiserver version", 
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

func getOSArch() string{
	goos := runtime.GOOS
	goarch := runtime.GOARCH

	return goos + "/" + goarch
}