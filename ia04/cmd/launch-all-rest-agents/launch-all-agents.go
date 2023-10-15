package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	const n = 100
	const url1 = ":8080"
	const url2 = "http://localhost:8080"
	ops := [...]string{"+", "-", "*"}

	clAgts := make([]restclientagent.RestClientAgent, 0, n)
	servAgt := restserveragent.NewRestServerAgent(url1)

	log.Println("démarrage du serveur...")
	go servAgt.Start()

	log.Println("démarrage des clients...")
	for i := 0; i < n; i++ {
		id := fmt.Sprintf("id%02d", i)
		op := ops[rand.Intn(3)]
		op1 := rand.Intn(100)
		op2 := rand.Intn(100)
		agt := restclientagent.NewRestClientAgent(id, url2, op, op1, op2)
		clAgts = append(clAgts, *agt) //Fais un slice d'agents
	}

	for _, agt := range clAgts {
		// attention, obligation de passer par cette lambda pour faire capturer la valeur de l'itération par la goroutine
		// pour récupérer la bonne valeur du pointeur qui va sur l'agent
		func(agt restclientagent.RestClientAgent) {
			go agt.Start()
		}(agt)
	}

	fmt.Scanln()
}
