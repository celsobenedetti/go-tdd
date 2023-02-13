package main

import (
	"io"
	"testing"

	"github.com/matryer/is"
)

func TestTape(t *testing.T) {
	is := is.New(t)

	t.Run("write", func(t *testing.T) {
		file, clean := createTempFile(t, "some_old_data_that_is_very_big")
		defer clean()

        newData := "New Data :)"

		tape := &tape{file}
		tape.Write([]byte(newData))

		file.Seek(0, 0)
		newFileContent, _ := io.ReadAll(file)

        is.Equal(newData, string(newFileContent))
	})
}
