package pkg

type State uint8

const (
	Backlog State = iota
	Distributing
	Transcoding
)

type Progress interface {
	CurrentJobState() (State, error)
}

type Distributor interface {
	// JobIds should return a list of all the current jobs in the distributor
	JobIds() []int
	// AddTranscodeJob gets a path to a media file and should create a new job. Return an error if something went wrong
	AddTranscodeJob(path string) error
	// TranscodeJobProgress gets an id and returns the current progress.
	TranscodeJobProgress(id int) (<-chan Progress, error)
}
