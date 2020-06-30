//Package Structure implements the type def and related functions for All Objects
// MinorConstituency.go is for handling type def and functions for Minor constituencies
package Structure

import (
	"fmt"
	"strconv"
)

type MinorConstituency struct {
	Name           string
	Electorate     int
	ValidVotes     int
	RealSeatHolder Candidate
}

// Assigns values to each Super Constituency based on CSV data
func (a *MinorConstituency) MakeMinorConstituency(name string, sc_info *CSVData, mp_info *CSVData) {

	a.Name = name
	fmt.Println("			Making Minor " + a.Name)

	for _, row := range sc_info.Rows {
		if row.Cols["original_con"].Data == a.Name {
			elec, err := strconv.Atoi(row.Cols["electorate"].Data)
			if err == nil {
				a.Electorate = elec
			}
			valvot, err := strconv.Atoi(row.Cols["valid_votes"].Data)
			if err == nil {
				a.ValidVotes = valvot
			}
		}
	}

	// Find out who actually won the seat

}
