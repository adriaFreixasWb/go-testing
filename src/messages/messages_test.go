package messages

import (
	"testing"
)

func TestGreetTableDriven(t *testing.T) {
	scenarios := []struct {
		input  string
		expect string
	}{
		{input: "Gopher", expect: "Hello, Gopher!\n"},
		{input: "", expect: "Hello, !\n"},
	}
	for _, s := range scenarios {
		got := Greet(s.input)
		if got != s.expect {
			t.Errorf("Did not get exptected result for input '%v'. Expected %q, got %q",
				s.input, s.expect, got)
		}
	}
}

func TestDepart(t *testing.T) {
	got := depart("Gopher")
	expect := "Goodbye, Gopher\n"
	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got: %q\n", expect, got)
	}
}
