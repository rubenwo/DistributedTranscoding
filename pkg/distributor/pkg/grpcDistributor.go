package pkg

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	v1 "github.com/rubenwo/DistributedTranscoding/pkg/api/v1"
	"github.com/rubenwo/DistributedTranscoding/pkg/distributor/pkg/registry"
	"github.com/rubenwo/DistributedTranscoding/pkg/distributor/pkg/splitter"
	"io/ioutil"
	"log"
	"strings"
)

type GrpcDistributor struct {
	Registry *registry.Registry
}

func NewGrpcDistributor() *GrpcDistributor {
	return &GrpcDistributor{
		Registry: registry.NewRegistry(),
	}
}

func (d *GrpcDistributor) JobIds() []int {
	panic("JobIds not implemented")
	return nil
}
func (d *GrpcDistributor) AddTranscodeJob(path string) error {
	// TODO: use registry
	filePaths, err := splitter.SplitVideoTempDir(path)
	if err != nil {
		return err
	}
	for i, filePath := range filePaths {
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Println(err)
			continue
		}

		split := strings.Split(filePath, "/")
		fmt.Println(split)
		job := &v1.Job{
			Id:              uuid.New().String(),
			ReferenceNumber: int64(i),
			TranscodingSettings: &v1.TranscodingSettings{
				VideoSettings: &v1.VideoSettings{
					Codec: v1.VideoCodec_Libx264,
				},
				AudioSettings: &v1.AudioSettings{},
				MediaFileType: v1.MediaFileType_Mp4,
			},
			InputFileName: split[len(split)-1],
			InputFileData: data,
		}

		d.Registry.Jobs <- job
	}
	return nil
}

func (d *GrpcDistributor) TranscodeJobProgress(id int) (<-chan Progress, error) {
	panic("TranscodeJobProgress not implemented")
	return nil, nil
}

func (d *GrpcDistributor) JoinCluster(req *v1.ClusterClientOffer, srv v1.DistributorService_JoinClusterServer) error {
	for {
		select {
		case <-srv.Context().Done():
			log.Println("Context->Done(), closing stream")
			return nil
		case job := <-d.Registry.Jobs:
			fmt.Println("sending job")
			if err := srv.Send(job); err != nil {
				return err
			}
		}
	}
}

func (d *GrpcDistributor) UploadResult(ctx context.Context, req *v1.Result) (*v1.Empty, error) {
	d.Registry.Results <- req
	return &v1.Empty{}, nil
}
