package main

import (
	"fmt"
	"ia04/agt/voteragent"
	"ia04/comsoc"
	"time"
)

const nb_agents int = 10
const nb_candidats int = 5

func main() {
	// pas de channel dans l'exemple car on va le faire en client/serveur ?
	c := make(chan []comsoc.Alternative)
	//p := make(comsoc.Profile, nb_agents)
	for i := 0; i < nb_agents; i++ {
		go func(i int) {
			a := voteragent.NewAgent("1", "A1", nb_candidats)
			//a.Start()
			new_slice := make([]comsoc.Alternative, len(a.Prefs))
			copy(new_slice, a.Prefs)
			c <- new_slice
		}(i)
	}
	time.Sleep(3 * time.Second)
	// on ne peut pas transmettre un slice dans un channel
	/*
		    for i:=0; i<nb_agents;i++{
				p[i] <- c
			}*/
	fmt.Scanln()
}
