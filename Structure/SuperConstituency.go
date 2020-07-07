//Package Structure implements the type def and related functions for All Objects
//SuperConstituency.go implements the type def and related functions for Super Constituencies
package Structure

import (
	"fmt"
)

// Candidates points towards a candidate that exists in a Party
type SuperConstituency struct {
	Name               string
	Minors             []MinorConstituency
	Electorate         int
	ValidVotes         int
	Turnout            float64
	SeatsNum           int
	Seats              []Seat
	OriginalDroopQuota float64
	Voters             []Voter
	Candidates         map[int]*Candidate
}

// Assigns values to each Super Constituency based on CSV data
func (a *SuperConstituency) MakeSuperConstituency(name string) {

	a.Name = name

	minor_names := Sc_info.findUnique("original_con", Reference{"super_con", a.Name})
	var minors []MinorConstituency

	for _, minor_name := range minor_names {
		m := MinorConstituency{}
		m.MakeMinorConstituency(minor_name)
		minors = append(minors, m)
	}
	a.Minors = minors

	a.SumElectorate()
	a.SumVotes()
	a.Turnout = (float64(a.ValidVotes) / float64(a.Electorate)) * 100

	a.SeatsNum = len(minor_names)

	for i := 0; i < a.SeatsNum; i++ {
		seat := Seat{}
		seat.MakeNew(a)
		a.Seats = append(a.Seats, seat)
	}
	a.OriginalDroopQuota = (float64(a.ValidVotes) / float64(a.SeatsNum+1)) + 1.0
}

// Sum Electorate across all super constituencies
func (a *SuperConstituency) SumElectorate() {

	for _, minor := range a.Minors {
		a.Electorate += minor.Electorate
	}

}

// Sum Valid Votes across all super constituencies
func (a *SuperConstituency) SumVotes() {

	for _, minor := range a.Minors {
		a.ValidVotes += minor.ValidVotes
	}

}

// Finds every Candidate that stood in a minor Constituency that resides in this SC and returns an array of them
func (a *SuperConstituency) FindCandidates() {

	num_parties := len(Gb.Parties)
	candidates := make(map[int]*Candidate)

	for _, minor := range a.Minors {
		for i := 0; i < num_parties; i++ {
			party_size := len(Gb.Parties[i].Members)
			for l := 0; l < party_size; l++ {
				if Gb.Parties[i].Members[l].StoodIn == minor.Name {
					candidates[Gb.Parties[i].Members[l].ID] = &Gb.Parties[i].Members[l]
					Gb.Parties[i].Members[l].StandingIn = a
				}
			}
		}
	}

	a.Candidates = candidates

}

// Finds which standing candidate actually won the seat in each minor constituency
func (a *SuperConstituency) FindWinners() {

	for _, minor := range a.Minors {
		var candidates []*Candidate = nil
		for id, _ := range a.Candidates {
			if a.Candidates[id].StoodIn == minor.Name {
				candidates = append(candidates, a.Candidates[id])
			}
		}

		high := candidates[0]
		for _, running := range candidates {
			if running.Votes > high.Votes {
				high = running
			}
		}

		minor.RealSeatHolder = high
	}
}

func (a *SuperConstituency) NewLocalElection() {

	e := LocalElection{}

	e.MakeNew(a)

	e.RunRound()

}

// Generates a series of fictional ballots based on the real world voting data
func (a *SuperConstituency) MakeBallots() {
	fmt.Printf("Making Ballots for: %s\n", a.Name)
	// Group candidates by party
	parties := make(map[string][]*Candidate)

	for candidate, _ := range a.Candidates {
		parties[a.Candidates[candidate].Party.Brev] = append(parties[a.Candidates[candidate].Party.Brev], a.Candidates[candidate])
	}

	// Order each group most votes to lowest
	for party, _ := range parties {
		parties[party] = OrderCandidatesByVotes(parties[party])
	}

	// Now make Ballots, Presume nobody votes outside their voted party
	for party, _ := range parties {
		for member, _ := range parties[party] {
			for ballot := 0; ballot < parties[party][member].Votes; ballot++ {
				voter := Voter{}
				voter.MakeNew(a)

				// Add this voter's first choice
				voter.Ballot = append(voter.Ballot, parties[party][member].ID)

				// Construct the ballot by going over each party member
				for choice, _ := range parties[party] {
					// Skip their first choice
					if parties[party][choice].ID == voter.Ballot[0] {
						continue
					}

					voter.Ballot = append(voter.Ballot, parties[party][choice].ID)
				}

				a.Voters = append(a.Voters, voter)
			}
		}
	}

}

// I can't believe I've done this. Identical copy of OrderCandidates except looks at .Votes instead of .LiveVotes
func OrderCandidatesByVotes(candidates []*Candidate) []*Candidate {

	length := len(candidates)

	if length <= 1 {
		return candidates
	}

	half := length / 2
	count := 0
	one := make([]*Candidate, half)
	two := make([]*Candidate, length-half)

	// Make two halves
	for k, _ := range candidates {
		if count < half {
			one[k] = candidates[k]
		} else {
			two[k-half] = candidates[k]
		}
		count++
	}

	// Sort each one
	one = OrderCandidatesByVotes(one)
	two = OrderCandidatesByVotes(two)

	// Merge them back together
	var new_candidates []*Candidate
	lens := []int{len(one), len(two)}
	// Loop until both arrays have been checked through
	for xy := []int{0, 0}; xy[0] < lens[0] || xy[1] < lens[1]; {

		winner := 1
		if xy[0] >= lens[0] {
			winner = 2
		} else if xy[1] < lens[1] {
			if one[xy[0]].Votes <= two[xy[1]].Votes {
				winner = 2
			}
		}

		switch winner {
		case 2:
			new_candidates = append(new_candidates, two[xy[1]])
			xy[1]++
		default:
			new_candidates = append(new_candidates, one[xy[0]])
			xy[0]++
		}

	}
	return new_candidates

}
