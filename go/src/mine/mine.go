package main

import (
	"encoding/hex"
	"fmt"
	"mine/common"
	"mine/state"
	"os"
	"strconv"
	"unicode/utf8"
)

var usage string = "mine.py <n_values> <n_processes> <state_path>\n\n" +
	"n_values     number of values to mine\n" +
	"n_processes  number of processes to use for mining\n" +
	"state_path   path to YAML state file\n"

var hexHash string = "5b4da95f5fa08280fc9879df44f418c8f9f12ba424b7757de02bb" +
	"dfbae0d4c4fdf9317c80cc5fe04c6429073466cf29706b8c25999ddd2f6540d4475" +
	"cc977b87f4757be023f19b8f4035d7722886b78869826de916a79cf9c94cc79cd43" +
	"47d24b567aa3e2390a573a373a48a5e676640c79cc70197e1c5e7f902fb53ca1858b6"

var binHash []byte

// strToInt converts unicode string to integer
func strToInt(value string) int {
	return common.FromBase(common.UnicodeToList(value), utf8.MaxRune)
}

// intToStr converts integer to unicode string
func intToStr(value int) string {
	return common.ListToUnicode(common.ToBase(value, utf8.MaxRune))
}

// worker mines for hash values
func worker(
	start int, nValues int, bestHamming int, bestValue string,
	ch chan state.State,
) {
	var lastValue []byte

	for guessChan := range common.GenGuess(nValues, start) {
		value := []byte(
			common.ListToUnicode(common.ToBase(guessChan, utf8.MaxRune)))

		hamming := common.CalcHamming(binHash, common.HashIt(value))

		if hamming < bestHamming {
			// Maximize
			fmt.Println("New Best", hamming, value)
			bestHamming = hamming
			bestValue = string(value)
		}

		lastValue = value
	}

	ch <- state.State{
		Hamming:   bestHamming,
		Value:     bestValue,
		LastValue: string(lastValue),
	}
}

// mine does some mining based on persistent state
func mine(nValues int, nProcesses int, statePath string) {
	// Check if state file exists
	var myState state.State
	if _, err := os.Stat(statePath); err == nil {
		// Load state
		myState.Load(statePath)
	} else {
		// Generate initial state
		myState = state.State{Hamming: 1024, Value: "", LastValue: ""}
	}

	fmt.Println("before", myState)
	fmt.Println("")

	// Create channel for results
	ch := make(chan state.State)

	// Divide work & start multiple workers
	startValue := strToInt(myState.LastValue)
	for i := 0; i < nProcesses; i++ {
		start := (nValues/nProcesses)*i + startValue
		go worker(
			start,
			nValues/nProcesses,
			myState.Hamming,
			myState.Value,
			ch,
		)
	}

	// Determine best result
	nResults := 0
	for result := range ch {
		if result.Hamming < myState.Hamming {
			myState.Hamming = result.Hamming
			myState.Value = result.Value
			fmt.Println("value", result.Value)
			fmt.Println("reset", myState)
			fmt.Println("")
		}
		nResults++

		if nResults == nProcesses {
			// Got a result for each worker: close channel & exit loop
			close(ch)
		}
	}

	myState.LastValue = intToStr(startValue + nValues)

	fmt.Println("Best this run", myState.Hamming, myState.Value)

	fmt.Println("after", myState)
	fmt.Println("")

	// Persist state
	myState.Save(statePath)
}

func main() {
	var err error
	// Decode hex hash to bytes
	binHash, err = hex.DecodeString(hexHash)
	common.Check(err)

	if len(os.Args) != 4 {
		fmt.Println(usage)
		os.Exit(-1)
	}

	nValues, err := strconv.Atoi(os.Args[1])
	common.Check(err)
	nProcesses, err := strconv.Atoi(os.Args[2])
	common.Check(err)
	statePath := os.Args[3]

	mine(nValues, nProcesses, statePath)
}
