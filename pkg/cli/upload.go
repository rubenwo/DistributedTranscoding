package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/apex/log"
	"github.com/spf13/cobra"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type uploadCmd struct {
	cmd          *cobra.Command
	vCodec       string
	vPreset      string
	aCodec       string
	aPreset      string
	outputFormat string
}

func newUploadCmd() *uploadCmd {
	root := &uploadCmd{}
	root.cmd = &cobra.Command{
		Use:     "upload <filename>",
		Aliases: nil,
		Short:   "Upload a media file",
		Long:    "This command uploads a media file to the transcoding cluster.",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("upload requires at least a file")
			}
			return nil
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]
			log.Infof("opening file: %s", path)
			file, err := os.Open(path)

			if err != nil {
				return err
			}
			defer file.Close()

			log.Debugf("creating writer...")
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))

			if err != nil {
				return err
			}

			log.Debugf("writing file contents to writer buffer")
			io.Copy(part, file)
			writer.Close()
			log.Debugf("creating POST request")
			request, err := http.NewRequest("POST", "http://localhost/job", body)

			if err != nil {
				return err
			}

			log.Debugf("setting request headers")
			request.Header.Add("Content-Type", writer.FormDataContentType())
			client := &http.Client{}

			log.Infof("sending request to the distributor")
			response, err := client.Do(request)

			if err != nil {
				return err
			}
			defer response.Body.Close()

			log.Infof("reading response from server")
			var msg struct {
				Id string `json:"id"`
			}
			err = json.NewDecoder(response.Body).Decode(&msg)
			if err != nil {
				return err
			}

			log.Infof("cluster is now transcoding. job id: %s", msg.Id)
			return nil
		},
	}

	root.cmd.PersistentFlags().StringVar(&root.vCodec, "vcodec", "",
		"Video codec for the output file: copy, libx264, libx265, vpx, vp8, vp9")
	root.cmd.PersistentFlags().StringVar(&root.vPreset, "videopreset", "",
		"Video preset: optional")
	root.cmd.PersistentFlags().StringVar(&root.aCodec, "acodec", "",
		"Audio codec for output file: copy, aac, ac3, opus")
	root.cmd.PersistentFlags().StringVar(&root.aPreset, "audiopreset", "",
		"Audio preset: optional")
	root.cmd.PersistentFlags().StringVar(&root.outputFormat, "format", "",
		"Format of the output type: mkv, mp4, webm")

	return root
}
