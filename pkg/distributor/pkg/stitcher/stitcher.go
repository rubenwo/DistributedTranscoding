package stitcher

import (
	"errors"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

var (
	ErrCouldNotGenerateManifest    = errors.New("couldn't generate the manifest file")
	ErrCouldNotWalkOutputDirectory = errors.New("couldn't walk the output directory")
	ErrCouldNotStitchVideo         = errors.New("couldn't stitch the video together")
	ErrNoManifestProvided          = errors.New("manifest was not provided")
)

func CreateManifest(dir string) (string, error) {
	var filePathFiles []string

	root := dir

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		filePathFiles = append(filePathFiles, strings.TrimPrefix(filepath.ToSlash(path), filepath.ToSlash(dir)+"/"))
		return nil
	})
	if err != nil {
		return "", ErrCouldNotWalkOutputDirectory
	}
	fmt.Println(filePathFiles)
	var nums []int
	for _, file := range filePathFiles {
		num, err := strconv.Atoi(strings.Split(file, ".")[0])
		if err != nil {
			return "", ErrCouldNotGenerateManifest
		}
		nums = append(nums, num)
	}
	sort.Ints(nums)
	outputType := strings.Split(filePathFiles[0], ".")[1]
	files := make([]string, len(nums))
	for i, n := range nums {
		files[i] = fmt.Sprintf("%d.%s", n, outputType)
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
	f.Close()
	return dir + "/manifest.txt", nil
}

func StitchVideo(manifestPath, outFileName string) error {
	if manifestPath == "" {
		return ErrNoManifestProvided
	}

	err := ffmpeg.Input(manifestPath, ffmpeg.KwArgs{"f": "concat", "safe": "0"}).
		Output(outFileName, ffmpeg.KwArgs{"c": "copy"}).
		OverWriteOutput().
		Run()
	if err != nil {
		return ErrCouldNotStitchVideo
	}
	return nil
}
