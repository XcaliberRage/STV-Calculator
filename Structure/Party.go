//Package Structure implements the type def and related functions for All Objects
//Party.go implements the type def and related functions for Super Constituencies
package Structure

import "fmt"

type Party struct {
	Name    string
	Brev    string
	members []Candidate
	Votes   int
}

func (a *Party) MakeParty(name string, mp_info *CSVData) {

	a.Name = name

	// Get the abbreviation of the party_name
	for _, row := range mp_info.Rows {

		if row.Cols["party_name"].Data != a.Name {
			continue
		}

		a.Brev = row.Cols["party_abbreviation"].Data
		fmt.Println(a.Name + "[" + a.Brev + "]")
		break
	}

	// Populate the MPs
	for _, row := range mp_info.Rows {

		if row.Cols["party_name"].Data != a.Name {
			continue
		}
	}

}
