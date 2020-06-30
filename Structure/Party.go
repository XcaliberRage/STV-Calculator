//Package Structure implements the type def and related functions for All Objects
//Party.go implements the type def and related functions for Super Constituencies
package Structure

import "fmt"

type Party struct {
	Name    string
	Brev    string
	Members []Candidate
	Votes   int
}

func (a *Party) MakeParty(name string, mp_info *CSVData) {

	a.Name = name

	// Iterate over each row
	for _, row := range mp_info.Rows {

		// If the row does not concern this party, skip it
		if row.Cols["party_name"].Data != a.Name {
			continue
		}

		// Mark Brev if it's not set
		if a.Brev == "" {
			a.Brev = row.Cols["party_abbreviation"].Data
		}

		mp := Candidate{}
		mp.MakeNewCandidate(&row, mp_info)
		a.Members = append(a.Members, mp)
	}

	fmt.Println(a.Name + "[" + a.Brev + "]" + ":")
	for _, v := range a.Members {
		fmt.Println("	" + v.Sname + ", " + v.Fname + ": " + v.StoodIn)
	}

}
