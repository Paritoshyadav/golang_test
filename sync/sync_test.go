package main

import "testing"

func TestSync(t *testing.T) {

	val := MyCounter{1}

	for i := 1; i < 10; i++ {

		val.add()

	}
	got := val.value()

	want := 10

	if got != want {

		t.Errorf("got %d and want %d", got, want)

	}

}
