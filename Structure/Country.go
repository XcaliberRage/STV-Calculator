//Package Structure implements the type def and related functions for All Objects
//Country.go implements the type def and related functions for Countries
package Structure

type Country struct {
	Name       string
	Regions    []Region
	Electorate int
	ValidVotes int
	SeatsNum   int
	CtSupers   int
}

func (a *Country) MakeCountry(name string) {

	a.Name = name

	region_names := Sc_info.findUnique("region", Reference{"country", a.Name})
	var regions []Region

	for _, region_name := range region_names {
		r := Region{}
		r.MakeRegion(region_name)
		regions = append(regions, r)
		a.CtSupers += r.CtSupers
	}
	a.Regions = regions

	a.SumElectorate()
	a.SumVotes()
	a.SumSeats()

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
