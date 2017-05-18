// Module for saving and loading state
package state

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"mine/common"
)

// State is a struct for representing search state between runs
type State struct {
	Hamming    int	   `yaml:"hamming"`
	Value      string  `yaml:"value"`
	LastGuess  string  `yaml:"last_guess"`
}

// Save persists the State to the specified path
func (s *State) Save(path string) {
	data, err1 := yaml.Marshal(s)
	common.Check(err1)
	err2 := ioutil.WriteFile(path, data, 0644)
	common.Check(err2)
}

// Load reads the state from the specified path to the struct
func (s *State) Load(path string) {
	data, err1 := ioutil.ReadFile(path)
	common.Check(err1)
	err2 := yaml.Unmarshal(data, &s)
	common.Check(err2)
}
