package govote

import (
	"encoding/json"
	"errors"
)

//implementing Borda voting system
type BordaVote struct {
	Candidates []string
	Ballots    map[string]*Ballot
}

type Ballot struct {
	NumVotes int    `json:"num_votes"`
	Vote     string `json:"vote"`
}

func NewBordaVote() *BordaVote {
	bordaVote := new(BordaVote)
	bordaVote.Candidates = []string{}
	bordaVote.Ballots = make(map[string]*Ballot)
	return bordaVote
}

//submitting a a new vote
func (b *BordaVote) AddBallot(jsonFile []byte) error {
	ballot := new(Ballot)

	if err := json.Unmarshal(jsonFile, &ballot); err != nil {
		return errors.New("Unmarshaling ballot error")
	}

	// look up if there is such ballot already
	if _, err := b.FindBallot(ballot.Vote); err == nil {
		return errors.New("There is already such ballot")
	}

	//adding the new ballot
	b.Ballots[ballot.Vote] = ballot
	return nil
}

//we use this function to check for duplicates
//if it is found ballot, we return it
func (b *BordaVote) FindBallot(vote string) (*Ballot, error) {

	ballot, ok := b.Ballots[vote]
	if ok {
		return ballot, nil
	} else {
		return nil, err
	}
}

//let's the voting begin
func (b *BordaVote) CastVote() ([]string, []int, error) {
	if b.Candidates == nill || b.Ballots == nill {
		return []string{}, errors.New("There is a problem with casting the vote")
	}
	winners, scores := b.GetVote()
	return winners, scores, nil
}

//calculating the voting winner
func (b *BordaVote) GetVote() ([]string, []int) {
	// scores keyed by candidate name
	tally := make(map[string]int)

	// number of candidates
	num := len(b.candidates)

	//
	for ballot, numBallot := range b.Ballots {

		//get candidate order
		candidatNames = strings.Split(ballot.vote, ",")

		//here the system is more complex, the fist candidate gets
		// num points, netxt num-1 and ect.
		for i := 0; i < num; i++ {
			tally[candidatName[i]] += (num - i) * numBallot
		}

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

	return winners, scores

}
