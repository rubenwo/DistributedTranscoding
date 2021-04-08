package pkg

import (
	"context"
	"fmt"
	v1 "github.com/rubenwo/DistributedTranscoding/pkg/api/v1"
	"github.com/rubenwo/DistributedTranscoding/pkg/distributor/pkg/distributor"
	"github.com/rubenwo/DistributedTranscoding/pkg/distributor/pkg/registry"
	"log"
)

type GrpcDistributor struct {
	Registry *registry.Registry
	v1.UnimplementedDistributorServiceServer
}

func NewGrpcDistributor() *GrpcDistributor {
	return &GrpcDistributor{
		Registry: registry.NewRegistry(),
	}
}

func (d *GrpcDistributor) JobIds() []string {
	return d.Registry.JobIds()
}
func (d *GrpcDistributor) AddTranscodeJob(path string) (string, error) {
	return d.Registry.AddJob(path)
}

func (d *GrpcDistributor) TranscodeJobProgress(id string) (<-chan distributor.Progress, error) {
	channels := d.Registry.ProgressChannels([]string{id})
	return channels[0], nil
}

func (d *GrpcDistributor) RetrieveFilePath(id string) (string, error) {
	return d.Registry.RetrieveFile(id)
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
