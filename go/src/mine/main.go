package main

import (
	"os"
	"fmt"
	"mine/common"
	"strconv"
	"mine/state"
	"encoding/hex"
	"unicode/utf8"
)

// TODO: multithreading/processing using goroutines on worker?
// TODO: tests for each of the functions I implement
// TODO: program for posting results to site

var bin_hash []byte

// Worker to mine for hash values
func worker(start int, n_values int, best_hamming int, best_value string, ch chan state.State) {
	var last_value []byte

	for guess_chan := range common.GenGuess(n_values, start) {
		value := []byte(
			common.ListToUnicode(common.ToBase(guess_chan, utf8.MaxRune)))

		hamming := common.CalcHamming(bin_hash, common.HashIt(value))

		if hamming < best_hamming {
			// Maxmimze
			fmt.Println("New Best", hamming, value)
			best_hamming = hamming
			best_value = best_value
		}

		last_value = value
	}


	ch <- state.State{
		Hamming:   best_hamming,
		Value:     best_value,
		LastValue: string(last_value),
	}
}


func str_to_int(value string) int {
	return common.FromBase(common.UnicodeToList(value), utf8.MaxRune)
}

func int_to_str(value int) string {
	return common.ListToUnicode(common.ToBase(value, utf8.MaxRune))
}


func main() {
	usage := "mine.py <n_values> <n_processes> <state_path>\n\n" +
		"n_values     number of values to mine\n" +
		"n_processes  number of processes to use for mining\n" +
		"state_path   path to YAML state file"
	hex_hash := "5b4da95f5fa08280fc9879df44f418c8f9f12ba424b7757de02bbdfbae0" +
		"d4c4fdf9317c80cc5fe04c6429073466cf29706b8c25999ddd2f6540d4475cc977b" +
		"87f4757be023f19b8f4035d7722886b78869826de916a79cf9c94cc79cd4347d24b" +
		"567aa3e2390a573a373a48a5e676640c79cc70197e1c5e7f902fb53ca1858b6"
	var err error
	bin_hash, err = hex.DecodeString(hex_hash)
	common.Check(err)

	args := os.Args
	if len(args) != 4 {
		fmt.Println(usage)
		os.Exit(-1)
	}

	// Get arguments
	n_values, err := strconv.Atoi(os.Args[1])
	common.Check(err)
	n_processes, err := strconv.Atoi(os.Args[2])
	common.Check(err)
	state_path := os.Args[3]

	// Check if state file exists
	var my_state state.State
	if _, err := os.Stat(state_path); err == nil {
		// Load state
		my_state.Load(state_path)
	} else {
		// Generate initial state
		my_state = state.State{Hamming: 1024, Value: "", LastValue: ""}
	}

	// Create channel for results
	ch := make(chan state.State)

	// Divide work & start multiple workers
	start_value := str_to_int(my_state.LastValue)
	for i := 0; i < n_processes; i++ {
		start := (n_values / n_processes) * i + start_value
		go worker(
			start,
			n_values / n_processes,
			my_state.Hamming,
			my_state.Value,
			ch,
		)
	}

	// Determine best result
	n_results := 0
	for result := range ch {
		if result.Hamming < my_state.Hamming {
			my_state.Hamming = result.Hamming
			my_state.Value = result.Value
		}
		n_results++

		if n_results == n_processes {
			// Got a result for each worker: close channel & exit loop
			close(ch)
		}
	}

	my_state.LastValue = int_to_str(start_value + n_values)

	fmt.Println("Best this run", my_state.Hamming, my_state.Value)

	// Persist state
	my_state.Save(state_path)
}
