package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("Reading Db", func(t *testing.T) {
		database := strings.NewReader(`[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertFileSystem(t, got, want)

	})

	t.Run("Getting player score", func(t *testing.T) {
		database := strings.NewReader(`[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)

		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Chris")

		want := 33

		if got != want {
			t.Errorf("got %d but want %d", got, want)
		}

	})

}

func assertFileSystem(t *testing.T, got, want []Player) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)

	}

}
