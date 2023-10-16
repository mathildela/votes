package main

import (
	"fmt"
	"ia04/agt/ballotagent"
	"ia04/agt/voteragent"
	"ia04/comsoc"
	"log"
)

func main() {
	const n = 3
	const url1 = ":8080"
	const url2 = "http://localhost:8080"

	clAgts := make([]voteragent.RestClientAgent, 0, n)
	servAgt := ballotagent.NewRestServerAgent(url1)

	log.Println("démarrage du serveur...")
	go servAgt.Start()

	voter_ids := []string{"a1", "a2", "a3"}
	tiebreak := []comsoc.Alternative{1, 2, 3, 4}
	servAgt.NewBallot("ballot1", "majority", "2023-10-09T23:05:08+02:00", 4, voter_ids, tiebreak)

	prefs1 := []comsoc.Alternative{1, 2, 3, 4}
	prefs2 := []comsoc.Alternative{1, 2, 4, 3}
	prefs3 := []comsoc.Alternative{2, 1, 4, 3}
	var opt []int

	clAgts = append(clAgts, *voteragent.NewRestClientAgent("a1", url2, "ballot1", prefs1, opt))
	clAgts = append(clAgts, *voteragent.NewRestClientAgent("a2", url2, "ballot1", prefs2, opt))
	clAgts = append(clAgts, *voteragent.NewRestClientAgent("a3", url2, "ballot1", prefs3, opt))
	/*
		log.Println("démarrage des clients...")
		for i := 0; i < n; i++ {
			id := fmt.Sprintf("id%02d", i)
			op := ops[rand.Intn(3)]
			op1 := rand.Intn(100)
			op2 := rand.Intn(100)
			agt := restclientagent.NewRestClientAgent(id, url2, op, op1, op2)
			clAgts = append(clAgts, *agt) //Fais un slice d'agents
		}*/

	for _, agt := range clAgts {
		// attention, obligation de passer par cette lambda pour faire capturer la valeur de l'itération par la goroutine
		// pour récupérer la bonne valeur du pointeur qui va sur l'agent
		func(agt voteragent.RestClientAgent) {
			go agt.Start()
		}(agt)
	}

	var ballot ballotagent.Ballot = servAgt.GetBallot("ballot1")
	function_majority := comsoc.SWFFactory(comsoc.MajoritySWF, comsoc.TieBreakFactory(ballot.Tiebreak))
	res, err := function_majority(ballot.Prof)

	if err != nil {
		fmt.Println(res)
	} else {
		fmt.Println("Error: %s", err)
	}
	fmt.Scanln()
}
