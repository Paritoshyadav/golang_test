package main

import (
	"sync"
	"testing"
)

func TestSync(t *testing.T) {

	val := MyCounter{}

	var wg sync.WaitGroup

	counter := 1000

	wg.Add(counter)

	for i := 0; i < counter; i++ {

		go func() {
			val.add()
			wg.Done()

		}()

	}

	wg.Wait()
	got := val.value()

	if got != counter {

		t.Errorf("got %d and want %d", got, counter)

	}

}
