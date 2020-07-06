//Package Structure implements the type def and related functions for All Objects
//Round.go implements the type def and related functions for a Round of elections
package Structure

import (
	"fmt"
)

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

func (a *LocalElection) MakeNew(super *SuperConstituency) {

	runners := make(map[int]*Candidate)
	a.Super = super
	a.SeatsFilled = false
	a.EmptySeats = a.Super.SeatsNum

	a.Voters = a.Super.Voters

	a.DroopQuota = a.Super.OriginalDroopQuota

	for key, _ := range a.Super.Candidates {
		runners[key] = a.Super.Candidates[key]
	}

	a.RunningCandidates = runners

}

// Runs rounds until all seats are filled
func (a *LocalElection) RunRound() {

	rounds := 0
	fmt.Printf("\n%s  - Candidates INITIAL: \n", a.Super.Name)
	for k, _ := range a.RunningCandidates {
		fmt.Printf("%d	->	 %s[%s] 	= %d[%.0f]\n", a.RunningCandidates[k].ID, a.RunningCandidates[k].Sname, a.RunningCandidates[k].Party.Brev, a.RunningCandidates[k].Votes, a.RunningCandidates[k].LiveVotes)
	}
	fmt.Printf("\n\n")

	// When All seats are filled, stop running rounds
	for !a.SeatsFilled {

		rounds++
		// If the number of available seats left is equal to or greater than the number of
		// remaining candidates, then just award all candidates a seat and stop
		if a.EmptySeats >= len(a.RunningCandidates) {
			for candidate, _ := range a.RunningCandidates {
				a.GiveSeat(a.Super.SeatsNum-a.EmptySeats, a.RunningCandidates[candidate])
			}
			fmt.Println("ALL REMAINING CANDIDATES GAIN SEATS")
			continue
		}

		// If there are no more candidates then remaining seats go empty
		if len(a.RunningCandidates) == 0 {
			a.SeatsFilled = true
			fmt.Println("NO MORE CANDIDATES FOR SEATS")
		}

		a.ValidVotes = 0

		for k, _ := range a.RunningCandidates {
			a.RunningCandidates[k].LiveVotes = 0
		}

		// e := Menu{}
		// e.NewReader()
		// action := e.WaitInput()
		// fmt.Println(action)

		// Award each candidates votes as per the voters' weight
		length := len(a.Voters)
		voter := 0
		for voter < length {
			//fmt.Printf("Looking @ (%d): \n", a.Voters[voter].Ballot[0].ID)
			//fmt.Println(a.Voters[voter])
			if a.Voters[voter].IsStillInterested {
				id := a.Voters[voter].Ballot[0].ID
				a.RunningCandidates[id].LiveVotes += a.Voters[voter].Weight
				a.ValidVotes += a.Voters[voter].Weight
				voter++
			}
		}

		// Get a running order of candidates, highest to lowest
		var ordered_candidates []*Candidate
		for k, _ := range a.RunningCandidates {
			ordered_candidates = append(ordered_candidates, a.RunningCandidates[k])
		}
		ordered_candidates = a.OrderCandidatesByLiveVotes(ordered_candidates)

		fmt.Printf("R%d\n", rounds)
		for key, _ := range ordered_candidates {
			fmt.Printf("[%s] %s, %s 	-> %.2f Votes\n", ordered_candidates[key].Party.Brev, ordered_candidates[key].Sname, ordered_candidates[key].Fname, ordered_candidates[key].LiveVotes)
		}

		// Get the DQ for this round
		a.DroopQuota = a.ValidVotes/float64(a.EmptySeats+1) + 1
		fmt.Printf("Valid Votes = %.0f \nDQ = %.2f\n", a.ValidVotes, a.DroopQuota)

		// Get all winners
		var winners []*Candidate
		for k, _ := range ordered_candidates {
			if ordered_candidates[k].LiveVotes >= a.DroopQuota {
				winners = append(winners, ordered_candidates[k])
				continue
			}
			break
		}

		fmt.Printf("\nWinners:\n")
		for key, _ := range winners {
			fmt.Printf("[%s] %s, %s 			-> %.2f Votes\n", winners[key].Party.Brev, winners[key].Sname, winners[key].Fname, winners[key].LiveVotes)
		}

		if len(winners) <= 0 {
			fmt.Println("NO WINNERS")
		}
		fmt.Printf("\n\n")

		// If there are winners do the winner thing
		if len(winners) > 0 {

			a.SetWinners(winners)

			continue
		}

		// Otherwise get all losers
		low_val := ordered_candidates[len(ordered_candidates)-1].LiveVotes
		losers := []*Candidate{}
		for candidate, _ := range ordered_candidates {
			if ordered_candidates[candidate].LiveVotes == low_val {
				losers = append(losers, ordered_candidates[candidate])
				fmt.Printf("Loser -> 	%s	-> %.0f\n", ordered_candidates[candidate].Sname, ordered_candidates[candidate].LiveVotes)
			}
		}

		// And do the loser thing
		a.PurgeLosers(losers)

		fmt.Printf("\n")

	}

	fmt.Print("\n")
	fmt.Printf("ALL POSSIBLE SEATS AWARDED - %s\n", a.Super.Name)

	for count := 0; count < a.Super.SeatsNum; count++ {
		fmt.Printf("%d.	->	", count+1)
		if a.Super.Seats[count].IsEmpty {
			fmt.Printf("Unawarded\n")
			continue
		}
		fmt.Printf("[%s] 	with %d votes 	->	%s, %s\n", a.Super.Seats[count].MP.Party.Brev, a.Super.Seats[count].VotesRequiredToWin, a.Super.Seats[count].MP.Sname, a.Super.Seats[count].MP.Fname)
	}

}

// Eliminates each losing (by ties) candidate from the election
func (a *LocalElection) PurgeLosers(losers []*Candidate) {

	// For each loser:
	//		For each Ballot:
	//			Eliminate this candidate

	for candidate, _ := range losers {
		for ballot, _ := range a.Voters {

			a.RemoveFromBallot(losers[candidate].ID, ballot)

		}
		a.EliminateCandidate(losers[candidate].ID)
	}

}

// Analyses each candidate that exceeded droop quota in order of their victory, assigning seats in said order
func (a *LocalElection) SetWinners(winners []*Candidate) {

	// For each winner:
	// 		Award them a seat (marking the DQ required for them to win said seat)
	//		For each Ballot that awarded them the seat:
	//			Recalculate the Weight
	//			Remove this candidate from their ballot

	// Find the next empty seat
	var i int
	for i = 0; i < a.Super.SeatsNum; i++ {
		if a.Super.Seats[i].IsEmpty {
			break
		}
	}

	for candidate_index, _ := range winners {

		fmt.Printf("Awarding %s [%s] ID: %d Seat: %d\n", winners[candidate_index].Sname, winners[candidate_index].Party.Brev, winners[candidate_index].ID, i+1)
		if a.SeatsFilled {
			return
		}

		a.GiveSeat(i, winners[candidate_index])
		i++

		for ballot, _ := range a.Voters {

			// Ignore Ballots that have pulled out
			if !a.Voters[ballot].IsStillInterested {
				continue
			}

			// Remove this candidate
			a.RemoveFromBallot(winners[candidate_index].ID, ballot)

			// Set the weight
			if a.Voters[ballot].IsStillInterested {
				a.Voters[ballot].Weight = a.Voters[ballot].Weight * ((winners[candidate_index].LiveVotes - a.DroopQuota) / winners[candidate_index].LiveVotes)
			}

		}

		a.EliminateCandidate(winners[candidate_index].ID)
	}

}

func (a *LocalElection) RemoveFromBallot(target_id int, ballot int) {
	index := 0
	target_found := false
	// Look for the ID and identify the candidate_index
	for i, _ := range a.Voters[ballot].Ballot {
		blep := a.Voters[ballot]
		if blep.Ballot[i].ID == target_id {
			index = i
			target_found = true
			break
		}
	}

	if target_found {
		a.SliceOut(index, ballot)
	}
}

// Removes element from the slice
func (a *LocalElection) SliceOut(index int, ballot int) {

	// If it's the last candidate this voter cared about, strike them out of future counts
	length := len(a.Voters[ballot].Ballot)
	if length == 1 {
		fmt.Printf("ELIMINATE BALLOT")
		a.Voters[ballot].IsStillInterested = false
		a.Voters[ballot].Ballot = nil
		return
	}

	// Otherwise, pull the candidate out of the ballot
	i := 0
	found := false
	for i < length {
		t := i
		if i == index {
			found = true
			i++
		}

		if i < length {
			a.Voters[ballot].Ballot[t] = a.Voters[ballot].Ballot[i]
		}
		i++
	}
	if found {
		a.Voters[ballot].Ballot = a.Voters[ballot].Ballot[:length-1]
	}
}

// GiveSeat handles awarding a seat to candidate
// Takes i as an index reference to the given seat and candidate as a pointer to the winning candidate
func (a *LocalElection) GiveSeat(i int, candidate *Candidate) {
	a.Super.Seats[i].MP = candidate
	a.Super.Seats[i].IsEmpty = false

	if candidate.LiveVotes < a.DroopQuota {
		a.Super.Seats[i].VotesRequiredToWin = int(candidate.LiveVotes)
	} else {
		a.Super.Seats[i].VotesRequiredToWin = int(a.DroopQuota)
	}

	a.EmptySeats--
	if a.EmptySeats == 0 {
		a.SeatsFilled = true
	}
}

// Remove the candidate from RunningCandidates
func (a *LocalElection) EliminateCandidate(id int) {
	delete(a.RunningCandidates, id)
}

// Orders running candidates by their current vote tally
// Can do a simple merge sort methinks
func (a *LocalElection) OrderCandidatesByLiveVotes(candidates []*Candidate) []*Candidate {

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
	one = a.OrderCandidatesByLiveVotes(one)
	two = a.OrderCandidatesByLiveVotes(two)

	// Merge them back together
	var new_candidates []*Candidate
	lens := []int{len(one), len(two)}
	// Loop until both arrays have been checked through
	for xy := []int{0, 0}; xy[0] < lens[0] || xy[1] < lens[1]; {

		winner := 1
		if xy[0] >= lens[0] {
			winner = 2
		} else if xy[1] < lens[1] {
			if one[xy[0]].LiveVotes <= two[xy[1]].LiveVotes {
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
