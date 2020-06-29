//Package Structure implements the type def and related functions for All Objects
//Party.go implements the type def and related functions for Super Constituencies
package Structure

type Party struct {
	Name    string
	Brev    string
	members []Candidate
	Votes   int
}
