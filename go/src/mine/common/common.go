package common

import (
	"bytes"
	"github.com/whyrusleeping/FastGoSkein"
	"math"
	"unicode/utf8"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Converts unicode str to list of integers
func UnicodeToList(value string) []int {
	var result []int

	for _, item := range value {
		result = append(result, int(item))
	}

	return result
}

// Converts list of integers to unicode str
func ListToUnicode(value []int) string {
	var (
		buffer bytes.Buffer
	)
	temp := make([]byte, 4, 4)

	for _, num := range value {
		n := utf8.EncodeRune(temp, rune(num))
		buffer.Write(temp[0:n])
	}

	return buffer.String()
}

// Converts integer to a list of ints of base x
func ToBase(value int, base int) []int {
	var nItems int
	result := []int{}

	if value > 0 {
		nItems = int(
			math.Ceil(math.Log(float64(value)) / math.Log(float64(base))))
	} else {
		nItems = 1
	}

	for i := nItems - 1; i >= 0; i-- {
		result = append(
			result, value/int(math.Pow(float64(base), float64(i)))%base)
	}

	return result
}

// Converts to integer from a list of ints of base x
func FromBase(value []int, base int) int {
	result := 0

	for i, num := range value {
		result += num * int(
			math.Pow(float64(base), float64(len(value)-i-1)))
	}

	return result
}

// Generator for creating value guesses
func GenGuess(n int, start int) chan int {
	ch := make(chan int)

	go func() {
		for i := start; i < n+start; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

// Get the skein2014 hash of the value
func HashIt(value []byte) []byte {
	digest := new(skein.Skein1024)
	digest.Init(1024)
	digest.Update(value)

	sum := make([]byte, 128)
	digest.Final(sum)
	return sum
}

// Return the bitwise hamming distance
func CalcHamming(truth []byte, guess []byte) int {
	var (
		dist   int = 0
		length int
		diff   byte
	)

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
