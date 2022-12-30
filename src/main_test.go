package main_test

import "testing"

func TestAddition(t *testing.T) {
	got := 2 + 2
	exptected := 4
	if got != exptected {
		t.Errorf("Did not get exptected results. Got '%v', wanted: '%v'", got, exptected)
	}
}

func TestSubtraction(t *testing.T) {
	got := 10 - 5
	exptected := 4
	if got != exptected {
		t.Errorf("Did not get exptected results. Got '%v', wanted: '%v'", got, exptected)
	}
}
