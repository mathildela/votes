package voteragent

type Alternative int
type AgentID string

const nb_candidats = 10

type Agent struct {
	ID    AgentID
	Name  string
	Prefs []Alternative
}

type AgentI interface {
	Equal(ag AgentI) bool
	DeepEqual(ag AgentI) bool
	Clone() AgentI
	String() string
	Prefers(a Alternative, b Alternative) bool
	Start()
}

/*
func NewAgent(id AgentID, name string, nb_candidats int) *Agent {
	return &Point2D{id, name, CreatePref(nb_candidats)}
}*/

func (a Agent) Equal(ag Agent) bool {
	return a.ID == ag.ID && a.Name == ag.Name
}

/*
func (a Agent) Clone() Agent {
	return NewAgent(a.ID, a.Name, nb_candidats)
}*/
