// Nation.go is a simple struct that ties all data together
package Structure

import (
	"fmt"
	"strconv"
	"strings"
)

var Gb Nation

type Nation struct {
	Countries []Country
	Parties   []Party
	CtSupers  int
	CtSeats   int
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
		a.CtSupers += countries[index].CtSupers
		a.CtSeats += countries[index].SeatsNum

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
				e := Menu{}
				e.NewReader()
				action := e.WaitInput()
				fmt.Println(action)
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

// Provide information on each Party's performance
func (a *Nation) GetStats() {

	header := "--- ELECTION RESULTS FOR SIMULATED STV ACROSS GREAT BRITAIN AND NORTHERN ISLAND ---"
	tab := "	"

	subline := make([]string, 3)
	subline[0] = tab + strconv.Itoa(a.CtSupers) + "Super Constitnuencies held elections awarding " + strconv.Itoa(a.CtSeats) + " seats and assumed nobody voted outside of their first choice party."
	subline[1] = tab + "It was assumed each voter ordered their preference in order of most votes in real world (always putting their real world choice first however)."
	subline[2] = "\n"

	fmt.Println(header)
	for _, line := range subline {
		fmt.Println(line)
	}

	// For each party display information
	// name			|	Brev|	CtMembers	|	RealVotes	|	VotesElect	|	Seats	|	Seats:VotesElect|	RealSeats:SeatsElect |	RealSeatHolder|	HoldElect |	% Hold Change
	fmt.Println("name			|	Brev|	CtMembers|	RealVotes	|	VotesElect	|	Seats	|	Seats:VotesElect|	RealSeats:SeatsElect |	RealSeatHolder|	HoldElect |	% Hold Change")

	for party, _ := range a.Parties {
		fmt.Printf("%s 			| %s	|	%d |	%d | 	%.0f	|	%d 	|	")
	}
}
