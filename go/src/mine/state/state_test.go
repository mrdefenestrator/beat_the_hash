package state

import (
	"os"
	"fmt"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Print("HELP!")
		os.Exit(-1)
	}

	state_path := args[1]
	new_state_path := args[2]

	s := State{}
	s.load(state_path)

	fmt.Print(s)

	s.save(new_state_path)
}
