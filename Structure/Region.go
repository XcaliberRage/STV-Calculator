//Package Structure implements the type def and related functions for All Objects
//Region.go implements the type def and related functions for Regions
package Structure

import "fmt"

type Region struct {
	Name       string
	Supers     []SuperConstituency
	Electorate int
	ValidVotes int
	SeatsNum   int
}

// Assigns values to each region based on CSV data
func (a *Region) MakeRegion(name string, sc_info *CSVData, mp_info *CSVData) {

	a.Name = name
	fmt.Println("	----")
	fmt.Println("	Making Region " + a.Name + ":")

	sc_names := sc_info.findUnique("super_con", Reference{"region", a.Name})
	var supers []SuperConstituency

	for _, sc_name := range sc_names {
		sc := SuperConstituency{}
		sc.MakeSuperConstituency(sc_name, sc_info, mp_info)
		supers = append(supers, sc)
	}
	a.Supers = supers

	a.SumElectorate()
	a.SumVotes()
	a.SumSeats()

	fmt.Println("	SCs: ", len(sc_names))
	fmt.Printf("	Region seats: %d\n", a.SeatsNum)

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
