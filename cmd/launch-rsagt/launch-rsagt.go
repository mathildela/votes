package main

import (
	"fmt"

	"gitlab.utc.fr/langemat/ia04/agt/ballotagent"
)

func main() {
	server := ballotagent.NewRestServerAgent(":8080")
	server.Start()
	fmt.Scanln()
}
