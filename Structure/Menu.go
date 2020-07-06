//Package Structure implements the type def and related functions for All Objects
//Menu.go implements the type def and functions related to user interaction
package Structure

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Menu struct {
	reader bufio.Reader
}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (a *Menu) WipeScreen() {

	value, ok := clear["linux"] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                     //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

// Creates a new reader for the Menu Control
func (a *Menu) NewReader() {
	a.reader = *bufio.NewReader(os.Stdin)
}

// Displays the primary screen
func (a *Menu) MainMenu(should_wipe bool) string {

	opts := []string{"1", "2", "3", "4", "exit"}

	if should_wipe {
		a.WipeScreen()
	}

	menu := `
-----WELCOME TO STV CALCULATOR-----

	--MAIN MENU--
1.	Review Nation.
2.	Review Parties.
3.	Simulate STV.
4.	Print Candidates
exit	Quit

Please select an option:
`

	fmt.Print(menu)
	var input string
	for {
		input = a.WaitInput()

		if !IsIn(input, opts) {
			fmt.Println("Give a valid choice")
			continue
		}

		break
	}

	return input
}

func (a *Menu) WaitInput() string {

	fmt.Print("-> ")
	text, _ := a.reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	return text
}

func IsIn(key string, array []string) bool {
	for _, v := range array {
		if strings.Compare(key, v) == 0 {
			return true
		}
	}
	return false
}

func (a *Menu) ReviewNation(nat *Nation) (string, bool) {

	var action string
	exit := "exit"

	men_ct := len(nat.Countries)
	opts := make([]string, men_ct+2)

	for action != exit {
		a.WipeScreen()
		fmt.Println("The Nation")

		for i := 0; i < men_ct; i++ {
			fmt.Printf("%d.	%s\n", i+1, nat.Countries[i].Name)
			opts = append(opts, strconv.Itoa(i+1))
		}

		opts = append(opts, exit)
		opts = append(opts, "menu")
		fmt.Printf("%s	Quit\n", exit)
		fmt.Println("menu	Main Menu")

		for {
			action = a.WaitInput()

			if !IsIn(action, opts) {
				fmt.Println("Give a valid choice")
				continue
			}

			break
		}

		wipe := false

		if action == exit || action == "menu" {
			wipe = true
			return action, wipe
		}

		a.WipeScreen()
		num, _ := strconv.Atoi(action)
		country := nat.Countries[num-1]
		fmt.Printf("%s:\n", country.Name)
		fmt.Printf("Electorate: 	%d\n", country.Electorate)
		fmt.Printf("Turnout: 	%d\n", country.ValidVotes)
		fmt.Printf("	Regions: \n")
		for k, _ := range country.Regions {
			fmt.Printf("	%s\n", country.Regions[k].Name)
		}

		return "1", wipe

	}
	return action, false
}
