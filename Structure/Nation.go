// Nation.go is a simple struct that ties all data together
package Structure

import "fmt"

type Nation struct {
	Countries []Country
	Parties   []Party
}

// Takes both csv data structures and relates that into the overall data structure
func (a *Nation) NewNation(sc_info CSVData, mp_info CSVData) {

	// Populate the geography
	fmt.Println("Searching Countries")
	country_names := sc_info.findUnique("country", Reference{"", ""})
	fmt.Println(country_names)
	var countries []Country

	for _, country_name := range country_names {

		c := Country{}
		c.MakeCountry(country_name, &sc_info, &mp_info)
		countries = append(countries, c)
	}

	a.Countries = countries

	fmt.Printf("\n\n")

	// Populate the Politics
	fmt.Println("Searching Parties")
	party_names := mp_info.findUnique("party_name", Reference{"", ""})

	var parties []Party

	for _, party_name := range party_names {
		p := Party{}
		p.MakeParty(party_name, &mp_info)
		parties = append(parties, p)
	}
	a.Parties = parties
}
