package govote

import (
	"fmt"
	//"strings"
	"testing"
)

type PlurarityVoteTest struct {
	Ballots map[string][]byte
}

func newPlurarityVoteTest() (*PlurarityVoteTest, *PlurarityVote) {
	pt := new(PlurarityVoteTest)
	// candidates := "P,R,S,T"

	// pt.Candidates = strings.Split(candidates, ",")

	pt.Ballots = map[string][]byte{
		"P,R,S,T": []byte(`{"num_votes":130,"vote":"P,R,S,T"}`),
		"T,R,S,P": []byte(`{"num_votes":120,"vote":"T,R,S,P"}`),
		"T,R,P,S": []byte(`{"num_votes":100,"vote":"T,R,P,S"}`),
		"S,R,P,T": []byte(`{"num_votes":150,"vote":"S,R,P,T"}`),
	}

	return pt, NewPlurarityVote()
}

//helper function
func (pt *PlurarityVoteTest) AddBallot(p *PlurarityVote, bID ...string) error {
	for _, b := range bID {
		err := p.AddBallot(pt.Ballots[b])
		if err != nil {
			return err
		}
	}
	return nil
}

/*------------------
    ==Tests==
--------------------*/

func TestAddBallot(t *testing.T) {
	pt, p := newPlurarityVoteTest()
	//bolot id from above
	err := pt.AddBallot(p, "")
	if err != nil {
		t.Errorf("Failed to add a ballot, recieved: %s!", err.Error())
	}
}

//test voting
func TestCastVote(t *testing.T) {
	pt, p := newPlurarityVoteTest()

	//get all ballots
	//ballots := make([]string, 0, len(pt.Ballots))
	// for _, keys := range pt.Ballots {
	//  ballots := keys
	// }

	//add all ballots
	pt.AddBallot(p, "P,R,S,T", "T,R,S,P", "T,R,P,S", "S,R,P,T")

	fmt.Println(p.CastVote())
	// expected := []string{"", "", "", ""}, []int{220, 150, 130, 0}, nil
	// if got != expected {
	//  t.Errorf("Expected: %s, got: %s", expected, got)
	// }
}
