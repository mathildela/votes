package ballotagent

import (
	"bytes"
	"encoding/json"
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

func NewRestServerAgent(addr string) *RestServerAgent {
	return &RestServerAgent{id: addr, addr: addr}
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
	if len(req.Rule) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("rule empty")
		w.Write([]byte(msg))
		return
	} else if len(req.Deadline) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("deadline empty")
		w.Write([]byte(msg))
		return
	} else if !CheckImplemented(req.Rule) {
		w.WriteHeader(http.StatusNotImplemented)
		msg := fmt.Sprintf("Rule not implemented")
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
	} else if (req.Rule == "majority" || req.Rule == "borda" || req.Rule == "copeland") && !CheckTieBreak(req.TieBreak, req.Alts) {
		// VERIFIER
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("all values from 1 to %d must be in tiebreak", req.Alts)
		w.Write([]byte(msg))
		return
	} else {
		// Conversion Deadline string -> time
		TimeDeadline, err := time.Parse(time.RFC3339, req.Deadline)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			msg := fmt.Sprintf("Error in the deadline format: %s", err)
			w.Write([]byte(msg))
			return
		} else {
			if DeadlineExpired(TimeDeadline) {
				w.WriteHeader(http.StatusBadRequest)
				msg := fmt.Sprintf("Deadline is in the past")
				w.Write([]byte(msg))
				return
			} else {
				// On utilise les reqCount pour déterminer un nom de ballot
				resp.Ballot_id = fmt.Sprintf("scrutin%d", rsa.reqCount)
				err = rsa.NewBallot(resp.Ballot_id, req.Rule, TimeDeadline, req.Alts, req.Voter_ids, req.TieBreak)
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

	if len(req.Ballot_id) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("ballot-id is empty")
		w.Write([]byte(msg))
		return
	} else if !rsa.CheckBallot(req.Ballot_id) {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("ballot-id not found")
		w.Write([]byte(msg))
		return
	} else if len(req.Agent_id) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("agent-id is empty")
		w.Write([]byte(msg))
		return
	} else if !rsa.IdInList(req.Ballot_id, req.Agent_id) {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("agent-id does not exist")
		w.Write([]byte(msg))
		return
	} else if rsa.AVote(req.Ballot_id, req.Agent_id) {
		w.WriteHeader(http.StatusForbidden)
		msg := fmt.Sprintf("Agent already voted")
		w.Write([]byte(msg))
		return
	} else if rsa.CheckPref(req.Prefs, req.Ballot_id) != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("Error : %s", rsa.CheckPref(req.Prefs, req.Ballot_id))
		w.Write([]byte(msg))
		return
	} else if DeadlineExpired(rsa.ballot_list[req.Ballot_id].Deadline) {
		w.WriteHeader(http.StatusServiceUnavailable)
		msg := fmt.Sprintf("Deadline is over")
		w.Write([]byte(msg))
		return
	} else {
		var Ballot_copy Ballot = rsa.ballot_list[req.Ballot_id]
		Ballot_copy.Prof = append(Ballot_copy.Prof, req.Prefs)
		Ballot_copy.A_vote = append(Ballot_copy.A_vote, req.Agent_id)
		Ballot_copy.Options = append(Ballot_copy.Options, req.Options)
		rsa.ballot_list[req.Ballot_id] = Ballot_copy
		w.WriteHeader(http.StatusOK)
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
	if len(req.Ballot_id) == 0 {
		w.WriteHeader(http.StatusNotFound)
		msg := fmt.Sprintf("ballot-id is empty")
		w.Write([]byte(msg))
		return
	} else if !rsa.CheckBallot(req.Ballot_id) {
		w.WriteHeader(http.StatusNotFound)
		msg := fmt.Sprintf("ballot-id not found")
		w.Write([]byte(msg))
		return
	} else if EarlierThanDeadline(rsa.ballot_list[req.Ballot_id].Deadline) {
		w.WriteHeader(http.StatusTooEarly)
		msg := fmt.Sprintf("Deadline is not over")
		w.Write([]byte(msg))
		return
	} else {
		// utiliser les bonnes fonctions ici donc à partir d'un case
		var ballot Ballot = rsa.ballot_list[req.Ballot_id]
		// Cas où personne n'a voté
		if len(ballot.Prof) == 0 {
			w.WriteHeader(http.StatusOK)
			serial, _ := json.Marshal(resp)
			w.Write(serial)
		} else {
			// Cas où au moins une personne a voté
			switch ballot.Rule {
			case "majority":
				fmt.Println("Majority")
				func_winner := comsoc.SCFFactory(comsoc.MajoritySCF, comsoc.TieBreakFactory(ballot.Tiebreak))
				winner, err1 := func_winner(ballot.Prof)
				func_ranking := comsoc.SWFFactory(comsoc.MajoritySWF, comsoc.TieBreakFactory(ballot.Tiebreak))
				ranking, err2 := func_ranking(ballot.Prof)
				if err1 == nil && err2 == nil {
					resp.Winner = winner
					resp.Ranking = ranking
					w.WriteHeader(http.StatusOK)
					serial, _ := json.Marshal(resp)
					w.Write(serial)
				} else {
					if err1 != nil {
						w.WriteHeader(http.StatusBadRequest)
						msg := fmt.Sprintf("Error: %s", err1)
						w.Write([]byte(msg))
					}
					if err2 != nil {
						w.WriteHeader(http.StatusBadRequest)
						msg := fmt.Sprintf("Error: %s", err2)
						w.Write([]byte(msg))
						return
					}
				}
			case "borda":
				fmt.Println("Borda")
				func_winner := comsoc.SCFFactory(comsoc.BordaSCF, comsoc.TieBreakFactory(ballot.Tiebreak))
				winner, err1 := func_winner(ballot.Prof)
				func_ranking := comsoc.SWFFactory(comsoc.BordaSWF, comsoc.TieBreakFactory(ballot.Tiebreak))
				ranking, err2 := func_ranking(ballot.Prof)
				if err1 == nil && err2 == nil {
					resp.Winner = winner
					resp.Ranking = ranking
					w.WriteHeader(http.StatusOK)
					serial, _ := json.Marshal(resp)
					w.Write(serial)
				} else {
					if err1 != nil {
						w.WriteHeader(http.StatusBadRequest)
						msg := fmt.Sprintf("Error: %s", err1)
						w.Write([]byte(msg))
					}
					if err2 != nil {
						w.WriteHeader(http.StatusBadRequest)
						msg := fmt.Sprintf("Error: %s", err2)
						w.Write([]byte(msg))
						return
					}
				}
			case "approval":
				// Besoin des options + pas de gestion de tiebreak dans la fonction
				// Verifier que les options sont cohérentes
				winner, err := comsoc.ApprovalSCF(ballot.Prof, GetOptionsApproval(ballot.Options))
				if err == nil {
					resp.Winner = winner[0]
					resp.Ranking = nil
					w.WriteHeader(http.StatusOK)
					serial, _ := json.Marshal(resp)
					w.Write(serial)
				} else {
					w.WriteHeader(http.StatusBadRequest)
					msg := fmt.Sprintf("Error: %s", err)
					w.Write([]byte(msg))
					return
				}
			case "condorcet":
				fmt.Println("Condorcet")
				winner, err := comsoc.CondorcetWinner(ballot.Prof)
				if err == nil {
					if winner != nil {
						// Cas où on trouve un gagnant de Condorcet
						resp.Winner = winner[0]
						w.WriteHeader(http.StatusOK)
						serial, _ := json.Marshal(resp)
						w.Write(serial)
					}
					// Sinon on ne renvoie rien
				} else {
					w.WriteHeader(http.StatusBadRequest)
					msg := fmt.Sprintf("Error: %s", err)
					w.Write([]byte(msg))
					return
				}
			case "copeland":
				fmt.Println("Copeland")
				func_winner := comsoc.SCFFactory(comsoc.CopelandSCF, comsoc.TieBreakFactory(ballot.Tiebreak))
				winner, err1 := func_winner(ballot.Prof)
				func_ranking := comsoc.SWFFactory(comsoc.CopelandSWF, comsoc.TieBreakFactory(ballot.Tiebreak))
				ranking, err2 := func_ranking(ballot.Prof)
				if err1 == nil && err2 == nil {
					resp.Winner = winner
					resp.Ranking = ranking
					w.WriteHeader(http.StatusOK)
					serial, _ := json.Marshal(resp)
					w.Write(serial)
				} else {
					if err1 != nil {
						fmt.Println("erreur 1")
						w.WriteHeader(http.StatusBadRequest)
						msg := fmt.Sprintf("Error: %s", err1)
						w.Write([]byte(msg))
					}
					if err2 != nil {
						fmt.Println("erreur 1")
						w.WriteHeader(http.StatusBadRequest)
						msg := fmt.Sprintf("Error: %s", err2)
						w.Write([]byte(msg))
						return
					}
				}
			case "stv":
				fmt.Println("STV")
				winner, err1 := comsoc.STV_SCF(ballot.Prof)
				func_ranking := comsoc.SWFFactory(comsoc.STV_SWF, comsoc.TieBreakFactory(ballot.Tiebreak))
				ranking, err2 := func_ranking(ballot.Prof)
				if err1 == nil && err2 == nil {
					resp.Winner = winner[0]
					resp.Ranking = ranking
					w.WriteHeader(http.StatusOK)
					serial, _ := json.Marshal(resp)
					w.Write(serial)
				} else {
					if err1 != nil {
						w.WriteHeader(http.StatusBadRequest)
						msg := fmt.Sprintf("Error: %s", err1)
						w.Write([]byte(msg))
					}
					if err2 != nil {
						w.WriteHeader(http.StatusBadRequest)
						msg := fmt.Sprintf("Error: %s", err2)
						w.Write([]byte(msg))
						return
					}
				}
			default:
				fmt.Println("Unknown Method")
				w.WriteHeader(http.StatusNotImplemented)
				msg := fmt.Sprintf("Rule not implemented")
				w.Write([]byte(msg))
				return
			}
		}
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
