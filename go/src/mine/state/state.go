// Package state is for saving and loading miner state
package state

import (
	"io/ioutil"
	"mine/common"

	"gopkg.in/yaml.v2"
)

// State is a struct for representing search state between runs
type State struct {
	Hamming   int    `yaml:"hamming"`
	Value     string `yaml:"value"`
	LastValue string `yaml:"last_value"`
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
