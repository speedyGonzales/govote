package govote

import (
	"encoding/json"
	"errors"
	"sort"
	"strings"
)

//implementing plurarity voting system
type PairWiseVote struct {
	Candidates []string
	Ballots    map[string]*Ballot
}

type Ballot struct {
	NumVotes int    `json:"num_votes"`
	Vote     string `json:"vote"`
}

func NewPairWiseVote() *PairWiseVote {
	pairWiseVote := new(PairWiseVote)
	pairWiseVote.Candidates = []string{}
	pairWiseVote.Ballots = make(map[string]*Ballot)
	return pairWiseVote
}

//submitting a a new vote
func (p *PairWiseVote) AddBallot(jsonFile []byte) error {
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
func (p *PairWiseVote) FindBallot(vote string) (*Ballot, error) {

	ballot, ok := p.Ballots[vote]
	if ok {
		return ballot, nil
	} else {
		return nil, nil
	}
}

//let's the voting begin
func (p *PairWiseVote) CastVote() ([]string, error) {
	if p.Candidates == nil || p.Ballots == nil {
		return []string{}, []int{}, errors.New("There is a problem with casting the vote")
	}
	winners := p.GetVote()
	return winners, nil
}

//calculating the voting winner
func (p *PairWiseVote) GetVote() ([]string, []int) {

	// scores keyed by candidate name
	//tally := make(map[string]int)

	//the number of candidates
	num := len(p.Candidates)

	//comparison  matrix
	var mat [num][num]int

	//get candidate order
	for _, can1 := range p.Candidates {
		for _can2 := range p.Candidates {
			for key, value := range p.Ballots {
				candidatNames := strings.Split(key, ",")
				if getIndexOf(can, candidatNames) < getIndexOf(can, candidatNames) {
					mat[i][j] += p.Ballots
				}
			}
		}
	}

	// recalculating the matrix
	// mat is the result of pair wise comparison
	// we estimate how manu with what weight
	//difference between each two numbers and calculate the score
	//mat [i] [j] holds how many times i elem is greater then j-th elem
	for _, name := range p.Candidates {
		tally[name] = 0
		for i := range num {
			var sum = 0
			for j := range num {
				if mat[i][j] == mat[j][i] {
					sum += 0.5
					continue
				}
				if mat[i][j] > 0 {
					sum += 1
				}
			}
			tally[name] += sum
		}
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
