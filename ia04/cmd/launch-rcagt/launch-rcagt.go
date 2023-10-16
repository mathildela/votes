package main

import (
	"fmt"
	"ia04/agt/voteragent"
)

func main() {
	ag := voteragent.NewRestClientAgent("id1", "http://localhost:8080", "+", 11, 1)
	ag.Start()
	fmt.Scanln()
}
