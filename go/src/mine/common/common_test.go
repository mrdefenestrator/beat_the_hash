package common

import (
	"testing"
)

// hammingTriple is a struct for holding values for use as hamming distance
// test vectors
type hammingTriple struct {
	truth []byte
	check []byte
	dist  int
}

// equal returns true if both int arrays are equal, else false
func equal(a []int, b []int) bool {
	var length int

	if len(a) < len(b) {
		length = len(a)
	} else {
		length = len(b)
	}

	for i := 0; i < length; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// genN returns byte array containing n items of value
func genN(value byte, n int) []byte {
	values := make([]byte, n, n)

	for i := 0; i < n; i++ {
		values = append(values, value)
	}

	return values
}

func TestUnicodeToList(t *testing.T) {
	drift := "ドリフト"
	list := []int{0x30c9, 0x30ea, 0x30d5, 0x30c8}

	list2 := UnicodeToList(drift)

	if !equal(list2, list) {
		t.Error("Expected ", list, " got ", list2)
	}
}

func TestListToUnicode(t *testing.T) {
	drift := "ドリフト"
	list := []int{0x30c9, 0x30ea, 0x30d5, 0x30c8}

	drift2 := ListToUnicode(list)

	if drift2 != drift {
		t.Error("Expected ", drift, " got ", drift2)
	}
}

func TestToBase(t *testing.T) {
	num := 305419896
	list := []int{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8}

	list2 := ToBase(num, 16)

	for i := range list {
		if list2[i] != list[i] {
			t.Error("Expected ", list, " got ", list2)
		}
	}
}

func TestFromBase(t *testing.T) {
	num := 305419896
	list := []int{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8}

	num2 := FromBase(list, 16)

	if num2 != num {
		t.Error("Expected ", num, " got ", num2)
	}
}

func TestGenGuess(t *testing.T) {
	n := 10
	start := 15

	for i := start; i < n; i++ {
		guess := <-GenGuess(n, start)
		if guess != i {
			t.Error("Expected ", i, " got ", guess)
		}
	}
}

func TestHashIt(t *testing.T) {

}

func TestHammingIt(t *testing.T) {
	longZero := genN(byte(0x00), 128)
	longOne := genN(byte(0xff), 128)

	vectors := []hammingTriple{
		{
			[]byte{byte(0x10)},
			[]byte{byte(0x01)},
			2,
		},
		{
			[]byte{byte(0x01)},
			[]byte{byte(0x01)},
			0,
		},
		{
			[]byte{byte(0x00), byte(0x10)},
			[]byte{byte(0x00), byte(0x01)},
			2,
		},
		{
			[]byte{byte(0x10), byte(0x10)},
			[]byte{byte(0x01), byte(0x01)},
			4,
		},
		{
			longOne,
			longOne,
			0,
		},
		{
			longZero,
			longOne,
			1024,
		},
	}

	for n, vector := range vectors {
		dist := CalcHamming(vector.truth, vector.check)
		if vector.dist != dist {
			t.Error("Test ", n, " expected ", vector.dist, " got ", dist)
		}
	}

}
