package main

import (
	"fmt"
	v1 "github.com/rubenwo/DistributedTranscoding/pkg/api/v1"
	"github.com/rubenwo/DistributedTranscoding/pkg/distributor"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	data, err := os.ReadFile("./assets/big_buck_bunny.mp4")
	if err != nil {
		log.Fatal(err)
	}

	job := &v1.Job{
		Id:              "",
		ReferenceNumber: 0,
		TranscodingSettings: &v1.TranscodingSettings{
			VideoSettings: &v1.VideoSettings{
				Codec: v1.VideoCodec_Libx264,
			},
			AudioSettings: &v1.AudioSettings{},
			MediaFileType: v1.MediaFileType_Mp4,
		},
		InputFileName: "big_buck_bunny.mp4",
		InputFileData: data,
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatal("can't listen on port: 8443", err)
	}
	grpcServer := grpc.NewServer()
	dist := &distributor.GrpcDistributor{
		Jobs:    make(chan *v1.Job),
		Results: make(chan *v1.Result),
	}
	go func() {
		time.Sleep(time.Second * 5)
		dist.Jobs <- job
	}()
	v1.RegisterDistributorServiceServer(grpcServer, dist)
	log.Println("grpc server started listening on port:", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
	log.Println("grpc server stopped listening")

}
