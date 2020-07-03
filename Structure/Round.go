//Package Structure implements the type def and related functions for All Objects
//Round.go implements the type def and related functions for a Round of elections
package Structure

// Local Election manages the temporary data involved with a Super constituencies local election
type LocalElection struct {
	Super             *SuperConstituency // Points back at the parent Super
	SeatsFilled       bool               // Once this is true no more rounds are needed all remaining candidates instantly lose
	DroopQuota        float64            // Tracks the droop quota round to round
	Voters            []Voter            // Holds ballots with countable votes
	ValidVotes        float64            // Sums the total for all valid votes
	RunningCandidates map[int]*Candidate // Holds remaining candidates in the race
	EmptySeats        int
}

func (a *LocalElection) MakeNew() {

	a.SeatsFilled = false
	a.EmptySeats = a.Super.SeatsNum

	a.Voters = a.Super.Voters

	a.DroopQuota = a.Super.OriginalDroopQuota

	for key, _ := range a.Super.Candidates {
		a.RunningCandidates[key] = a.Super.Candidates[key]
	}

}

// Runs rounds until all seats are filled
func (a *LocalElection) RunRound() {

	for !a.SeatsFilled {

		a.ValidVotes = 0

		// Award each candidates votes as per the voters' weight
		for _, voter := range a.Voters {
			a.RunningCandidates[voter.Ballot[0].ID].LiveVotes += voter.Weight
			a.ValidVotes += voter.Weight
		}

		// Get a running order of candidates, highest to lowest
		var ordered_candidates []*Candidate
		for k, _ := range a.RunningCandidates {
			ordered_candidates = append(ordered_candidates, a.RunningCandidates[k])
		}
		ordered_candidates = a.OrderCandidates(ordered_candidates)

		a.DroopQuota = a.ValidVotes/float64(a.EmptySeats+1) + 1

		// Get all winners
		var winners []*Candidate
		for k, _ := range ordered_candidates {
			if ordered_candidates[k].LiveVotes >= a.DroopQuota {
				winners = append(winners, ordered_candidates[k])
				continue
			}
			break
		}

		// If there are winners do the winner thing
		if len(winners) > 0 {

			continue
		}

		// Otherwise do the loser thing
	}

}

// Orders running candidates by their current vote tally
// Can do a simple merge sort methinks
func (a *LocalElection) OrderCandidates(candidates []*Candidate) []*Candidate {

	length := len(candidates)

	if length <= 1 {
		return candidates
	}

	half := length / 2
	count := 0
	var one []*Candidate
	var two []*Candidate

	// Make two halves
	for k, _ := range candidates {
		if count < half {
			one[k] = candidates[k]
			continue
		} else {
			one[k] = candidates[k]
		}
		count++
	}

	// Sort each one
	one = a.OrderCandidates(one)
	two = a.OrderCandidates(two)

	// Merge them back together
	var new_candidates []*Candidate
	lens := []int{len(one), len(two)}
	// Loop until both arrays have been checked through
	for xy := []int{0, 0}; xy[0] < lens[0] || xy[1] < lens[1]; {

		winner := 1
		if xy[1] >= lens[1] {
			continue
		} else if xy[0] >= lens[0] {
			winner = 2
		} else if one[xy[0]].LiveVotes <= two[xy[1]].LiveVotes {
			winner = 2
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
