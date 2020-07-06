// Nation.go is a simple struct that ties all data together
package Structure

import (
	"fmt"
	"strings"
)

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
	for country, _ := range a.Countries {
		for region, _ := range a.Countries[country].Regions {
			for sc, _ := range a.Countries[country].Regions[region].Supers {
				a.Countries[country].Regions[region].Supers[sc].FindCandidates()
				//a.Countries[country].Regions[region].Supers[sc].FindWinners()
				// This function replaces a function that simply reads a CSV table of all ballots
				a.Countries[country].Regions[region].Supers[sc].MakeBallots()
			}
		}
	}
	fmt.Println("Ballots made")
	fmt.Println("Candidates assigned")
}

func (a *Nation) RunElection() {

	for country := range a.Countries {
		for region, _ := range a.Countries[country].Regions {
			for super, _ := range a.Countries[country].Regions[region].Supers {
				a.Countries[country].Regions[region].Supers[super].NewLocalElection()
			}
		}
	}

}

// Prints all candidates
func (a *Nation) GiveCandidates() {

	for party, _ := range a.Parties {
		fmt.Printf("%s [%s] --------\n", strings.ToUpper(a.Parties[party].Name), a.Parties[party].Brev)
		for member, _ := range a.Parties[party].Members {
			fmt.Printf("%d	:	%s ->	votes %d [%.0f]\n", a.Parties[party].Members[member].ID, a.Parties[party].Members[member].Sname, a.Parties[party].Members[member].Votes, a.Parties[party].Members[member].LiveVotes)
		}

	}

}
