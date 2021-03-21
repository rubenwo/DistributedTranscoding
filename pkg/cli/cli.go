package cli

import "fmt"

func Run(args ...interface{}) error {
	fmt.Println("running cli")
	return nil
}
