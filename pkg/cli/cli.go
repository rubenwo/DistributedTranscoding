package cli

import (
	"fmt"
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/spf13/cobra"
	"os"
)

func Execute() {
	log.SetHandler(cli.Default)
	newRootCmd().Execute()
}

type rootCmd struct {
	cmd     *cobra.Command
	backend string
	verbose bool
}

func (cmd *rootCmd) Execute() {
	if err := cmd.cmd.Execute(); err != nil {
		log.WithError(err).Error("command failed")
		os.Exit(1)
	}
}

func newRootCmd() *rootCmd {
	root := &rootCmd{}

	root.cmd = &cobra.Command{
		Use:           "dist-ffmpeg",
		Short:         "Cli tool for using distributed transcoding cluster",
		Long:          "Use this cli tool to send a media file to the distributed transcoding cluster to transcode it",
		SilenceErrors: true,
		SilenceUsage:  true,
		Version:       "0.0.1",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if root.backend != "" {
				fmt.Println(root.backend)
			}

			if root.verbose {
				log.SetLevel(log.DebugLevel)
			}
		},
	}

	root.cmd.PersistentFlags().StringVar(&root.backend, "distributor", "", "Distributor url")
	root.cmd.PersistentFlags().BoolVarP(&root.verbose, "verbose", "v", false, "Enable verbose output")

	root.cmd.AddCommand(newUploadCmd().cmd)
	root.cmd.AddCommand(newProgressCmd().cmd)

	return root
}
