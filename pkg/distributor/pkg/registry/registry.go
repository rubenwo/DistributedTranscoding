package registry

import (
	v1 "github.com/rubenwo/DistributedTranscoding/pkg/api/v1"
	"log"
)

type Registry struct {
	Jobs    chan *v1.Job
	Results chan *v1.Result

	jobs map[string][]int
}

func NewRegistry() *Registry {
	registry := &Registry{
		Jobs:    make(chan *v1.Job),
		Results: make(chan *v1.Result),
	}

	go registry.processResults()

	return registry
}

func (r *Registry) processResults() {
	for res := range r.Results {
		log.Println(res.JobId, res.JobReferenceNumber, res.StatusCode.String())
	}
}

//func test() {
//	splitter.SplitVideoTempDir()
//	manifestPath, _ := stitcher.CreateManifest()
//	stitcher.StitchVideo(manifestPath, "./assets/stitched.mp4")
//}
