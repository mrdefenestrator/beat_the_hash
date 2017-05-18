package main

import (
	"os"
	"fmt"
	"math"
	"bytes"
	"unicode"
	"mine/state"
	"mine/common"
	"encoding/hex"
	"unicode/utf8"
	"github.com/aead/skein"
)

// TODO: multithreading/processing using goroutines on worker?
// TODO: tests for each of the functions I implement
// TODO: program for posting results to site
// TODO: figure out GOPATH / project organization

// Converts unicode str to list of integers
func unicode_to_list(value string) []rune {
	var result []rune

	for _, item := range value {
		result = append(result, item)
	}

	return result
}

// Converts list of integers to unicode str
func list_to_unicode(value []rune) string {
	var temp []byte
	var buffer bytes.Buffer

	for _, num := range value {
		utf8.EncodeRune(temp, num)
		buffer.Write(temp)
	}

	return buffer.String()
}

// Converts integer to a list of ints of base x
func to_base(value int, base int) []int {
	var (
		result []int
		n_items int
	)

	if value > 0 {
		n_items = int(math.Ceil(math.Log(float64(value)) / math.Log(unicode.MaxRune)))
	} else {
		n_items = 1
	}

	for i := int(0); i < n_items; i++ {
		result = append(result, value / int(math.Pow(float64(base), float64(i))))
	}

	return result
}


// Converts to integer from a list of ints of base x
func from_base(value []int, base int) int {
	var result int = 0
	for i, num := range value {
		result += num * int(math.Pow(float64(base), float64(i)))
	}

	return result
}

// Generator for creating value guesses
func gen_guess(n int, start int) chan int {
	ch := make (chan int)

	go func () {
		for i := start; i < n + start; i++ {
			ch <- i
		}
		close(ch)
	} ()

	return ch
}

// Get the skein2014 hash of the value
func hash_it(value []byte) []byte {
	hash := skein.New(1024, &skein.Config{})
	return hash.Sum(value)
}

// Return the bitwise hamming distance
func hamming_it(truth []byte, guess []byte) int {
	var dist int = 0
	var length int
	var diff byte

	if len(truth) < len(guess) {
		length = len(truth)
	} else {
		length = len(guess)
	}

	for i := 0; i < length; i++ {
		diff = truth[i] ^ guess[i]
		for j := 0; j < 8; j++ {
			if (diff & byte(0x01)) != 0 {
				dist += 1
			}
			diff = diff >> 1
		}
	}

	return dist
}

//func worker(start int, n int, best_hamming int, best_value []byte) {
//	var last_value int = nil
//
//}

func main() {
	HEX_HASH := "5b4da95f5fa08280fc9879df44f418c8f9f12ba424b7757de02bbdfbae0d4c4fdf9317c80cc5fe04c6429073466cf29706b8c25999ddd2f6540d4475cc977b87f4757be023f19b8f4035d7722886b78869826de916a79cf9c94cc79cd4347d24b567aa3e2390a573a373a48a5e676640c79cc70197e1c5e7f902fb53ca1858b6"
	BIN_HASH, err := hex.DecodeString(HEX_HASH)
	common.Check(err)

	USAGE := "mine.py <n_values> <n_processes>\n\nn_values     number of values to mine\nn_processes  number of processes to use for mining"

	fmt.Print("TEST")

	args := os.Args
	if len(args) < 3 {
		fmt.Print("ERROR")
		fmt.Print(USAGE)
		os.Exit(-1)
	}

	digest := hash_it([]byte("asdffdsa"))

	s := state.State{}
	s.Load("state.yml")
	s.Save("main_state.yml")

	fmt.Print(BIN_HASH)
	fmt.Print(digest)
}
