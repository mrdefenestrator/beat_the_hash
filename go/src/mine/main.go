package main

import (
	"os"
	"fmt"
	"mine/state"
	"mine/common"
	"encoding/hex"
)

// TODO: multithreading/processing using goroutines on worker?
// TODO: tests for each of the functions I implement
// TODO: program for posting results to site
// TODO: figure out GOPATH / project organization

//func worker(start int, n int, best_hamming int, best_value []byte) {
//	var last_value int = nil
//
//}

func main() {
	HEX_HASH := "5b4da95f5fa08280fc9879df44f418c8f9f12ba424b7757de02bbdfbae0" +
		"d4c4fdf9317c80cc5fe04c6429073466cf29706b8c25999ddd2f6540d4475cc977b" +
		"87f4757be023f19b8f4035d7722886b78869826de916a79cf9c94cc79cd4347d24b" +
		"567aa3e2390a573a373a48a5e676640c79cc70197e1c5e7f902fb53ca1858b6"
	BIN_HASH, err := hex.DecodeString(HEX_HASH)
	common.Check(err)

	USAGE := "mine.py <n_values> <n_processes>\n\n" +
		"n_values     number of values to mine\n" +
		"n_processes  number of processes to use for mining"

	args := os.Args
	if len(args) < 3 {
		fmt.Print(USAGE)
		os.Exit(-1)
	}

	digest := common.HashIt([]byte("asdffdsa"))

	s := state.State{}
	s.Load("state.yml")
	s.Save("main_state.yml")

	fmt.Print(BIN_HASH)
	fmt.Print(digest)
}
