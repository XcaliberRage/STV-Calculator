//Package Structure implements the type def and related functions for All Objects
//SuperConstituency.go implements the type def and related functions for Super Constituencies
package Structure

// Candidates points towards a candidate that exists in a Party
type SuperConstituency struct {
	Name       string
	Minors     []MinorConstituency
	Electorate int
	ValidVotes int
	Turnout    float64
	Candidates []*Candidate
	SeatsNum   int
	Seats      []Seat
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

	for _, minor := range a.Minors {
		for i := 0; i < num_parties; i++ {
			party_size := len(Gb.Parties[i].Members)
			for l := 0; l < party_size; l++ {
				if Gb.Parties[i].Members[l].StoodIn == minor.Name {
					a.Candidates = append(a.Candidates, &Gb.Parties[i].Members[l])
				}
			}
		}
	}

}

// Finds which standing candidate actually won the seat in each minor constituency
func (a *SuperConstituency) FindWinners() {

	num_candidates := len(a.Candidates)

	for _, minor := range a.Minors {
		var candidates []*Candidate = nil
		for i := 0; i < num_candidates; i++ {
			if a.Candidates[i].StoodIn == minor.Name {
				candidates = append(candidates, a.Candidates[i])
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
