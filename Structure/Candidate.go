//Package Structure implements the type def and related functions for All Objects
//Candidate.go implements the type def and related functions for Super Constituencies
package Structure

import (
	"fmt"
	"strconv"
)

type Candidate struct {
	Fname   string
	Sname   string
	Gender  string
	WasMP   bool
	Votes   int
	StoodIn string
}

func (a *Candidate) MakeNewCandidate(row *Row, mp_info *CSVData) {

	a.Fname = row.Cols["firstname"].Data
	a.Sname = row.Cols["surname"].Data
	a.Gender = row.Cols["gender"].Data
	if a.WasMP = false; row.Cols["former_mp"].Data == "Yes" {
		a.WasMP = true
	}

	votes, err := strconv.Atoi(row.Cols["votes"].Data)
	if err != nil {
		fmt.Println(err)
	}
	a.Votes = votes
	a.StoodIn = row.Cols["constituency_name"].Data

}
