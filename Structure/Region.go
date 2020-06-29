//Package Structure implements the type def and related functions for All Objects
//Region.go implements the type def and related functions for Regions
package Structure

type Region struct {
	Name       string
	supers     []SuperConstituency
	Electorate int
	ValidVotes int
	SeatsNum   int
}

//TODO: MakeRegion
