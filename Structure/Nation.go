// Nation.go is a simple struct that ties all data together
package Structure

import "fmt"

var Gb Nation

type Nation struct {
	Countries []Country
	Parties   []Party
}

// Takes both csv data structures and relates that into the overall data structure
func (a *Nation) NewNation() {

	// Populate the geography
	fmt.Println("Searching Countries")
	country_names := Sc_info.findUnique("country", Reference{"", ""})
	countries := make([]Country, len(country_names))

	for index, country_name := range country_names {

		countries[index] = Country{}
		countries[index].MakeCountry(country_name)

	}

	a.Countries = countries

	fmt.Println("Countries populated")

	// Populate the Politics
	fmt.Println("Searching Parties")
	party_names := Mp_info.findUnique("party_name", Reference{"", ""})

	parties := make([]Party, len(party_names))

	for index, party_name := range party_names {
		p := Party{}
		p.MakeParty(party_name)
		parties[index] = p
	}
	a.Parties = parties
	fmt.Println("Parties populated")

	// With the Candidates and Geography prepared, the Candidates need to be attirbuted to a Super Constituency

	// For each SC
	for _, country := range a.Countries {
		for _, region := range country.Regions {
			for _, sc := range region.Supers {
				sc.FindCandidates()
				sc.FindWinners()
			}
		}
	}

	fmt.Println("Candidates assigned")
}
