package distributor

import (
	"context"
	"fmt"
	v1 "github.com/rubenwo/DistributedTranscoding/pkg/api/v1"
	"log"
)

type GrpcDistributor struct {
	Jobs    chan *v1.Job
	Results chan *v1.Result
}

func (d *GrpcDistributor) JoinCluster(req *v1.ClusterClientOffer, srv v1.DistributorService_JoinClusterServer) error {
	for {
		select {
		case <-srv.Context().Done():
			log.Println("Context->Done(), closing stream")
			return nil
		case job := <-d.Jobs:
			fmt.Println("sending job")
			if err := srv.Send(job); err != nil {
				return err
			}
		}
	}
}

func (d *GrpcDistributor) UploadResult(ctx context.Context, req *v1.Result) (*v1.Empty, error) {
	fmt.Println(req)
	return &v1.Empty{}, nil
}
