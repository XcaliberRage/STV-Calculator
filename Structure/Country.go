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

func (a *Country) MakeCountry(name string, sc_info *CSVData, mp_info *CSVData) {

	a.Name = name
	fmt.Println("----")
	fmt.Println("Making Country " + a.Name + ":")

	region_names := sc_info.findUnique("region", Reference{"country", a.Name})
	fmt.Println(region_names)
	var regions []Region

	for _, region_name := range region_names {
		r := Region{}
		r.MakeRegion(region_name, sc_info, mp_info)
		regions = append(regions, r)
	}
	a.Regions = regions

	a.SumElectorate()
	a.SumVotes()
	a.SumSeats()

	fmt.Println("Regions: ", len(region_names))
	fmt.Printf("Total Seats: %d\n", a.SeatsNum)

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

func (a *Country) SumSeats() {
	for _, region := range a.Regions {
		a.SeatsNum += region.SeatsNum
	}
}
