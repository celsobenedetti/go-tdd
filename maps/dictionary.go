package dictionary

import "errors"

type Dictionary map[string]string

var (
	ErrKeyNotFound       = errors.New("key not found in dictionary")
	ErrWordAlreadyExists = errors.New("word already exists in dictionary")
)

func (d Dictionary) Search(key string) (definition string, err error) {
	if definition, ok := d[key]; ok {
		return definition, nil
	}
	return "", ErrKeyNotFound
}

func (d Dictionary) Add(word, definition string) error {
	if _, ok := d[word]; ok {
		return ErrWordAlreadyExists
	}
	d[word] = definition
	return nil
}
