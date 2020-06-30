//Package Structure implements the type def and related functions for All Objects
//SuperConstituency.go implements the type def and related functions for Super Constituencies
package Structure

import "fmt"

type SuperConstituency struct {
	Name       string
	Minors     []MinorConstituency
	Electorate int
	ValidVotes int
	Turnout    float64
	Candidates []Candidate
	SeatsNum   int
	Seats      []Seat
}

// Assigns values to each Super Constituency based on CSV data
func (a *SuperConstituency) MakeSuperConstituency(name string, sc_info *CSVData, mp_info *CSVData) {

	a.Name = name
	fmt.Println("		----")
	fmt.Println("		Making Super " + a.Name + ":")

	minor_names := sc_info.findUnique("original_con", Reference{"super_con", a.Name})
	var minors []MinorConstituency

	for _, minor_name := range minor_names {
		m := MinorConstituency{}
		m.MakeMinorConstituency(minor_name, sc_info, mp_info)
		minors = append(minors, m)
	}
	a.Minors = minors

	a.SumElectorate()
	a.SumVotes()
	a.Turnout = (float64(a.ValidVotes) / float64(a.Electorate)) * 100

	a.SeatsNum = len(minor_names)
	fmt.Printf("		SC Seats: %d\n", a.SeatsNum)
	fmt.Printf("		SC Electorate: %d\n", a.Electorate)
	fmt.Printf("		SC Valid Votes: %d\n", a.ValidVotes)
	fmt.Printf("		SC Turnout: %.2f%%\n", a.Turnout)
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
