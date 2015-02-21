package govote

import (
	"encoding/json"
	"errors"
)

type PEliminationVote struct {
	Candidates []string
	Ballots    map[string]*Ballot
}

type Ballot struct {
	NumVotes int    `json:"num_votes"`
	Vote     string `json:"vote"`
}

func NewPEliminationVote() *PEliminationVote {
	pEliminationVote := new(PEliminationVote)
	pEliminationVote.Candidates = []string{}
	pEliminationVote.Ballots = make(map[string]*Ballot)
	return pEliminationVote
}

//submitting a a new vote
func (p *PEliminationVote) AddBallot(jsonFile []byte) error {
	ballot := new(Ballot)

	if err := json.Unmarshal(jsonFile, &ballot); err != nil {
		return errors.New("Unmarshaling ballot error")
	}

	// look up if there is such ballot already
	if _, err := p.FindBallot(ballot.Vote); err == nil {
		return errors.New("There is already such ballot")
	}

	//adding the new ballot
	p.Ballots[ballot.Vote] = ballot
	return nil
}

//we use this function to check for duplicates
//if it is found ballot, we return it
func (p *PEliminationVote) FindBallot(vote string) (*Ballot, error) {

	ballot, ok := p.Ballots[vote]
	if ok {
		return ballot, nil
	} else {
		return nil, err
	}
}

//let's the voting begin
func (p *PEliminationVote) CastVote() (PEliminationVote, error) {
	if p.Candidates == nill || p.Ballots == nill {
		return []string{}, errors.New("There is a problem with casting the vote")
	}
	nextRound, scores := p.GetVote()
	return nextRound, nil
}

//calculating the voting winner
func (p *PEliminationVote) GetVote() PEliminationVote {
	// scores keyed by candidate name
	tally := make(map[string]int)

	// number of candidates
	num := len(p.candidates)

	//
	for key, value := range p.Ballots {

		//get candidate order
		candidatNames = strings.Split(key, ",")

		//only the first pref candidate get votes
		tally[candidatName[0]] += value.NumVotes
	}

	// map with key, value ->score, candidate
	sorted = make(map[int]string)
	scores := []int{}
	winners := []string{}

	//reverse the map and get the scores
	for candidate, score := range tally {
		sorted[score] = candidate
		scores = append(scores, score)
	}

	sort.Ints(scores)

	//get the winning order
	for _, s := range scores {
		winners = append(winners, sorted[s])
	}

	//inizialize the instance for the next round
	nextRound := NewPEliminationVote()

	//remove the loser for this round
	nextRound.Candidates = winners[:num-1]

	return nextRound

}
