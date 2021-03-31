package transcoder

import (
	"context"
	"fmt"
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/spf13/cobra"
	"os"
)

func ExecuteCli() {
	log.SetHandler(cli.Default)
	newRootCmd().Execute()
}

type rootCmd struct {
	cmd               *cobra.Command
	backend           string
	verbose           bool
	maxConcurrentJobs int64
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
		Use:           "dist-ffmpeg-transcoder",
		Short:         "Cli tool for using distributed transcoding cluster",
		Long:          "Use this cli tool to send a media file to the distributed transcoding cluster to transcode it",
		SilenceErrors: true,
		SilenceUsage:  true,
		Version:       "0.0.1",
		RunE: func(cmd *cobra.Command, args []string) error {
			if root.backend == "" {
				return fmt.Errorf("distributor backend is not set")
			}

			client := NewClient(root.backend, root.maxConcurrentJobs)
			ctx := context.TODO()
			log.Infof("Joining cluster on: %s with support for %d concurrent job(s)", root.backend, root.maxConcurrentJobs)
			if err := client.JoinCluster(ctx); err != nil {
				return err
			}

			if root.verbose {
				log.SetLevel(log.DebugLevel)
			}
			return nil
		},
	}

	root.cmd.PersistentFlags().StringVar(&root.backend, "distributor", "localhost:8080", "Distributor url")
	root.cmd.PersistentFlags().Int64Var(&root.maxConcurrentJobs, "jobs", 1, "Set's the amount of concurrent jobs")
	root.cmd.PersistentFlags().BoolVarP(&root.verbose, "verbose", "v", false, "Enable verbose output")

	return root
}
