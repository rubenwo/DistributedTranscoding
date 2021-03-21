package transcoder

import "errors"

var (
	ErrConnectionFailed     = errors.New("couldn't connect to distributor server")
	ErrCouldNotJoinCluster  = errors.New("couldn't join the transcoding cluster")
	ErrCouldNotUploadResult = errors.New("couldn't upload results to the host server")
)
