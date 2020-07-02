package main

import (
	"fmt"

	"github.com/XcaliberRage/STV-Calculator/Structure"
)

func main() {

	Structure.Sc_info.ReadNew("cons_info.csv")
	fmt.Printf("Number of rows = %d\n", Structure.Sc_info.Length)

	Structure.Mp_info.ReadNew("votes_by_mp.csv")
	fmt.Printf("Number of rows = %d\n", Structure.Mp_info.Length)

	Structure.Gb.NewNation()

	menu := Structure.Menu{}
	menu.NewReader()
	var action string

	for action != "exit" {
		action = menu.MainMenu()

		switch action {
		default:
			continue
		}
	}

	menu.WipeScreen()
	fmt.Println("GOODBYE!")
}
