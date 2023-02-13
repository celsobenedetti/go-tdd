package poker

import (
	"os"
	"testing"

	"github.com/matryer/is"
)

func TestFileSystemStore(t *testing.T) {
	is := is.New(t)

	playerData := `[{"Name": "Celso", "Wins": 20},
    {"Name": "Joao", "Wins": 10}]`

	t.Run("get league from a reader", func(t *testing.T) {
		store, clean := createDatabaseAndStore(t, playerData)
		defer clean()

		got := store.GetLeague()
		want := League{
			{"Celso", 20},
			{"Joao", 10},
		}

		is.Equal(got, want) // store returned different league object
	})

	t.Run("get score for a player", func(t *testing.T) {
		store, clean := createDatabaseAndStore(t, playerData)
		defer clean()

		is.Equal(store.GetPlayerScore("Celso"), 20) // got different player score
	})

	t.Run("record win for existing player", func(t *testing.T) {
		store, clean := createDatabaseAndStore(t, playerData)
		defer clean()

		player := "Celso"
		store.RecordWin(player)

		is.Equal(store.GetPlayerScore(player), 21) // expected score to be increased by 1
	})

	t.Run("record win for new player", func(t *testing.T) {
		store, clean := createDatabaseAndStore(t, playerData)
		defer clean()

		player := "Jeniffer"
		store.RecordWin(player)

		is.Equal(store.GetPlayerScore(player), 1) // expected score for new player to be 1
	})

	t.Run("works with empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)
		is.NoErr(err) // should work with empty file
	})

	t.Run("store get league is sorted", func(t *testing.T) {
		playerData := `[{"Name": "Celso", "Wins": 1},
    {"Name": "Joao", "Wins": 2}]`

		database, cleanDatabase := createTempFile(t, playerData)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		is.NoErr(err)

		want := League{{"Joao", 2}, {"Celso", 1}}

		is.Equal(store.GetLeague(), want) // League should be sorted by wins

        //read again
		is.Equal(store.GetLeague(), want) // League should be sorted by wins
	})
}

func createDatabaseAndStore(t testing.TB, playerData string) (*FileSystemPlayerStore, func()) {
	t.Helper()

	database, clean := createTempFile(t, playerData)
	store, err := NewFileSystemPlayerStore(database)

	if err != nil {
		t.Fatalf("didnt expect an error, but got one, %v", err)
	}

	return store, clean
}

func createTempFile(t testing.TB, data string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("failed to create temp file %v", err)
	}

	tmpFile.Write([]byte(data))

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}
