//Package Structure implements the type def and related functions for All Objects
//Translator.go implements the type def and functions related to translating a csv file into the data structure
package Structure

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CSVData struct {
	Headers []string
	Rows    []Row
	Length  int
}

type Row struct {
	Cols map[string]Col
}

type Col struct {
	Data string
}

type Reference struct {
	Header string
	Value  string
}

// Reads a csv file and formats the data into convenient info
func (a *CSVData) ReadNew(input string) {
	csvFile, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Successfully Opened CSV file: %s\n", input)
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for index, line := range csvLines {
		// Iterate over each line.

		// For the first line create the Headers
		// Every other line requires a new row
		switch index {
		case 0:
			a.Headers = a.MakeHeaders(line)
		default:
			a.Rows = append(a.Rows, a.MakeRow(line))
		}
	}
	a.Length = len(a.Rows)
}

// Formats the header of a CSVData struct
func (a *CSVData) MakeHeaders(line []string) []string {

	var head []string

	for _, col := range line {
		head = append(head, col)
	}
	return head

}

// Formats a row of a CSVData struct
func (a *CSVData) MakeRow(line []string) Row {

	row_col := make(map[string]Col)

	for i, col := range line {
		row_col[a.Headers[i]] = Col{col}
	}
	row := Row{row_col}
	return row
}

// Takes a string and searches the CSVData for every unique instance of that string returning a string array of each found
// Reference contains optional parameters to filter by a specific common field under a different header
// E.g. findUnique("region", Reference{"country","england"}) will return a list of each unique entry found under the column "region" found but only counting rows that have "england" under the "column" country
func (a *CSVData) findUnique(key string, ref Reference) []string {
	var results []string
	var has_ref = false
	if ref.Header != "" {
		has_ref = true
	}

	for _, row := range a.Rows {

		// Ignore rows that don't have the specific data under the specific column
		if has_ref {
			if row.Cols[ref.Header].Data != ref.Value {
				continue
			}
		}

		// Check it's not a duplicate
		duplicate := false
		for _, v := range results {
			if row.Cols[key].Data == v {
				duplicate = true
				break
			}
		}

		if !duplicate {
			results = append(results, row.Cols[key].Data)
		}

	}
	return results
}
