package main

import (
	"fmt"
	"reflect"
	"testing"
)

func assertCorrectMsg(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {

		t.Errorf("got %q and wanted %q", got, want)
	}

}

func TestHello(t *testing.T) {

	t.Run("with a arugument name", func(t *testing.T) {

		got := hello("test", "default")
		want := "Hello test"
		assertCorrectMsg(t, got, want)

	})
	t.Run("with a no name or any arugument", func(t *testing.T) {

		got := hello("", "default")
		want := "Hello world"

		assertCorrectMsg(t, got, want)

	})
	t.Run("with language arugument", func(t *testing.T) {

		got := hello("test", "french")
		want := "Hola test"

		assertCorrectMsg(t, got, want)

	})

}

func TestSumAll(t *testing.T) {

	assertSliceCorrectMsg := func(t testing.TB, got []int, want []int) {

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	}

	t.Run("Testing SUmm all", func(t *testing.T) {

		got := sumAll([]int{1, 2, 3}, []int{1, 2, 3})
		want := []int{6, 6}
		assertSliceCorrectMsg(t, got, want)

	})

}

func ExampleHello() {

	result := hello("test", "hindi")
	fmt.Println(result)
	// Output: Namesta test

}
