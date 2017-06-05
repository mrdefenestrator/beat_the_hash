package state

import (
	"testing"
)

func TestState(t *testing.T) {
	s := State{
		432,
		"asdf",
		"fdsa",
	}
	s.Save("test.yml")

	var s2 State
	s2.Load("test.yml")

	if s2.Hamming != s.Hamming {
		t.Fail()
	}

	if s2.Value != s.Value {
		t.Fail()
	}

	if s2.LastValue != s.LastValue {
		t.Fail()
	}
}
