package registry

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	v1 "github.com/rubenwo/DistributedTranscoding/pkg/api/v1"
	"github.com/rubenwo/DistributedTranscoding/pkg/distributor/pkg/distributor"
	"github.com/rubenwo/DistributedTranscoding/pkg/distributor/pkg/splitter"
	"github.com/rubenwo/DistributedTranscoding/pkg/distributor/pkg/stitcher"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type Registry struct {
	Jobs    chan *v1.Job
	Results chan *v1.Result

	jobs    map[string]int
	results map[string]int

	finished map[string]string
}

func NewRegistry() *Registry {
	registry := &Registry{
		Jobs:     make(chan *v1.Job),
		Results:  make(chan *v1.Result),
		jobs:     make(map[string]int),
		results:  make(map[string]int),
		finished: make(map[string]string),
	}

	go registry.processResults()

	return registry
}

func (r *Registry) JobIds() []string {
	var ids []string
	for k := range r.jobs {
		ids = append(ids, k)
	}

	return ids
}

type EasyProgress struct {
	currentChunks int
	maxChunks     int
}

func (e *EasyProgress) CurrentJobState() (distributor.State, error) {
	return distributor.Transcoding, nil
}

func (e *EasyProgress) CurrentChunks() int {
	return e.currentChunks
}

func (e *EasyProgress) MaxChunks() int {
	return e.maxChunks
}

func (r *Registry) ProgressChannels(ids []string) []chan distributor.Progress {
	channels := make([]chan distributor.Progress, len(ids))
	for i, id := range ids {
		go func(idx int, id string) {
			for x := 0; x < 100; x++ {
				channels[idx] <- &EasyProgress{
					currentChunks: r.results[id],
					maxChunks:     r.jobs[id],
				}
				time.Sleep(time.Second)
			}
		}(i, id)
	}

	return channels
}

func (r *Registry) RetrieveFile(id string) (string, error) {
	res, ok := r.finished[id]
	if !ok {
		return "", fmt.Errorf("id: %s does not exist", id)
	}
	if res == "" {
		return "", fmt.Errorf("job for id: %s is not done", id)
	}
	return res, nil
}

func (r *Registry) AddJob(path string) (string, error) {
	jobId := uuid.New().String()
	filePaths, err := splitter.SplitVideoTempDir(path)
	if err != nil {
		return "", err
	}
	r.jobs[jobId] = len(filePaths)
	r.finished[jobId] = ""

	go func(fp []string, id string) {
		for i, filePath := range fp {
			data, err := ioutil.ReadFile(filePath)
			if err != nil {
				log.Println(err)
				continue
			}

			split := strings.Split(filePath, "/")
			fmt.Println(split)
			job := &v1.Job{
				Id:              id,
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

			r.Jobs <- job
		}
	}(filePaths, jobId)

	return jobId, nil
}

func (r *Registry) processResults() {
	dir, err := os.MkdirTemp("", "ProcessedTempDir")
	if err != nil {
		log.Fatal(err)
	}
	// TODO: Make separate directory for each job

	for res := range r.Results {
		//dir += "/" + res.JobId
		log.Println(res.JobId, res.JobReferenceNumber, res.StatusCode.String())
		f, err := os.Create(fmt.Sprintf("%s/%d.mp4", dir, res.JobReferenceNumber))
		if err != nil {
			log.Println(err)
			continue
		}
		io.Copy(f, bytes.NewBuffer(res.OutputFileData))
		f.Close()
		r.results[res.JobId]++

		log.Printf("got %d chunks out of: %d", r.results[res.JobId], r.jobs[res.JobId])

		if r.results[res.JobId] == r.jobs[res.JobId] {
			go func() {
				manifest, err := stitcher.CreateManifest(dir)
				if err != nil {
					log.Fatal(err)
				}
				dir, err := os.MkdirTemp("", "Finished")
				if err != nil {
					log.Fatal(err)
				}
				err = stitcher.StitchVideo(manifest, dir+"/finished.mp4")
				if err != nil {
					log.Fatal(err)
				}
				r.finished[res.JobId] = dir + "/finished.mp4"
			}()
		}
	}
}

//func test() {
//	splitter.SplitVideoTempDir()
//	manifestPath, _ := stitcher.CreateManifest()
//	stitcher.StitchVideo(manifestPath, "./assets/stitched.mp4")
//}
