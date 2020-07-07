//Package Structure implements the type def and related functions for All Objects
//Region.go implements the type def and related functions for Regions
package Structure

type Region struct {
	Name       string
	Supers     []SuperConstituency
	Electorate int
	ValidVotes int
	SeatsNum   int
	CtSupers   int
}

// Assigns values to each region based on CSV data
func (a *Region) MakeRegion(name string) {

	a.Name = name

	sc_names := Sc_info.findUnique("super_con", Reference{"region", a.Name})
	var supers []SuperConstituency

	for _, sc_name := range sc_names {
		sc := SuperConstituency{}
		sc.MakeSuperConstituency(sc_name)
		supers = append(supers, sc)
	}
	a.Supers = supers

	a.CtSupers = len(a.Supers)

	a.SumElectorate()
	a.SumVotes()
	a.SumSeats()

}

// Sum Electorate across all regions
func (a *Region) SumElectorate() {

	for _, super := range a.Supers {
		a.Electorate += super.Electorate
	}

}

// Sum Valid Votes across all regions
func (a *Region) SumVotes() {

	for _, super := range a.Supers {
		a.ValidVotes += super.ValidVotes
	}

}

func (a *Region) SumSeats() {

	for _, super := range a.Supers {
		a.SeatsNum += super.SeatsNum
	}
}
