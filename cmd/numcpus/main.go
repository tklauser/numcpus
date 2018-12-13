package main

import (
	"fmt"

	"github.com/tklauser/numcpus"
)

func main() {
	present, _ := numcpus.GetPresent()
	online, _ := numcpus.GetOnline()
	offline, _ := numcpus.GetOffline()
	possible, _ := numcpus.GetPossible()
	fmt.Printf("present:  %v\n", present)
	fmt.Printf("online:   %v\n", online)
	fmt.Printf("offline:  %v\n", offline)
	fmt.Printf("possible: %v\n", possible)
}
