// Module for saving and loading state
package main

import (
	"fmt"
	"os"
	//"path"
	"io/ioutil"
	//"path/filepath"
	"encoding/json"
	//"encoding/base64"
	//"gopkg.in/yaml.v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//var (
//	THIS_PATH string
//	STATE_PATH string
//)

// State is a struct for representing search state between runs
type State struct {
	Hamming    int	   `json:"hamming"`
	Value      []byte  `json:"value"`
	LastGuess  []byte  `json:"last_guess"`
}

// Save persists the State to the specified path
func (s *State) save(path string) {
	//data, err1 := yaml.Marshal(s)
	data, err1 := json.Marshal(s)
	check(err1)
	err2 := ioutil.WriteFile(path, data, 0644)
	check(err2)
}

// Load reads the state from the specified path to the struct
func (s *State) load(path string) {
	data, err1 := ioutil.ReadFile(path)
	check(err1)
	//err2 := yaml.Unmarshal(data, &s)
	err2 := json.Unmarshal(data, &s)
	check(err2)
}

//func (s *State) MarshalJSON() ([]byte, error) {
//	type Alias State
//	return json.Marshal(&struct {
//		Value     string `json:"value"`
//		LastGuess string `json:"last_guess"`
//		*Alias
//	}{
//		Value:     base64.StdEncoding.EncodeToString(s.Value),
//		LastGuess: base64.StdEncoding.EncodeToString(s.LastGuess),
//		Alias:     (*Alias)(s),
//	})
//}
//
//func (s *State) UnmarshalJSON(data []byte) error {
//	type Alias State
//	aux := &struct {
//		Value      string  `json:"value"`
//		LastGuess  string  `json:"last_guess"`
//		*Alias
//	} {
//		Alias: (*Alias)(s),
//	}
//
//	if err := json.Unmarshal(data, &aux); err != nil {
//		return err
//	}
//
//	value, err := base64.StdEncoding.DecodeString(aux.Value)
//	check(err)
//	s.Value = value
//
//	last_guess, err := base64.StdEncoding.DecodeString(aux.LastGuess)
//	check(err)
//	s.LastGuess = last_guess
//
//	return nil
//}


func main() {
	//THIS_PATH, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//check(err)
	//STATE_PATH = path.Join(THIS_PATH, "state.yml")


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