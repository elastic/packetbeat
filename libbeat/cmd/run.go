package cmd

import (
	"flag"
	"os"

	"github.com/spf13/cobra"

	"github.com/elastic/beats/libbeat/beat"
)

func genRunCmd(name string, beatCreator beat.Creator) *cobra.Command {
	runCmd := cobra.Command{
		Use:   "run",
		Short: "Run " + name,
		Run: func(cmd *cobra.Command, args []string) {
			if err := beat.Run(name, "", beatCreator); err != nil {
				os.Exit(1)
			}
		},
	}

	// Run subcommand flags, only available to *beat run
	runCmd.Flags().AddGoFlag(flag.CommandLine.Lookup("N"))
	runCmd.Flags().AddGoFlag(flag.CommandLine.Lookup("e"))
	runCmd.Flags().AddGoFlag(flag.CommandLine.Lookup("v"))
	runCmd.Flags().AddGoFlag(flag.CommandLine.Lookup("httpprof"))
	runCmd.Flags().AddGoFlag(flag.CommandLine.Lookup("cpuprofile"))
	runCmd.Flags().AddGoFlag(flag.CommandLine.Lookup("memprofile"))

	// TODO deprecate in favor of subcommands (7.0):
	runCmd.Flags().AddGoFlag(flag.CommandLine.Lookup("configtest"))
	runCmd.Flags().AddGoFlag(flag.CommandLine.Lookup("setup"))
	runCmd.Flags().AddGoFlag(flag.CommandLine.Lookup("version"))

	return &runCmd
}
