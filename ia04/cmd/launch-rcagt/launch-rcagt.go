package main

import (
	"fmt"
)

func main() {
	ag := restclientagent.NewRestClientAgent("id1", "http://localhost:8080", "+", 11, 1)
	ag.Start()
	fmt.Scanln()
}
