package transcoder

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	v1 "github.com/rubenwo/DistributedTranscoding/pkg/api/v1"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"os"
)

type Client struct {
	ServerAddr        string
	MaxConcurrentJobs int64

	transportCredentials credentials.TransportCredentials

	tempDir string

	jobs    chan *v1.Job
	results chan *v1.Result
}

func NewClient(serverAddr string, maxConcurrentJobs int64) *Client {
	dir, err := os.MkdirTemp("", "TranscoderTempDir")
	if err != nil {
		log.Fatal(err)
	}
	client := &Client{
		ServerAddr:        serverAddr,
		MaxConcurrentJobs: maxConcurrentJobs,
		tempDir:           dir,
		jobs:              make(chan *v1.Job),
		results:           make(chan *v1.Result),
	}

	for i := 0; i < int(maxConcurrentJobs); i++ {
		go client.processJobs()
	}
	go client.uploadResult(context.TODO())

	return client
}

func (c *Client) JoinCluster(ctx context.Context) error {
	var opts []grpc.DialOption
	if c.transportCredentials != nil {
		opts = append(opts, grpc.WithTransportCredentials(c.transportCredentials))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(4294967296)))
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(4294967296)))

	conn, err := grpc.Dial(c.ServerAddr, opts...)
	if err != nil {
		return ErrConnectionFailed
	}
	defer conn.Close()

	client := v1.NewDistributorServiceClient(conn)

	clusterClient, err := client.JoinCluster(ctx, &v1.ClusterClientOffer{MaxConcurrentJobs: c.MaxConcurrentJobs})
	if err != nil {
		return ErrCouldNotJoinCluster
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("Context is done")
			return nil
		default:
			job, err := clusterClient.Recv()
			if err == nil {
				c.jobs <- job
			} else {
				code, ok := status.FromError(err)
				if !ok {
					log.Println(err)
				}
				switch code.Code() {
				case codes.Unavailable:
					return nil
				}
			}
		}
	}
}

func (c *Client) uploadResult(ctx context.Context) error {
	var opts []grpc.DialOption
	if c.transportCredentials != nil {
		opts = append(opts, grpc.WithTransportCredentials(c.transportCredentials))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(4294967296)))
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(4294967296)))

	conn, err := grpc.Dial(c.ServerAddr, opts...)
	if err != nil {
		return ErrConnectionFailed
	}
	defer conn.Close()

	client := v1.NewDistributorServiceClient(conn)

	for res := range c.results {
		log.Println("Job done, returning result")
		_, err := client.UploadResult(ctx, res, grpc.MaxCallSendMsgSize(4294967296), grpc.MaxCallRecvMsgSize(4294967296))
		if err != nil {
			log.Println(err)
			return ErrCouldNotUploadResult
		}
	}
	return nil
}

func (c *Client) processJobs() {
	for job := range c.jobs {
		log.Println("Got new job:", job.Id)
		inputFileName := fmt.Sprintf("%s/%s", c.tempDir, job.InputFileName)

		log.Println("Writing data to file")
		f, err := os.Create(inputFileName)
		if err != nil {
			log.Println(err)
			c.results <- &v1.Result{
				JobId:              job.Id,
				JobReferenceNumber: job.ReferenceNumber,
				StatusCode:         v1.UploadStatusCode_Failed,
				OutputFileData:     []byte{},
			}
			continue
		}

		n, err := f.Write(job.InputFileData)
		if err != nil {
			log.Println(err)
			c.results <- &v1.Result{
				JobId:              job.Id,
				JobReferenceNumber: job.ReferenceNumber,
				StatusCode:         v1.UploadStatusCode_Failed,
				OutputFileData:     []byte{},
			}
			f.Close()
			continue
		}
		if n != len(job.InputFileData) {
			log.Println("n not equal to len(job.InputFileData")
			c.results <- &v1.Result{
				JobId:              job.Id,
				JobReferenceNumber: job.ReferenceNumber,
				StatusCode:         v1.UploadStatusCode_Failed,
				OutputFileData:     []byte{},
			}
			f.Close()
			continue
		}
		f.Close()

		log.Println("Setting ffmpeg args")
		args := ffmpeg.KwArgs{}
		id := uuid.New().String()
		ft := ""
		if job.TranscodingSettings != nil {
			switch job.TranscodingSettings.MediaFileType {
			case v1.MediaFileType_Mp4:
				ft = "mp4"
			case v1.MediaFileType_Mkv:
				ft = "mkv"
			case v1.MediaFileType_Webm:
				ft = "webm"
			default:
				ft = "mkv"
			}

			switch job.TranscodingSettings.VideoSettings.Codec {
			case v1.VideoCodec_Libx264:
				args["c:v"] = "libx264"
			case v1.VideoCodec_Libx265:
				args["c:v"] = "libx264"
			case v1.VideoCodec_Vp8:
				args["c:v"] = "libx264"
			case v1.VideoCodec_Vp9:
				args["c:v"] = "libx264"
			case v1.VideoCodec_Vpx:
				args["c:v"] = "libx264"
			default:
				args["c:v"] = "copy"
			}

			switch job.TranscodingSettings.AudioSettings.Codec {
			case v1.AudioCodec_Aac:
				args["c:a"] = "aac"
			case v1.AudioCodec_Ac3:
				args["c:a"] = "ac3"
			case v1.AudioCodec_Opus:
				args["c:a"] = "libopus"
			default:
				args["c:a"] = "copy"
			}
		}

		outFileName := fmt.Sprintf("%s/%s.%s", c.tempDir, id, ft)
		log.Println("Running ffmpeg")
		err = ffmpeg.Input(inputFileName).
			Output(outFileName, args).
			OverWriteOutput().
			Run()
		if err != nil {
			log.Println(err)
			c.results <- &v1.Result{
				JobId:              job.Id,
				JobReferenceNumber: job.ReferenceNumber,
				StatusCode:         v1.UploadStatusCode_Failed,
				OutputFileData:     []byte{},
			}
			continue
		}

		outFile, err := os.Open(outFileName)
		if err != nil {
			log.Println(err)
			c.results <- &v1.Result{
				JobId:              job.Id,
				JobReferenceNumber: job.ReferenceNumber,
				StatusCode:         v1.UploadStatusCode_Failed,
				OutputFileData:     []byte{},
			}
			continue
		}

		data, err := io.ReadAll(outFile)
		if err != nil {
			log.Println(err)
			c.results <- &v1.Result{
				JobId:              job.Id,
				JobReferenceNumber: job.ReferenceNumber,
				StatusCode:         v1.UploadStatusCode_Failed,
				OutputFileData:     []byte{},
			}
			outFile.Close()
			continue
		}
		outFile.Close()
		c.results <- &v1.Result{
			JobId:              job.Id,
			JobReferenceNumber: job.ReferenceNumber,
			StatusCode:         v1.UploadStatusCode_Ok,
			OutputFileData:     data,
		}
	}
}
