package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/version"
)

func genVersionCmd(name string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show current version info",
		Run: func(cmd *cobra.Command, args []string) {
			beat, err := beat.New(name, "")
			if err != nil {
				os.Exit(1)
			}

			fmt.Printf("%s version %s (%s), libbeat %s\n",
				beat.Info.Beat, beat.Info.Version, runtime.GOARCH, version.GetDefaultVersion())
		},
	}
}
