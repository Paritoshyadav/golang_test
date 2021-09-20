package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (fs *FileSystemPlayerStore) GetLeague() League {
	fs.database.Seek(0, 0)
	league, _ := Newleague(fs.database)

	return league

}

func (fs *FileSystemPlayerStore) GetPlayerScore(playerName string) int {
	data := fs.GetLeague()
	player := data.Find(playerName)
	return player.Wins

}

func (fs *FileSystemPlayerStore) RecordWins(playerName string) {
	data := fs.GetLeague()
	player := data.Find(playerName)
	if player != nil {
		player.Wins++
	} else {
		data = append(data, Player{playerName, 1})
	}

	fs.database.Seek(0, 0)
	json.NewEncoder(fs.database).Encode(data)

}

func Newleague(rdr io.Reader) ([]Player, error) {
	var league []Player

	err := json.NewDecoder(rdr).Decode(&league)

	if err != nil {
		return nil, err
	}

	return league, nil

}
