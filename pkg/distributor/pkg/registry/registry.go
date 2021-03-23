package registry

import (
	"bytes"
	"fmt"
	v1 "github.com/rubenwo/DistributedTranscoding/pkg/api/v1"
	"github.com/rubenwo/DistributedTranscoding/pkg/distributor/pkg/stitcher"
	"io"
	"log"
	"os"
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
		f, err := os.Create(fmt.Sprintf("./assets/transcoded/%d.mp4", res.JobReferenceNumber))
		if err != nil {
			log.Println(err)
			continue
		}
		io.Copy(f, bytes.NewBuffer(res.OutputFileData))
		f.Close()
		if res.JobReferenceNumber == 10 {
			go func() {
				//manifest, err := stitcher.CreateManifest("./assets/transcoded")
				//if err != nil {
				//	log.Fatal(err)
				//}
				manifest := "./assets/manifest.txt"
				err = stitcher.StitchVideo(manifest, "./assets/t.mp4")
				if err != nil {
					log.Fatal(err)
				}
			}()
		}
	}
}

//func test() {
//	splitter.SplitVideoTempDir()
//	manifestPath, _ := stitcher.CreateManifest()
//	stitcher.StitchVideo(manifestPath, "./assets/stitched.mp4")
//}
