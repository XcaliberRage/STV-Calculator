//Package Structure implements the type def and related functions for All Objects
//SuperConstituency.go implements the type def and related functions for Super Constituencies
package Structure

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
