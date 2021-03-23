package splitter

import (
	"errors"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrCouldNotCreateDir           = errors.New("couldn't create temporary directory")
	ErrCouldNotSplitVideo          = errors.New("couldn't split the video file")
	ErrCouldNotProbeVideo          = errors.New("couldn't probe video file")
	ErrCouldNotWalkOutputDirectory = errors.New("couldn't walk the output directory")
)

// SplitVideoTempDir splits a video into smaller videos and writes them to a temp dir
// Returns array of the paths to the individual chunks, the manifest directory and
func SplitVideoTempDir(inFileName string) ([]string, error) {
	path, err := os.MkdirTemp("", "SplitVideoTempDir")
	if err != nil {
		log.Println(err)
		return nil, ErrCouldNotCreateDir
	}

	_, err = ffmpeg.Probe(inFileName)
	if err != nil {
		return nil, ErrCouldNotProbeVideo
	}

	err = ffmpeg.Input(inFileName).
		Output(path+"/output%03d.mp4", ffmpeg.KwArgs{"c": "copy", "segment_time": "00:00:10", "f": "segment"}).
		OverWriteOutput().
		Run()
	if err != nil {
		return nil, ErrCouldNotSplitVideo
	}

	var files []string

	root := path
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files = append(files, strings.ReplaceAll(path, "\\", "/"))
		return nil
	})
	if err != nil {
		return nil, ErrCouldNotWalkOutputDirectory
	}
	fmt.Println(files)
	return files, nil
}
