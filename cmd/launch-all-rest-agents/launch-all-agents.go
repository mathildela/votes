package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"gitlab.utc.fr/langemat/ia04/agt/voteragent"

	"gitlab.utc.fr/langemat/ia04/comsoc"

	"gitlab.utc.fr/langemat/ia04/agt/ballotagent"
)

type DataNewBallot struct {
	Rule     string               `json:"rule"`
	Deadline string               `json:"deadline"`
	VoterIDs []string             `json:"voter-ids"`
	Alts     int                  `json:"#alts"`
	TieBreak []comsoc.Alternative `json:"tie-break"`
}

type DataResult struct {
	Ballot_id string `json:"ballot-id"`
}

func generatePrefs(nb_max int) []comsoc.Alternative {
	if nb_max < 1 {
		return nil
	}
	generatedNumbers := make(map[comsoc.Alternative]struct{})
	uniqueNumbers := []comsoc.Alternative{}
	for len(uniqueNumbers) < nb_max {
		num := comsoc.Alternative(rand.Intn(nb_max) + 1)
		if _, exists := generatedNumbers[num]; !exists {
			uniqueNumbers = append(uniqueNumbers, num)
			generatedNumbers[num] = struct{}{}
		}
	}
	return uniqueNumbers
}

func generateAgentIDs(n int) []string {
	if n <= 0 {
		return nil
	}
	agentIDs := make([]string, n)
	for i := 1; i <= n; i++ {
		agentIDs[i-1] = fmt.Sprintf("ag_id%02d", i)
	}
	return agentIDs
}

func AddBallot(url_server string, rule string, deadline string, voter_ids []string, alts int, tiebreak []comsoc.Alternative) (string, error) {
	data := DataNewBallot{
		Rule:     rule,
		Deadline: deadline,
		VoterIDs: voter_ids,
		Alts:     alts,
		TieBreak: tiebreak,
	}

	// Convertir la structure en JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Erreur lors de la conversion en JSON :", err)
		return "", err
	}

	// Créer un objet bytes.Buffer pour contenir les données JSON
	buffer := bytes.NewBuffer(jsonData)
	// Effectuer la requête POST avec http.Post
	resp, err := http.Post(url_server+"/new_ballot", "application/json", buffer)
	if err != nil {
		fmt.Println("Erreur lors de la requête HTTP :", err)
		return "", err
	}

	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == http.StatusCreated {
		fmt.Println("La requête de création de ballot a réussi!")
	} else {
		fmt.Printf("La requête de création de ballot a échoué avec le code d'état : %d\n", resp.StatusCode)
		msg := fmt.Sprintf("Error: %d", resp.StatusCode)
		err := errors.New(msg)
		return "", err
	}

	//Lire reponse pour la création du scrutin
	defer resp.Body.Close()
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	ballot_id := fmt.Sprint(res["ballot-id"])
	return ballot_id, nil
}

func Vote(url_server string, nomscrutin string, nb_votants int, nb_alts int) (lAgts []voteragent.RestClientAgent) {
	clAgts := make([]voteragent.RestClientAgent, 0, nb_votants)

	log.Println("Démarrage des clients voters...")
	for i := 1; i <= nb_votants; i++ {
		id := fmt.Sprintf("ag_id%02d", i)
		prefs := generatePrefs(nb_alts)
		options := make([]int, 0)
		agt := voteragent.NewRestClientAgent(id, url_server, nomscrutin, prefs, options)
		clAgts = append(clAgts, *agt) //Fais un slice d'agents
	}

	for _, agt := range clAgts {
		// attention, obligation de passer par cette lambda pour faire capturer la valeur de l'itération par la goroutine
		// pour récupérer la bonne valeur du pointeur qui va sur l'agent
		func(agt voteragent.RestClientAgent) {
			go agt.Start()
		}(agt)
	}
	return clAgts
}

func VoteApproval(url_server string, nomscrutin string, nb_votants int, nb_alts int) (lAgts []voteragent.RestClientAgent) {
	clAgts := make([]voteragent.RestClientAgent, 0, nb_votants)

	log.Println("Démarrage des clients voters...")
	for i := 1; i <= nb_votants; i++ {
		id := fmt.Sprintf("ag_id%02d", i)
		prefs := generatePrefs(nb_alts)
		options := make([]int, 0)
		options = append(options, rand.Intn(nb_alts)+1)
		agt := voteragent.NewRestClientAgent(id, url_server, nomscrutin, prefs, options)
		clAgts = append(clAgts, *agt) //Fais un slice d'agents
	}

	for _, agt := range clAgts {
		// attention, obligation de passer par cette lambda pour faire capturer la valeur de l'itération par la goroutine
		// pour récupérer la bonne valeur du pointeur qui va sur l'agent
		func(agt voteragent.RestClientAgent) {
			go agt.Start()
		}(agt)
	}
	return clAgts
}

func getResult(url_serveur string, nomscrutin string) (map[string]interface{}, error) {
	data := DataResult{
		Ballot_id: nomscrutin,
	}
	// Convertir la structure en JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Erreur lors de la conversion en JSON :", err)
		return nil, err
	}

	// Créer un objet bytes.Buffer pour contenir les données JSON
	buffer := bytes.NewBuffer(jsonData)

	resp, err := http.Post(url_serveur+"/result", "application/json", buffer)
	if err != nil {
		fmt.Println("Erreur lors de la requête HTTP :", err)
		return nil, err
	}
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == http.StatusOK {
		fmt.Println("La requête de résultat a réussi!")
	} else {
		fmt.Printf("La requête a échoué avec le code d'état : %s\n", resp.Status)
		return nil, errors.New(resp.Status)
	}

	//Lire reponse pour la création du scrutin
	defer resp.Body.Close()
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	return res, nil
}

func main() {
	const url1 = ":8080"
	const url2 = "http://localhost:8080"
	const nb_alts = 4
	const attente = 3

	var n int
	fmt.Print("\nEntrez le nombre de votants : ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	voters_ids := generateAgentIDs(n)
	alts := nb_alts
	tiebreak := generatePrefs(alts)

	servAgt := ballotagent.NewRestServerAgent(url1)

	log.Println("Démarrage du serveur...")
	go servAgt.Start()

	time.Sleep(time.Second)

	// ************************** MAJORITE **************************
	// newBallot
	fmt.Println("\n######## MAJORITY ########")

	rule := "majority"

	deadline_t := time.Now().Add(time.Second * attente)
	deadline := deadline_t.Format(time.RFC3339)

	nomscrutin, err := AddBallot(url2, rule, deadline, voters_ids, alts, tiebreak)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Printf("> %s créé [méthode %s, %d alts, deadline %s, tie-break %v]\n", nomscrutin, rule, alts, deadline, tiebreak)
	}

	// Vote
	clAgts := Vote(url2, nomscrutin, n, nb_alts)
	if clAgts == nil {
		fmt.Println("Error")
		return
	}

	fmt.Printf("\n -- Attente de %ds pour la fin de la deadline -- \n", attente)
	time.Sleep(attente * time.Second)

	// Result
	res, err := getResult(url2, nomscrutin)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("Le gagnant est :", res["winner"])
		fmt.Println("Le classement est :", res["ranking"])
	}
	fmt.Println("\nAppuyez sur Entrée pour passer à la prochaine méthode de vote")
	fmt.Scanln()

	// ************************** BORDA **************************
	fmt.Println("\n######## BORDA ########")

	rule = "borda"
	deadline_t = time.Now().Add(time.Second * attente)
	deadline = deadline_t.Format(time.RFC3339)

	nomscrutin, err = AddBallot(url2, rule, deadline, voters_ids, alts, tiebreak)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Printf("> %s créé [méthode %s, %d alts, deadline %s, tie-break %v]\n", nomscrutin, rule, alts, deadline, tiebreak)
	}

	// Vote
	clAgts = Vote(url2, nomscrutin, n, nb_alts)
	if clAgts == nil {
		fmt.Println("Error")
		return
	}

	fmt.Printf("\n -- Attente de %ds pour la fin de la deadline -- \n", attente)
	time.Sleep(attente * time.Second)

	// Result
	res, err = getResult(url2, nomscrutin)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("Le gagnant est :", res["winner"])
		fmt.Println("Le classement est :", res["ranking"])
	}
	fmt.Println("\nAppuyez sur Entrée pour passer à la prochaine méthode de vote")
	fmt.Scanln()

	// ************************** APPROVAL **************************
	fmt.Println("######## APPROVAL ########")

	rule = "approval"
	deadline_t = time.Now().Add(time.Second * attente)
	deadline = deadline_t.Format(time.RFC3339)

	nomscrutin, err = AddBallot(url2, rule, deadline, voters_ids, alts, tiebreak)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Printf("> %s créé [méthode %s, %d alts, deadline %s, tie-break %v]\n", nomscrutin, rule, alts, deadline, tiebreak)
	}

	// Vote
	clAgts = VoteApproval(url2, nomscrutin, n, nb_alts)
	if clAgts == nil {
		fmt.Println("Error")
		return
	}

	fmt.Printf("\n -- Attente de %ds pour la fin de la deadline -- \n", attente)
	time.Sleep(attente * time.Second)

	// Result
	res, err = getResult(url2, nomscrutin)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("Le gagnant est :", res["winner"])
	}
	fmt.Println("\nAppuyez sur Entrée pour passer à la prochaine méthode de vote")
	fmt.Scanln()

	// ************************** CONDORCET **************************
	fmt.Println("\n######## CONDORCET ########")

	rule = "condorcet"
	deadline_t = time.Now().Add(time.Second * attente)
	deadline = deadline_t.Format(time.RFC3339)

	nomscrutin, err = AddBallot(url2, rule, deadline, voters_ids, alts, tiebreak)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Printf("> %s créé [méthode %s, %d alts, deadline %s, tie-break %v]\n", nomscrutin, rule, alts, deadline, tiebreak)
	}

	// Vote
	clAgts = Vote(url2, nomscrutin, n, nb_alts)
	if clAgts == nil {
		fmt.Println("Error")
		return
	}

	fmt.Printf("\n -- Attente de %ds pour la fin de la deadline -- \n", attente)
	time.Sleep(attente * time.Second)

	// Result
	res, err = getResult(url2, nomscrutin)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		if res["winner"] == nil {
			fmt.Println("Pas de vainqueur de Condorcet.")
		} else {
			fmt.Println("Le gagnant est :", res["winner"])
		}

	}
	fmt.Println("\nAppuyez sur Entrée pour passer à la prochaine méthode de vote")
	fmt.Scanln()

	// ************************** COPELAND **************************
	fmt.Println("\n######## COPELAND ########")

	rule = "copeland"
	deadline_t = time.Now().Add(time.Second * attente)
	deadline = deadline_t.Format(time.RFC3339)

	nomscrutin, err = AddBallot(url2, rule, deadline, voters_ids, alts, tiebreak)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Printf("> %s créé [méthode %s, %d alts, deadline %s, tie-break %v]\n", nomscrutin, rule, alts, deadline, tiebreak)
	}

	// Vote
	clAgts = Vote(url2, nomscrutin, n, nb_alts)
	if clAgts == nil {
		fmt.Println("Error")
		return
	}

	fmt.Printf("\n -- Attente de %ds pour la fin de la deadline -- \n", attente)
	time.Sleep(attente * time.Second)

	// Result
	res, err = getResult(url2, nomscrutin)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("Le gagnant est :", res["winner"])
		fmt.Println("Le classement est :", res["ranking"])
	}
	fmt.Println("\nAppuyez sur Entrée pour passer à la prochaine méthode de vote")
	fmt.Scanln()

	// ************************** STV **************************
	fmt.Println("######## STV ########")

	rule = "stv"
	deadline_t = time.Now().Add(time.Second * attente)
	deadline = deadline_t.Format(time.RFC3339)

	nomscrutin, err = AddBallot(url2, rule, deadline, voters_ids, alts, tiebreak)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Printf("> %s créé [méthode %s, %d alts, deadline %s, tie-break %v]\n", nomscrutin, rule, alts, deadline, tiebreak)
	}

	// Vote
	clAgts = Vote(url2, nomscrutin, n, nb_alts)
	if clAgts == nil {
		fmt.Println("Error")
		return
	}

	fmt.Printf("\n -- Attente de %ds pour la fin de la deadline -- \n", attente)
	time.Sleep(attente * time.Second)

	// Result
	res, err = getResult(url2, nomscrutin)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("Le gagnant est :", res["winner"])
		fmt.Println("Le classement est :", res["ranking"])
	}

}
