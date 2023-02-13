package poker

import (
	"os"
)

// tape is an implementation of io.Writer
// It encapsulates a database/file which is an instance of io.ReadWriteSeeker
// Whenever the database/file is writen to, it will be wrapped by tape.Write(p []byte)
type tape struct {
    file *os.File
}

func (t *tape) Write(p []byte) (int, error)  {
    t.file.Truncate(0)
    t.file.Seek(0,0)
    return t.file.Write(p)
}
