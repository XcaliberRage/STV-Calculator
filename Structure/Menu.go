//Package Structure implements the type def and related functions for All Objects
//Menu.go implements the type def and functions related to user interaction
package Structure

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
func (a *Menu) MainMenu() string {

	opts := []string{"1", "2", "3", "exit"}

	a.WipeScreen()

	menu := `
-----WELCOME TO STV CALCULATOR-----

	--MAIN MENU--
1.	Review Nation.
2.	Review Parties.
3.	Simulate STV.
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
