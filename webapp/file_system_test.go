package main

import (
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("Reading Db", func(t *testing.T) {
		database, closeFile := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)

		defer closeFile()

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertFileSystem(t, got, want)

	})

	t.Run("Getting player score", func(t *testing.T) {
		database, closeFile := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)

		defer closeFile()

		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Chris")

		want := 33

		assetPlayerscore(t, got, want)

	})

	t.Run("store existing player wins", func(t *testing.T) {
		database, closeFile := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)

		defer closeFile()

		store := FileSystemPlayerStore{database}

		store.RecordWins("Chris")

		got := store.GetPlayerScore("Chris")

		want := 34
		assetPlayerscore(t, got, want)

	})
	t.Run("store new player ", func(t *testing.T) {
		database, closeFile := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)

		defer closeFile()

		store := FileSystemPlayerStore{database}

		store.RecordWins("Pepper")

		got := store.GetPlayerScore("Pepper")

		want := 1
		assetPlayerscore(t, got, want)

	})

}
func assetPlayerscore(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}

}

func createTempFile(t testing.TB, initalData string) (io.ReadWriteSeeker, func()) {
	t.Helper()
	fs, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}
	fs.Write([]byte(initalData))

	removefunc := func() {

		fs.Close()
		os.Remove(fs.Name())

	}

	return fs, removefunc

}

func assertFileSystem(t *testing.T, got, want []Player) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)

	}

}
