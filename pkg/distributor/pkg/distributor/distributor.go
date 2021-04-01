package distributor

type State uint8

const (
	Backlog State = iota
	Distributing
	Transcoding
)

type Progress interface {
	CurrentJobState() (State, error)
	CurrentChunks() int
	MaxChunks() int
}

type Distributor interface {
	// JobIds should return a list of all the current jobs in the distributor
	JobIds() []string
	// AddTranscodeJob gets a path to a media file and should create a new job. Return a job id
	AddTranscodeJob(path string) (string, error)
	// TranscodeJobProgress gets an id and returns the current progress.
	TranscodeJobProgress(id string) (<-chan Progress, error)
	// RetrieveFilePath returns the filepath where the transcoded file corresponding to the id is located
	RetrieveFilePath(id string) (string, error)
}
