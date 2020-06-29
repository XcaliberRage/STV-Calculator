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
	Cols []Col
}

type Col struct {
	Header string
	Data   string
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

	row := Row{}

	for i, col := range line {
		data := Col{a.Headers[i], col}
		row.Cols = append(row.Cols, data)
	}
	return row
}

func (a *CSVData) findUnique(key string, ref Reference) []string {
	var results []string
	var has_ref = false
	var header_index int

	if ref.Value != "" {
		has_ref = true
		for i, v := range a.Headers {
			if v == ref.Header {
				header_index = i
			}
		}
	}

	for _, row := range a.Rows {

		if has_ref {
			if row.Cols[header_index].Data != ref.Value {
				continue
			}
		}

		for _, col := range row.Cols {

			if col.Header != key {
				continue
			}
			dupe := false

			for _, v := range results {
				if col.Data == v {
					dupe = true
					break
				}
			}

			if dupe {
				continue
			}
			results = append(results, col.Data)
		}
	}
	return results
}
