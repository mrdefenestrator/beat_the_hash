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
			[]byte{byte(0x00) * 128},
			[]byte{byte(0xff) * 128},
			1024,
		},
		{
			[]byte{byte(0xff) * 128},
			[]byte{byte(0xff) * 128},
			0,
		},
	}

	for _, vector := range vectors {
		dist := HammingIt(vector.truth, vector.check)
		if vector.dist != dist {
			t.Fail()
		}
	}

}
