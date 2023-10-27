package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"ia04/agt/ballotagent"
	"ia04/agt/voteragent"
	"ia04/comsoc"
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
	// Vérifier si nb_max est inférieur à 1
	if nb_max < 1 {
		return nil
	}

	// Créer une carte pour suivre les nombres déjà générés
	generatedNumbers := make(map[comsoc.Alternative]struct{})

	// Créer une slice pour stocker les nombres uniques générés
	uniqueNumbers := []comsoc.Alternative{}

	// Générer des nombres uniques aléatoires jusqu'à ce que la taille soit égale à nb_max
	for len(uniqueNumbers) < nb_max {
		// Générer un nombre aléatoire entre 1 et nb_max
		num := comsoc.Alternative(rand.Intn(nb_max) + 1)

		// Vérifier si le nombre a déjà été généré
		if _, exists := generatedNumbers[num]; !exists {
			// Ajouter le nombre unique à la liste
			uniqueNumbers = append(uniqueNumbers, num)
			// Marquer le nombre comme généré
			generatedNumbers[num] = struct{}{}
		}
	}

	return uniqueNumbers
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
	if resp.StatusCode == http.StatusOK {
		fmt.Println("La requête a réussi!")
	} else {
		fmt.Printf("La requête a échoué avec le code d'état : %d\n", resp.StatusCode)
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

	log.Println("démarrage des clients voters...")
	for i := 1; i <= nb_votants; i++ {
		id := fmt.Sprintf("ag_id%02d", i)
		prefs := generatePrefs(nb_alts)
		//fmt.Println(prefs)
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
	fmt.Scanln()
	return clAgts
}

func VoteApproval(url_server string, nomscrutin string, nb_votants int, nb_alts int) (lAgts []voteragent.RestClientAgent) {
	clAgts := make([]voteragent.RestClientAgent, 0, nb_votants)

	log.Println("démarrage des clients voters...")
	for i := 1; i <= nb_votants; i++ {
		id := fmt.Sprintf("ag_id%02d", i)
		prefs := generatePrefs(nb_alts)
		//fmt.Println(prefs)
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
	fmt.Scanln()
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
		fmt.Println("La requête a réussi!")
	} else {
		fmt.Printf("La requête a échoué avec le code d'état : %d\n", resp.StatusCode)
	}

	//Lire reponse pour la création du scrutin
	defer resp.Body.Close()
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res)
	fmt.Scanln()
	return res, nil
}

func main() {
	const n = 3
	const url1 = ":8080"
	const url2 = "http://localhost:8080"
	const nb_alts = 5

	servAgt := ballotagent.NewRestServerAgent(url1)

	log.Println("démarrage du serveur...")
	go servAgt.Start()

	// ************************** MAJORITE **************************
	// newBallot
	fmt.Println("######## MAJORITY ########")

	rule := "majority"
	deadline := "2023-10-28T12:34:08+02:00"
	voters_ids := []string{"ag_id01", "ag_id02", "ag_id03"}
	alts := nb_alts
	tiebreak := []comsoc.Alternative{4, 2, 3, 5, 1}

	nomscrutin, err := AddBallot(url2, rule, deadline, voters_ids, alts, tiebreak)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Println(nomscrutin)
	}
	fmt.Scanln()

	// Vote
	clAgts := Vote(url2, nomscrutin, 3, nb_alts)
	if clAgts == nil {
		fmt.Println("Error")
		return
	}

	// Result
	res, err := getResult(url2, nomscrutin)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("The winner is :", res["winner"])
		fmt.Println("The ranking is :", res["ranking"])
	}

	// ************************** BORDA **************************
	fmt.Println("######## BORDA ########")

	rule = "borda"
	deadline = "2023-10-28T12:34:08+02:00"
	voters_ids = []string{"ag_id01", "ag_id02", "ag_id03"}
	alts = nb_alts
	tiebreak = []comsoc.Alternative{4, 2, 3, 5, 1}

	nomscrutin, err = AddBallot(url2, rule, deadline, voters_ids, alts, tiebreak)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Println(nomscrutin)
	}
	fmt.Scanln()

	// Vote
	clAgts = Vote(url2, nomscrutin, 3, nb_alts)
	if clAgts == nil {
		fmt.Println("Error")
		return
	}

	// Result
	res, err = getResult(url2, nomscrutin)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("The winner is :", res["winner"])
		fmt.Println("The ranking is :", res["ranking"])
	}

	// ************************** APPROVAL **************************
	fmt.Println("######## APPROVAL ########")

	rule = "approval"
	deadline = "2023-10-28T12:34:08+02:00"
	voters_ids = []string{"ag_id01", "ag_id02", "ag_id03"}
	alts = nb_alts
	tiebreak = []comsoc.Alternative{4, 2, 3, 5, 1}

	nomscrutin, err = AddBallot(url2, rule, deadline, voters_ids, alts, tiebreak)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Println(nomscrutin)
	}
	fmt.Scanln()

	// Vote
	clAgts = VoteApproval(url2, nomscrutin, 3, nb_alts)
	if clAgts == nil {
		fmt.Println("Error")
		return
	}

	// Result
	res, err = getResult(url2, nomscrutin)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("The winner is :", res["winner"])
		fmt.Println("The ranking is :", res["ranking"])
	}

	// ************************** CONDORCET **************************
	fmt.Println("######## CONDORCET ########")

	rule = "condorcet"
	deadline = "2023-10-28T12:34:08+02:00"
	voters_ids = []string{"ag_id01", "ag_id02", "ag_id03"}
	alts = nb_alts
	tiebreak = []comsoc.Alternative{4, 2, 3, 5, 1}

	nomscrutin, err = AddBallot(url2, rule, deadline, voters_ids, alts, tiebreak)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Println(nomscrutin)
	}
	fmt.Scanln()

	// Vote
	clAgts = Vote(url2, nomscrutin, 3, nb_alts)
	if clAgts == nil {
		fmt.Println("Error")
		return
	}

	// Result
	res, err = getResult(url2, nomscrutin)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("The winner is :", res["winner"])
		fmt.Println("The ranking is :", res["ranking"])
	}

	// ************************** COPELAND **************************
	fmt.Println("######## COPELAND ########")

	rule = "copeland"
	deadline = "2023-10-28T12:34:08+02:00"
	voters_ids = []string{"ag_id01", "ag_id02", "ag_id03"}
	alts = nb_alts
	tiebreak = []comsoc.Alternative{4, 2, 3, 5, 1}

	nomscrutin, err = AddBallot(url2, rule, deadline, voters_ids, alts, tiebreak)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Println(nomscrutin)
	}
	fmt.Scanln()

	// Vote
	clAgts = Vote(url2, nomscrutin, 3, nb_alts)
	if clAgts == nil {
		fmt.Println("Error")
		return
	}

	// Result
	res, err = getResult(url2, nomscrutin)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("The winner is :", res["winner"])
		fmt.Println("The ranking is :", res["ranking"])
	}

	// ************************** STV **************************
	fmt.Println("######## STV ########")

	rule = "stv"
	deadline = "2023-10-28T12:34:08+02:00"
	voters_ids = []string{"ag_id01", "ag_id02", "ag_id03"}
	alts = nb_alts
	tiebreak = []comsoc.Alternative{4, 2, 3, 5, 1}

	nomscrutin, err = AddBallot(url2, rule, deadline, voters_ids, alts, tiebreak)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Println(nomscrutin)
	}
	fmt.Scanln()

	// Vote
	clAgts = Vote(url2, nomscrutin, 3, nb_alts)
	if clAgts == nil {
		fmt.Println("Error")
		return
	}

	// Result
	res, err = getResult(url2, nomscrutin)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("The winner is :", res["winner"])
		fmt.Println("The ranking is :", res["ranking"])
	}

}
