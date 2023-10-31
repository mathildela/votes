package voteragent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gitlab.utc.fr/langemat/ia04/comsoc"
)

type RestClientAgent struct {
	id        string
	url       string
	ballot_id string
	prefs     []comsoc.Alternative
	options   []int
}

func NewRestClientAgent(id string, url string, ballot_id string, prefs []comsoc.Alternative, options []int) *RestClientAgent {
	return &RestClientAgent{id, url, ballot_id, prefs, options}
}

// Quelle réponse du serveur traiter ? Les agents client ne font que voter ?
/*
func (rca *RestClientAgent) treatResponse(r *http.Response) int {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	var resp comsoc.ResponseVote
	json.Unmarshal(buf.Bytes(), &resp)

	return resp.Result
}*/

func (rca *RestClientAgent) doRequest() (err error) {
	req := comsoc.RequestVote{
		Agent_id:  rca.id,
		Ballot_id: rca.ballot_id,
		Prefs:     rca.prefs,
		Options:   rca.options,
	}

	// sérialisation de la requête
	url := rca.url + "/vote"
	data, _ := json.Marshal(req)

	// envoi de la requête
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	// traitement de la réponse
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[%d] %s", resp.StatusCode, resp.Status)
		return err
	}
	//res = rca.treatResponse(resp)
	return nil
}

func (rca *RestClientAgent) Start() {
	log.Printf("démarrage de %s", rca.id)
	err := rca.doRequest()

	if err != nil {
		log.Fatal(rca.id, "error:", err.Error())
	} else {
		log.Printf("[%s] a voté pour %s: %v (si options %v)\n", rca.id, rca.ballot_id, rca.prefs, rca.options)
	}
}
