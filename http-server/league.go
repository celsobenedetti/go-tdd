package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

// Find returns a reference to the first Player with matching name, or nil
func (l League) Find(playerName string) *Player {
	for i, p := range l {
		if p.Name == playerName {
			return &l[i]
		}
	}
    return nil
}

// ParseLeague from a reader with JSON format and return players
func ParseLeagueJSON(r io.Reader) (league League, err error) {
	err = json.NewDecoder(r).Decode(&league)

	if err != nil {
		err = fmt.Errorf("error parsing league JSON, %v", err)
	}

	return league, err
}
