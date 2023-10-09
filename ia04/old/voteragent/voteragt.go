package voteragent

import (
	"fmt"
	"ia04/comsoc"
	"math/rand"
)

type AgentID string

type Agent struct {
	ID    AgentID
	Name  string
	Prefs []comsoc.Alternative
}

type AgentI interface {
	Equal(ag AgentI) bool
	DeepEqual(ag AgentI) bool
	Clone() AgentI
	String() string
	Prefers(a comsoc.Alternative, b comsoc.Alternative) bool
	Start()
}

func CreatePref(nb_candidats int) []comsoc.Alternative {
	alts := make([]comsoc.Alternative, nb_candidats)
	for i := 0; i < nb_candidats; i++ {
		alts[i] = comsoc.Alternative(i + 1)
	}
	rand.Shuffle(len(alts), func(i, j int) { alts[i], alts[j] = alts[j], alts[i] })
	fmt.Println(alts)
	return alts
}

func NewAgent(id AgentID, name string, nb_candidats int) *Agent {
	return &Agent{id, name, CreatePref(nb_candidats)}
}

func (a Agent) Equal(ag Agent) bool {
	return a.ID == ag.ID && a.Name == ag.Name
}

func (a Agent) DeepEqual(ag Agent) bool {
	if !a.Equal(ag) || len(a.Prefs) != len(ag.Prefs) {
		return false
	}
	for idx := range a.Prefs {
		if a.Prefs[idx] != ag.Prefs[idx] {
			return false
		}
	}
	return true
}

func (a Agent) Clone() Agent {
	return *NewAgent(a.ID, a.Name, len(a.Prefs))
}

func (a Agent) String() string {
	return a.Name
}

func rank(alt comsoc.Alternative, prefs []comsoc.Alternative) int {
	for idx := range prefs {
		if prefs[idx] == alt {
			return idx
		}
	}
	return -1
}

func (ag Agent) Prefers(a comsoc.Alternative, b comsoc.Alternative) bool {
	if rank(a, ag.Prefs) == -1 || rank(b, ag.Prefs) == -1 {
		return false
	}
	return rank(a, ag.Prefs) < rank(b, ag.Prefs)
}
