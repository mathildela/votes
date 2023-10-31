package ballotagent

import (
	"errors"
	"time"

	"gitlab.utc.fr/langemat/ia04/comsoc"
)

type Ballot struct {
	Rule      string
	Deadline  time.Time
	Voter_ids []string
	Alts      int
	Tiebreak  []comsoc.Alternative
	Prof      comsoc.Profile
	Options   [][]int
	A_vote    []string
}

func (rsa *RestServerAgent) NewBallot(ballot_id string, rule string, deadline time.Time, alts int, voter_ids []string, tiebreak []comsoc.Alternative) error {
	_, ok := rsa.ballot_list[ballot_id]
	if ok {
		return errors.New("Ballot already exists")
	} else {
		var p comsoc.Profile = make(comsoc.Profile, 0)
		var a_vote []string = make([]string, 0)
		var options [][]int = make([][]int, 0)
		var ballot Ballot = Ballot{
			Rule:      rule,
			Deadline:  deadline,
			Voter_ids: voter_ids,
			Alts:      alts,
			Tiebreak:  tiebreak,
			Prof:      p,
			Options:   options,
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
	} else if len(prefs) > rsa.ballot_list[ballot_id].Alts {
		return errors.New("too much values in prefs")
	} else {
		for i := 1; i <= rsa.ballot_list[ballot_id].Alts; i++ {
			if !comsoc.Contains(prefs, comsoc.Alternative(i)) {
				return errors.New("missing value(s) in prefs")
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

func GetOptionsApproval(options [][]int) []int {
	var res []int = make([]int, 0)
	for _, val := range options {
		res = append(res, val[0])
	}
	return res
}
func (rsa *RestServerAgent) GetBallot(ballot_id string) Ballot {
	return rsa.ballot_list[ballot_id]
}

func DeadlineExpired(t time.Time) bool {
	return time.Now().After(t)
}

func EarlierThanDeadline(t time.Time) bool {
	return time.Now().Before(t)
}
