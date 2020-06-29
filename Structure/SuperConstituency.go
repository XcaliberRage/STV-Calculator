//Package Structure implements the type def and related functions for All Objects
//SuperConstituency.go implements the type def and related functions for Super Constituencies
package Structure

type SuperConstituency struct {
	Name       string
	minors     []MinorConstituency
	Electorate int
	ValidVotes int
	candidates []Candidate
	SeatsNum   int
	seats      []Seat
}
