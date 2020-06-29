//Package Structure implements the type def and related functions for All Objects
// MinorConstituency.go is for handling type def and functions for Minor constituencies
package Structure

type MinorConstituency struct {
	Name           string
	Electorate     int
	ValidVotes     int
	RealSeatHolder Candidate
}
