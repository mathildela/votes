package comsoc

type RequestNewBallot struct {
	Rule      string        `json:"rule"`
	Deadline  string        `json:"deadline"`
	Voter_ids []string      `json:"voter-ids"`
	Alts      int           `json:"#alts"`
	TieBreak  []Alternative `json:"tie-break"`
}

type ResponseNewBallot struct {
	Ballot_id string `json:"ballot-id"`
}

type RequestVote struct {
	Agent_id  string        `json:"agent-id"`
	Ballot_id string        `json:"ballot-id"`
	Prefs     []Alternative `json:"prefs"`
	Options   []int         `json:"options"`
}

type RequestResult struct {
	Ballot_id string `json:"ballot-id"`
}

type ResponseResult struct {
	Winner  Alternative   `json:"winner"`
	Ranking []Alternative `json:"ranking"`
}
