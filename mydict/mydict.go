package mydict

import "errors"

var errWordNotFound = errors.New("Word not found")
var errWordExists = errors.New("That word aleady exists")

// Dictionary type
type Dictionary map[string]string

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exist := d[word]
	if exist {
		return value, nil
	}
	return "", errWordNotFound
}

// Add a word to the Dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errWordNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}
