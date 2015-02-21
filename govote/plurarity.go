package govote

import (
	"encoding/json"
	"errors"
	"sort"
	"strings"
)

//implementing plurarity voting system
type PlurarityVote struct {
	Candidates []string
	Ballots    map[string]*Ballot
}

type Ballot struct {
	NumVotes int    `json:"num_votes"`
	Vote     string `json:"vote"`
}

func NewPlurarityVote() *PlurarityVote {
	plurarityVote := new(PlurarityVote)
	plurarityVote.Candidates = []string{}
	plurarityVote.Ballots = make(map[string]*Ballot)
	return plurarityVote
}

//submitting a a new vote
func (p *PlurarityVote) AddBallot(jsonFile []byte) error {
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
func (p *PlurarityVote) FindBallot(vote string) (*Ballot, error) {

	ballot, ok := p.Ballots[vote]
	if ok {
		return ballot, nil
	} else {
		return nil, nil
	}
}

//let's the voting begin
func (p *PlurarityVote) CastVote() ([]string, []int, error) {
	if p.Candidates == nil || p.Ballots == nil {
		return []string{}, []int{}, errors.New("There is a problem with casting the vote")
	}
	winners, scores := p.GetVote()
	return winners, scores, nil
}

//calculating the voting winner
func (p *PlurarityVote) GetVote() ([]string, []int) {
	// scores keyed by candidate name
	tally := make(map[string]int)

	//
	for key, value := range p.Ballots {

		//get candidate order
		candidatNames := strings.Split(key, ",")

		//only the first pref candidate get votes
		tally[candidatNames[0]] += value.NumVotes
	}

	// map with key, value ->score, candidate
	sorted := make(map[int]string)
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

	return winners, scores

}
