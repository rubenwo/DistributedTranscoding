package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/apex/log"
	"github.com/gorilla/websocket"
	"github.com/rubenwo/DistributedTranscoding/pkg/distributor/pkg/distributor"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type Client struct{}

func New() *Client { return &Client{} }

func t() error {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost/ws/jobs/status", nil)
	if err != nil {
		return err
	}

	var msg struct {
		CurrentJobState distributor.State `json:"state"`
		Error           string            `json:"error"`
	}

	fmt.Println(msg)
	fmt.Println(conn)
	return nil
}

func (c *Client) AddJob(path string) error {
	log.Infof("opening file: %s", path)
	file, err := os.Open(path)

	if err != nil {
		return err
	}
	defer file.Close()

	log.Debugf("creating writer...")
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))

	if err != nil {
		return err
	}

	log.Debugf("writing file contents to writer buffer")
	io.Copy(part, file)
	writer.Close()
	log.Debugf("creating POST request")
	request, err := http.NewRequest("POST", "http://localhost/job", body)

	if err != nil {
		return err
	}

	log.Debugf("setting request headers")
	request.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}

	log.Infof("sending request to the distributor")
	response, err := client.Do(request)

	if err != nil {
		return err
	}
	defer response.Body.Close()

	log.Infof("reading response from server")
	var msg struct {
		Id string `json:"id"`
	}
	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return err
	}

	log.Infof("cluster is now transcoding. job id: %s", msg.Id)
	return nil
}
