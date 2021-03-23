package stitcher

import (
	"errors"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
	"path/filepath"
)

var (
	ErrCouldNotGenerateManifest    = errors.New("couldn't generate the manifest file")
	ErrCouldNotWalkOutputDirectory = errors.New("couldn't walk the output directory")
	ErrCouldNotStitchVideo         = errors.New("couldn't stitch the video together")
)

func CreateManifest(dir string) (string, error) {
	var files []string

	root := dir
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		return "", ErrCouldNotWalkOutputDirectory
	}

	f, err := os.Create(dir + "/manifest.txt")
	if err != nil {
		return "", ErrCouldNotGenerateManifest
	}
	str := ""
	for _, fileName := range files {
		str += fmt.Sprintf("file '%s'\n", fileName)
	}

	_, err = f.Write([]byte(str))
	if err != nil {
		return "", ErrCouldNotGenerateManifest
	}
	return dir + "/manifest.txt", nil
}

func StitchVideo(manifestPath, outFileName string) error {
	err := ffmpeg.Input(manifestPath, ffmpeg.KwArgs{"f": "concat", "safe": "0"}).
		Output(outFileName, ffmpeg.KwArgs{"c": "copy"}).
		OverWriteOutput().
		Run()
	if err != nil {
		return ErrCouldNotStitchVideo
	}
	return nil
}
