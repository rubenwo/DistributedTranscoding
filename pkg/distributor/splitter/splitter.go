package splitter

import (
	"errors"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
)

var (
	ErrCouldNotCreateDir        = errors.New("couldn't create temporary directory")
	ErrCouldNotSplitVideo       = errors.New("couldn't split the video file")
	ErrCouldNotProbeVideo       = errors.New("couldn't probe video file")
	ErrCouldNotGenerateManifest = errors.New("couldn't generate the manifest file")
)

// SplitVideoTempDir splits a video into smaller videos and writes them to a temp dir, which is then returned
func SplitVideoTempDir(inFileName string) (string, error) {
	path, err := os.MkdirTemp("", "SplitVideoTempDir")
	if err != nil {
		log.Println(err)
		return "", ErrCouldNotCreateDir
	}

	_, err = ffmpeg.Probe(inFileName)
	if err != nil {
		return "", ErrCouldNotProbeVideo
	}

	err = ffmpeg.Input(inFileName).
		Output(path+"/output%03d.mp4", ffmpeg.KwArgs{"c": "copy", "segment_time": "00:00:10", "f": "segment"}).
		OverWriteOutput().
		Run()
	if err != nil {
		return "", ErrCouldNotSplitVideo
	}

	return path, nil
}
