package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

type progressCmd struct {
	cmd *cobra.Command
}

func newProgressCmd() *progressCmd {
	root := &progressCmd{}
	root.cmd = &cobra.Command{
		Use:     "progress",
		Aliases: nil,
		Short:   "Get the progress of the cluster",
		Long:    "This command uploads a media file to the transcoding cluster.",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("upload requires at least a file")
			}
			return nil
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(args)
			return nil
		},
	}

	return root
}
