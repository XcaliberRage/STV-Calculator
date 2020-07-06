package main

import (
	"fmt"

	"github.com/XcaliberRage/STV-Calculator/Structure"
)

func main() {

	Structure.Sc_info.ReadNew("cons_info_test.csv")
	fmt.Printf("Number of rows = %d\n", Structure.Sc_info.Length)

	Structure.Mp_info.ReadNew("votes_by_mp_test.csv")
	fmt.Printf("Number of rows = %d\n", Structure.Mp_info.Length)

	// A more practical system would only do this per Super instead of per nation
	Structure.Gb.NewNation()

	Menu(true)

}

func Menu(wipe bool) {
	menu := Structure.Menu{}
	menu.NewReader()
	var action string

	for action != "exit" {
		action = menu.MainMenu(wipe)
		fmt.Println(action)

		switch action {
		case "1":
			action, wipe = menu.ReviewNation(&Structure.Gb)
			break
		case "3":
			Structure.Gb.RunElection()
			wipe = false
			break
		case "4":
			Structure.Gb.GiveCandidates()
			wipe = false
		default:
			wipe = true
		}
	}

	menu.WipeScreen()
	fmt.Println("GOODBYE!")
}
