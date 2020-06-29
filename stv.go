package main

import (
	"fmt"
	"stv/Structure"
)

func main() {

	cons_info := Structure.CSVData{}
	cons_info.ReadNew("cons_info.csv")
	fmt.Printf("Number of rows = %d\n", cons_info.Length)

	mp_info := Structure.CSVData{}
	mp_info.ReadNew("votes_by_mp.csv")
	fmt.Printf("Number of rows = %d\n", mp_info.Length)

	gb := Structure.Nation{}
	gb.NewNation(cons_info, mp_info)
}
