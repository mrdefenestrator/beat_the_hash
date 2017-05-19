package common

import (
	"unicode/utf8"
	"math"
	"unicode"
	"github.com/aead/skein"
	"bytes"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Converts unicode str to list of integers
func UnicodeToList(value string) []rune {
	var result []rune

	for _, item := range value {
		result = append(result, item)
	}

	return result
}

// Converts list of integers to unicode str
func ListToUnicode(value []rune) string {
	var temp []byte
	var buffer bytes.Buffer

	for _, num := range value {
		utf8.EncodeRune(temp, num)
		buffer.Write(temp)
	}

	return buffer.String()
}

// Converts integer to a list of ints of base x
func ToBase(value int, base int) []int {
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
func FromBase(value []int, base int) int {
	var result int = 0
	for i, num := range value {
		result += num * int(math.Pow(float64(base), float64(i)))
	}

	return result
}

// Generator for creating value guesses
func GenGuess(n int, start int) chan int {
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
func HashIt(value []byte) []byte {
	hash := skein.New(1024, &skein.Config{})
	return hash.Sum(value)
}

// Return the bitwise hamming distance
func HammingIt(truth []byte, guess []byte) int {
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
