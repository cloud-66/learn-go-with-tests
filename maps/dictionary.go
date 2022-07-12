package maps

import "errors"

var ErrNotFound = errors.New("could not find the  word you were looking for")
var ErrWordExists = errors.New("cannot add word because it already exists")

type Dictionary map[string]string

type DictionaryErr string

func (d Dictionary) Search(word string) (string, error) {

	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (e DictionaryErr) Error() string {
	return string(e)
}
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	d[word] = definition
	return nil
}
