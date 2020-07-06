//Package Structure implements the type def and related functions for All Objects
//Seat.go implements the type def and related functions for Super Constituencies
package Structure

type Seat struct {
	MP                 *Candidate
	VotesRequiredToWin int
	Super              *SuperConstituency
	IsEmpty            bool
}

func (a *Seat) MakeNew(super *SuperConstituency) {
	a.Super = super
	a.IsEmpty = true
}
