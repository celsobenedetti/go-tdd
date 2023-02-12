package main



type InMemoryPlayerStore struct {
    wins map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
    return &InMemoryPlayerStore{map[string]int{}}
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) ( int, bool ) {
    wins, ok := s.wins[name]
    return wins, ok
}

func (s *InMemoryPlayerStore) RecordWin(name string)  {
    s.wins[name]++
}

func (s *InMemoryPlayerStore) GetLeague() ( league []Player ) {
    for player, wins := range s.wins {
        league = append(league, Player{player, wins})
    }
    return
}
