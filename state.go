// Module for saving and loading state
package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// State is a struct for representing search state between runs
type State struct {
	Hamming    int	   `json:"hamming"`
	Value      []byte  `json:"value"`
	LastGuess  []byte  `json:"last_guess"`
}

// Save persists the State to the specified path
func (s *State) save(path string) {
	data, err1 := json.Marshal(s)
	check(err1)
	err2 := ioutil.WriteFile(path, data, 0644)
	check(err2)
}

// Load reads the state from the specified path to the struct
func (s *State) load(path string) {
	data, err1 := ioutil.ReadFile(path)
	check(err1)
	err2 := json.Unmarshal(data, &s)
	check(err2)
}


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
