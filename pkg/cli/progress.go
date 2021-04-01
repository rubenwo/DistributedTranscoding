package cli

import (
	"github.com/cheggaaa/pb/v3"
	"github.com/gorilla/websocket"
	"github.com/rubenwo/DistributedTranscoding/pkg/distributor/pkg/distributor"
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
		Long:    "This command gets the progress of the transcoding cluster.",
		Args: func(cmd *cobra.Command, args []string) error {
			//if len(args) < 1 {
			//	return fmt.Errorf("upload requires at least a file")
			//}
			return nil
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			conn, _, err := websocket.DefaultDialer.Dial("ws://localhost/ws/jobs/status", nil)
			if err != nil {
				return err
			}

			var msg struct {
				CurrentJobState distributor.State `json:"state"`
				Error           string            `json:"error"`
			}
			count := 100
			bar := pb.StartNew(count)

			for err := conn.ReadJSON(&msg); ; {
				if err != nil {
					return err
				}
				bar.Increment()
			}

			bar.Finish()
			return nil
		},
	}

	return root
}
