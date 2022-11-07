package pointers

type Dictionary map[string]string
type DictionaryErr string

func (de DictionaryErr) Error() string {
	return string(de)
}

const (
	ErrNonExistingKey = DictionaryErr("could find a value mapped to key")
	ErrKeyExists      = DictionaryErr("key already exists")
)

/**
A gotcha with maps is that they can be a nil value.
A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic.
*/
// Initialize empty map like following
// var dictionary = map[string]string{}
// OR
// var dictionary = make(map[string]string)
// read https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/maps#pointers-copies-et-al

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]

	if !ok {
		return "", ErrNonExistingKey
	}

	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, ok := d[key]

	if ok {
		return ErrKeyExists
	}
	d[key] = value
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
