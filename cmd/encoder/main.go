package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"math/rand"
	"net"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ExampleShowProgress is an example of using the ffmpeg `-progress` option with a
//    unix-domain socket to report progress
func ExampleShowProgress(inFileName, outFileName string) {
	a, err := ffmpeg.Probe(inFileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
	totalDuration := gjson.Get(a, "format.duration").Float()

	err = ffmpeg.Input(inFileName).
		Output(outFileName, ffmpeg.KwArgs{"c:v": "libx264", "preset": "veryslow"}).
		GlobalArgs("-progress", "unix://"+TempSock(totalDuration)).
		OverWriteOutput().
		Run()
	if err != nil {
		log.Fatal(err)
	}

}

func TempSock(totalDuration float64) string {
	// serve
	rand.Seed(time.Now().Unix())
	sockFileName := path.Join(os.TempDir(), fmt.Sprintf("%d_sock", rand.Int()))
	l, err := net.Listen("unix", sockFileName)
	if err != nil {
		panic(err)
	}

	go func() {
		re := regexp.MustCompile(`out_time_ms=(\d+)`)
		fd, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		buf := make([]byte, 16)
		data := ""
		progress := ""
		for {
			_, err := fd.Read(buf)
			if err != nil {
				return
			}
			data += string(buf)
			a := re.FindAllStringSubmatch(data, -1)
			cp := ""
			if len(a) > 0 && len(a[len(a)-1]) > 0 {
				c, _ := strconv.Atoi(a[len(a)-1][len(a[len(a)-1])-1])
				cp = fmt.Sprintf("%.2f", float64(c)/totalDuration/1000000)
			}
			if strings.Contains(data, "progress=end") {
				cp = "done"
			}
			if cp == "" {
				cp = ".0"
			}
			if cp != progress {
				progress = cp
				fmt.Println("progress: ", progress)
			}
		}
	}()

	return sockFileName
}

func splitFile(inFileName, outFileDir string) error {
	a, err := ffmpeg.Probe(inFileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)
	err = ffmpeg.Input(inFileName).
		Output(outFileDir+"/output%03d.mp4", ffmpeg.KwArgs{"c": "copy", "segment_time": "00:00:10", "f": "segment"}).
		OverWriteOutput().
		Run()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func transcodeParallel(fileNames []string, outFileDir string) error {
	args := ffmpeg.KwArgs{"c:v": "libx264", "preset": "slow"}
	var wg sync.WaitGroup
	for i, fileName := range fileNames {
		wg.Add(1)
		go func(idx int, fileName string) {
			err := ffmpeg.Input(fileName).
				Output(fmt.Sprintf("%s/output%03d.mp4", outFileDir, idx), args).
				OverWriteOutput().
				Run()
			if err != nil {
				log.Println(err)
			}
			wg.Done()
		}(i, fileName)
	}
	wg.Wait()
	return nil
}

func stitchVideosTogether(fileNames []string, outFileName string) error {
	f, err := os.Create("./assets/manifest.txt")
	if err != nil {
		return err
	}
	str := ""
	for _, fileName := range fileNames {
		str += fmt.Sprintf("file '%s'\n", fileName)
	}

	_, err = f.Write([]byte(str))
	if err != nil {
		return err
	}

	err = ffmpeg.Input("./assets/manifest.txt", ffmpeg.KwArgs{"f": "concat", "safe": "0"}).
		Output(outFileName, ffmpeg.KwArgs{"c": "copy"}).
		OverWriteOutput().
		Run()
	if err != nil {
		log.Println(err)
	}

	return nil
}

func main() {
	//ExampleShowProgress("./assets/demo.mp4", "./assets/demo_out.mp4")
	if err := os.Mkdir("./assets/split", 0644); err != nil {
		log.Fatal(err)
	}
	if err := os.Mkdir("./assets/encoded", 0644); err != nil {
		log.Fatal(err)
	}

	if err := splitFile("./assets/big_buck_bunny.mp4", "./assets/split"); err != nil {
		log.Fatal(err)
	}
	if err := transcodeParallel([]string{
		"./assets/split/output000.mp4",
		"./assets/split/output001.mp4",
		"./assets/split/output002.mp4",
		"./assets/split/output003.mp4",
		"./assets/split/output004.mp4",
		"./assets/split/output005.mp4",
		"./assets/split/output006.mp4",
		"./assets/split/output007.mp4",
		"./assets/split/output008.mp4",
		"./assets/split/output009.mp4",
		"./assets/split/output010.mp4",
	}, "./assets/encoded"); err != nil {
		log.Fatal(err)
	}

	if err := stitchVideosTogether([]string{
		"./split/output000.mp4",
		"./split/output001.mp4",
		"./split/output002.mp4",
		"./split/output003.mp4",
		"./split/output004.mp4",
		"./split/output005.mp4",
		"./split/output006.mp4",
		"./split/output007.mp4",
		"./split/output008.mp4",
		"./split/output009.mp4",
		"./split/output010.mp4",
	}, "./assets/stitched.mp4"); err != nil {
		log.Fatal(err)
	}
}
