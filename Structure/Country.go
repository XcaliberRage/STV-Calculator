//Package Structure implements the type def and related functions for All Objects
//Country.go implements the type def and related functions for Countries
package Structure

import "fmt"

type Country struct {
	Name       string
	Regions    []Region
	Electorate int
	ValidVotes int
	SeatsNum   int
}

func (a *Country) MakeCountry(name string, sc_info CSVData, mp_info CSVData) {

	a.Name = name
	fmt.Println("----")
	fmt.Println("Making " + a.Name + ":")

	region_names := sc_info.findUnique("region", Reference{"country", a.Name})

	for _, region_name := range region_names {
		r := Region{}
		r.MakeRegion(region_name, sc_info, mp_info)
	}

	a.SumElectorate()
	a.SumVotes()

	a.SeatsNum = len(region_names)
	fmt.Println("Seats: ", a.SeatsNum)

}

// Sum Electorate across all regions
func (a *Country) SumElectorate() {

	for _, region := range a.Regions {
		a.Electorate += region.Electorate
	}

}

// Sum Valid Votes across all Regions
func (a *Country) SumVotes() {

	for _, region := range a.Regions {
		a.ValidVotes += region.ValidVotes
	}

}
