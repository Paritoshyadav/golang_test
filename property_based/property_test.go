package main

import (
	"testing"
)

func TestRomanNumber(t *testing.T) {
	number := 2
	got := convertToRomanNumber(number)

	if got != "I" {

		t.Errorf("got %s and want %s", got, "I")

	}
}
