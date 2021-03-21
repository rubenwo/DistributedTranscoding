package distributor

type Distributor interface {
	Distribute(files []string) error
}
