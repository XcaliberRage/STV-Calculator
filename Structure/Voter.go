//Package Structure implements the type def and related functions for All Objects
//Voter.go implements the type def and related functions for Voters
package Structure

type Voter struct {
	RegisteredIn      *SuperConstituency
	Weight            float64
	Ballot            []*Candidate
	IsStillInterested bool
}

func (a *Voter) MakeNew(super *SuperConstituency) {

	a.RegisteredIn = super
	a.Weight = 1.0
	a.IsStillInterested = true

}
