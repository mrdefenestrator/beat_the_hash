package common

import "testing"

func TestCheck(t *testing.T) {

}

func TestUnicodeToList(t *testing.T) {

}

func TestListToUnicode(t *testing.T) {

}

func TestToBase(t *testing.T) {

}

func TestFromBase(t *testing.T) {

}

func TestGenGuess(t *testing.T) {

}

func TestHashIt(t *testing.T) {

}

type hamming_triple struct {
	truth []byte
	check []byte
	dist int
}

func TestHammingIt(t *testing.T) {
	long_zero := []byte{}
	long_one := []byte{}

	for i := 0; i < 128; i++ {
		long_zero = append(long_zero, byte(0x00))
		long_one = append(long_one, byte(0xff))
	}


	vectors := []hamming_triple{
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
			long_one,
			long_one,
			0,
		},
		{
			long_zero,
			long_one,
			1024,
		},
	}

	for n, vector := range vectors {
		dist := HammingIt(vector.truth, vector.check)
		if vector.dist != dist {
			t.Error("Test ", n, " expected ", vector.dist, " got ", dist)
		}
	}

}
