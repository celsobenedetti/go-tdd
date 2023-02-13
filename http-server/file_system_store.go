package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// FileSystemPlayerStore stores information about players in a JSON file
type FileSystemPlayerStore struct {
	database json.Encoder
	league   League
}

func NewFileSystemPlayerStore(db *os.File) (*FileSystemPlayerStore, error) {
	err := initializePlayerDBFile(db)

	if err != nil {
		return nil, fmt.Errorf("problem initializing player db file %s, %v", db.Name(), err)
	}

	league, err := ParseLeagueJSON(db)

	if err != nil {
		return nil, fmt.Errorf("problem loading store from file %s, %v", db.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: *json.NewEncoder(&tape{db}),
		league:   league}, nil
}

func (s *FileSystemPlayerStore) GetLeague() League {
    sort.Slice(s.league, func(i, j int) bool {
        return s.league[i].Wins > s.league[j].Wins
    })
	return s.league
}

func (s *FileSystemPlayerStore) GetPlayerScore(playerName string) int {
	player := s.league.Find(playerName)

	if player != nil {
		return player.Wins
	}
	return 0
}

func (s *FileSystemPlayerStore) RecordWin(playerName string) {
	player := s.league.Find(playerName)

	if player != nil {
		player.Wins++
	} else {
		s.league = append(s.league, Player{playerName, 1})
	}

	// NewEncoder Writes on s.database, which is an instance of 'tape', which wrapes the dataabse file
	s.database.Encode(s.league)
}

func initializePlayerDBFile(file *os.File) error {
	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
	}

	file.Seek(0, 0) // return reader to beginning of file

	return nil
}
