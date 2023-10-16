package ballotagent

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"ia04/comsoc"
	"log"
	"net/http"
	"sync"
	"time"
)

type RestServerAgent struct {
	sync.Mutex
	id          string
	addr        string
	reqCount    int
	ballot_list map[string]Ballot
}

type Ballot struct {
	Rule      string
	Deadline  string
	Voter_ids []string
	Alts      int
	Tiebreak  []comsoc.Alternative
	Prof      comsoc.Profile
	A_vote    []string
}

func NewRestServerAgent(addr string) *RestServerAgent {
	return &RestServerAgent{id: addr, addr: addr}
}

func (rsa *RestServerAgent) NewBallot(ballot_id string, rule string, deadline string, alts int, voter_ids []string, tiebreak []comsoc.Alternative) error {
	// Vérifie si cd ballot existe deja
	_, ok := rsa.ballot_list[ballot_id]
	if ok {
		return errors.New("Ballot already exists")
	} else {
		var p comsoc.Profile = make(comsoc.Profile, 0)
		var a_vote []string = make([]string, 0)
		var ballot Ballot
		ballot = Ballot{
			Rule:      rule,
			Deadline:  deadline,
			Voter_ids: voter_ids,
			Alts:      alts,
			Tiebreak:  tiebreak,
			Prof:      p,
			A_vote:    a_vote,
		}
		rsa.ballot_list[ballot_id] = ballot
		return nil
	}
}

func CheckImplemented(rule string) bool {
	ReferenceList := []string{"majority", "borda", "approval", "stv", "condorcet", "copeland"}
	for _, value := range ReferenceList {
		if value == rule {
			return true
		}
	}
	return false
}

// A coder : verifie si la deadline est bien ecrite et si elle est bien dans le futur
func CheckDeadline(deadline string) bool {
	return true
}

func CheckTieBreak(tiebreak []comsoc.Alternative, alts int) bool {
	for i := 1; i <= alts; i++ {
		if !comsoc.Contains(tiebreak, comsoc.Alternative(i)) {
			return false
		}
	}
	return true
}

func (rsa *RestServerAgent) AVote(ballot string, agent string) bool {
	for _, val := range rsa.ballot_list[ballot].A_vote {
		if val == agent {
			return true
		}
	}
	return false
}

func (rsa *RestServerAgent) CheckBallot(ballot_id string) bool {
	_, ok := rsa.ballot_list[ballot_id]
	return ok
}

func (rsa *RestServerAgent) CheckPref(prefs []comsoc.Alternative, ballot_id string) error {
	if len(prefs) == 0 {
		return errors.New("prefs is empty")
	} else {
		for i := 1; i <= rsa.ballot_list[ballot_id].Alts; i++ {
			if !comsoc.Contains(prefs, comsoc.Alternative(i)) {
				return errors.New("Missing value(s) in prefs")
			}
		}
	}
	return nil
}

func (rsa *RestServerAgent) IdInList(ballot_id string, agent_id string) bool {
	for _, val := range rsa.ballot_list[ballot_id].Voter_ids {
		if val == agent_id {
			return true
		}
	}
	return false
}

// Test de la méthode
func (rsa *RestServerAgent) checkMethod(method string, w http.ResponseWriter, r *http.Request) bool {
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "method %q not allowed", r.Method)
		return false
	}
	return true
}

func (*RestServerAgent) decodeRequestNewBallot(r *http.Request) (req comsoc.RequestNewBallot, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (*RestServerAgent) decodeRequestVote(r *http.Request) (req comsoc.RequestVote, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

func (*RestServerAgent) decodeRequestResult(r *http.Request) (req comsoc.RequestResult, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	err = json.Unmarshal(buf.Bytes(), &req)
	return
}

// La fonction doNewBallot permet la création d'un Profile
func (rsa *RestServerAgent) doNewBallot(w http.ResponseWriter, r *http.Request) {
	rsa.Lock()
	defer rsa.Unlock()
	rsa.reqCount++

	// vérification de la méthode de la requête
	if !rsa.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := rsa.decodeRequestNewBallot(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	// traitement de la requête
	var resp comsoc.ResponseNewBallot

	// Faire différents tests pour renvoyer les bons messages d'erreurs
	// 1) Vérifier que les strings de rule sont bien dans la liste d'implémentés -> not implemented
	// 2) Vérifier que la deadline est dans le futur -> faire une fonction A FAIRE
	// 3) Vérifier que alts > 0
	// 4) Vérifier que voter-ids n'est pas vide
	// 5) Vérifier que le tie-break n'est pas vide et que sa longueur est égale à alts
	if !CheckImplemented(req.Rule) {
		w.WriteHeader(http.StatusNotImplemented)
		msg := fmt.Sprintf("Rule not implemented")
		w.Write([]byte(msg))
		return
	} else if !CheckDeadline(req.Deadline) {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("deadline is incorrect")
		w.Write([]byte(msg))
		return
	} else if req.Alts <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("alts must be strictly positive")
		w.Write([]byte(msg))
		return
	} else if len(req.Voter_ids) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("voter-ids empty")
		w.Write([]byte(msg))
		return
	} else if len(req.TieBreak) != req.Alts {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("tiebreak list must contain %d values", req.Alts)
		w.Write([]byte(msg))
		return
	} else if !CheckTieBreak(req.TieBreak, req.Alts) {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("all values from 1 to %d must be in tiebreak", req.Alts)
		w.Write([]byte(msg))
		return
	} else {
		// On utilise les reqCount pour déterminer un nom de ballot

		resp.Ballot_id = fmt.Sprintf("scrutin%d", rsa.reqCount)
		err := rsa.NewBallot(resp.Ballot_id, req.Rule, req.Deadline, req.Alts, req.Voter_ids, req.TieBreak)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			msg := fmt.Sprintf("Error: %s", err)
			w.Write([]byte(msg))
			return
		} else {
			w.WriteHeader(http.StatusOK)
			serial, _ := json.Marshal(resp)
			w.Write(serial)
		}
	}
}

// La fonction doVote permet de remplir le profile créé avec la fonction précédente
// Il faut vérifier que le profile existe, que le vote n'a pas déjà été fait et que la deadline
// n'est pas passée
func (rsa *RestServerAgent) doVote(w http.ResponseWriter, r *http.Request) {
	rsa.Lock()
	defer rsa.Unlock()

	// vérification de la méthode de la requête
	if !rsa.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := rsa.decodeRequestVote(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	// Verifier que les valeurs sont cohérentes (vérifications plus poussées à l'avenir)
	if len(req.Agent_id) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("agent-id is empty")
		w.Write([]byte(msg))
		return
	} else if !rsa.IdInList(req.Ballot_id, req.Agent_id) {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("agent-id does not exist")
		w.Write([]byte(msg))
		return
	} else if len(req.Ballot_id) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("ballot-id is empty")
		w.Write([]byte(msg))
		return
	} else if !rsa.CheckBallot(req.Ballot_id) {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("ballot-id not found")
		w.Write([]byte(msg))
		return
	} else if rsa.AVote(req.Ballot_id, req.Agent_id) { //NE MARCHE PAS ENCORE
		w.WriteHeader(http.StatusForbidden)
		msg := fmt.Sprintf("Agent already voted")
		w.Write([]byte(msg))
		return
	} else if rsa.CheckPref(req.Prefs, req.Ballot_id) != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("Error : %s", rsa.CheckPref(req.Prefs, req.Ballot_id))
		w.Write([]byte(msg))
		return
		// CheckDeadline aussi pour verifier si la deadline est dépassée
	} else {
		var Ballot_copy Ballot = rsa.ballot_list[req.Ballot_id]
		Ballot_copy.Prof = append(Ballot_copy.Prof, req.Prefs)
		rsa.ballot_list[req.Ballot_id] = Ballot_copy
		w.WriteHeader(http.StatusOK)
		// A COMPLETER
	}
}

// La fonction doResult permet d'obtenir le résultat à partir du profile créé
func (rsa *RestServerAgent) doResult(w http.ResponseWriter, r *http.Request) {
	rsa.Lock()
	defer rsa.Unlock()

	// vérification de la méthode de la requête
	if !rsa.checkMethod("POST", w, r) {
		return
	}

	// décodage de la requête
	req, err := rsa.decodeRequestResult(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	// traitement de la requête
	var resp comsoc.ResponseResult

	// Verifier que les valeurs sont cohérentes (vérifications plus poussées à l'avenir)
	if len(req.Ballot_id) != 0 {
		// utiliser les bonnes fonctions ici donc à partir d'un case
		resp.Winner = 12
		ranking := []int{1, 2, 3, 4}
		resp.Ranking = ranking // Pas trop compris ce que c'est
		w.WriteHeader(http.StatusOK)
		serial, _ := json.Marshal(resp)
		w.Write(serial)

	} else {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("Error")
		w.Write([]byte(msg))
		return
	}
}

func (rsa *RestServerAgent) Start() {
	// création du multiplexer
	mux := http.NewServeMux()
	mux.HandleFunc("/new_ballot", rsa.doNewBallot)
	mux.HandleFunc("/vote", rsa.doVote)
	mux.HandleFunc("/result", rsa.doResult)

	// création du serveur http
	s := &http.Server{
		Addr:           rsa.addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20}

	// lancement du serveur
	rsa.ballot_list = make(map[string]Ballot)

	log.Println("Listening on", rsa.addr)
	go log.Fatal(s.ListenAndServe())
}
