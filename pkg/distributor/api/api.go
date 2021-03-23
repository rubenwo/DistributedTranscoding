package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/websocket"
	v1 "github.com/rubenwo/DistributedTranscoding/pkg/api/v1"
	"github.com/rubenwo/DistributedTranscoding/pkg/distributor/api/config"
	"github.com/rubenwo/DistributedTranscoding/pkg/distributor/pkg"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

type api struct {
	distributor pkg.Distributor
}

func Run(cfg *config.Configuration) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		return fmt.Errorf("failed to listen on port 8080")
	}
	grpcServer := grpc.NewServer(grpc.MaxRecvMsgSize(4294967296), grpc.MaxSendMsgSize(4294967296))
	dist := pkg.NewGrpcDistributor()
	v1.RegisterDistributorServiceServer(grpcServer, dist)
	go func() {
		log.Println("grpc server started listening on port:", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
		log.Println("grpc server stopped listening")
	}()
	a := &api{distributor: dist}

	router := chi.NewRouter()
	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	router.Get("/healthz", a.healthz)
	router.Post("/job", a.createJob)

	router.Get("/job/{}/status", a.jobStatus)

	router.Get("/ws/jobs/status", a.streamJobStatus)

	router.Get("/results/{id}", a.result)

	log.Printf("Distributor is listening on: %s\n", cfg.Addr)
	if err := http.ListenAndServe(cfg.Addr, router); err != nil {
		return fmt.Errorf("ListenAndServe: %w", err)
	}
	return nil
}

func (a *api) healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(&HealthzModel{
		IsHealthy:    true,
		ErrorMessage: "",
	}); err != nil {
		log.Printf("error sending healthz: %s\n", err.Error())
	}
}

func (a *api) createJob(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20) // limit your max input length!
	if err != nil {
		panic(err)
	}
	// in your case file would be fileupload
	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	name := strings.Split(header.Filename, ".")
	fmt.Printf("File name %s\n", name[0])
	// Copy the file data to a file
	f, err := os.Create("./assets/" + name[0] + "." + name[len(name)-1])
	if err != nil {
		panic(err)
	}
	defer f.Close()
	n, err := io.Copy(f, file)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

	id, err := a.distributor.AddTranscodeJob("./assets/" + name[0] + "." + name[len(name)-1])
	if err != nil {
		panic(err)
	}

	var msg struct {
		Id string `json:"id"`
	}
	msg.Id = id
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(&msg); err != nil {
		panic(err)
	}
}

func (a *api) result(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	path, err := a.distributor.RetrieveFilePath(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, path)
}

func (a *api) jobStatus(w http.ResponseWriter, r *http.Request) {

}

func (a *api) streamJobStatus(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not upgrade request: %s", err.Error()), http.StatusBadRequest)
		return
	}

	jobIds := a.distributor.JobIds()
	results := make(chan pkg.Progress, 10)
	for _, id := range jobIds {
		go func(rc chan<- pkg.Progress) {
			c, err := a.distributor.TranscodeJobProgress(id)
			if err != nil {
				http.Error(w, fmt.Sprintf("could find progress for job: %d", id), http.StatusBadRequest)
				return
			}
			for p := range c {
				rc <- p
			}

		}(results)
	}

	for r := range results {
		var msg struct {
			CurrentJobState pkg.State `json:"state"`
			Error           string    `json:"error"`
		}
		state, err := r.CurrentJobState()
		msg.CurrentJobState = state
		if err != nil {
			msg.Error = err.Error()
		}

		data, err := json.Marshal(&msg)
		if err != nil {
			_ = conn.WriteMessage(websocket.TextMessage, []byte("error marshalling msg"))
			continue
		}

		_ = conn.WriteMessage(websocket.TextMessage, data)
	}
}
