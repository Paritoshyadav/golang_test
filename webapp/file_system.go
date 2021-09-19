package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func (fs *FileSystemPlayerStore) GetLeague() []Player {
	fs.database.Seek(0, 0)
	league, _ := Newleague(fs.database)

	return league

}

func (fs *FileSystemPlayerStore) GetPlayerScore(playerName string) int {
	data := fs.GetLeague()
	var wins int
	for _, v := range data {

		if v.Name == playerName {
			wins = v.Wins
			break

		}

	}
	return wins

}

func Newleague(rdr io.Reader) ([]Player, error) {
	var league []Player

	err := json.NewDecoder(rdr).Decode(&league)

	if err != nil {
		return nil, err
	}

	return league, nil

}
