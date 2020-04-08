package dictionary

type Comparable struct {
	number int
	word   string
}
type Entry struct {
	key   string
	value string
}
type (
	Dictionary map[string]string
	Error      string
)

const (
	ErrNotFound         = Error("could not find the word you were looking for")
	ErrWordExists       = Error("this word is already included in your dictionary")
	ErrWordDoesNotExist = Error("can not find the word you are looking for")
)

func (e Error) Error() string {
	return string(e)
}

func (d Dictionary) Search(search Comparable) (string, error) {
	definition, ok := d[search.word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// pass by reference is not needed for maps; map is a "reference type" because maps can be huge => saving memory
// attempts to write to a nil map will cause a runtime panic
// so never do this: var m map[string]string
// var dictionary = map[string]string{}
// OR
// var dictionary = make(map[string]string)
func (d Dictionary) Add(entry Entry) error {
	_, err := d.Search(Comparable{word: entry.key})

	if err == ErrNotFound {
		d[entry.key] = entry.value
		return nil
	}

	if err == nil {
		return ErrWordExists
	}

	return err
}

func (d Dictionary) Update(entry Entry) error {
	_, err := d.Search(Comparable{word: entry.key})

	if err == ErrNotFound {
		return ErrWordDoesNotExist
	}

	if err == nil {
		d[entry.key] = entry.value
		return nil
	}

	return err
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(Comparable{word: key})

	if err == ErrNotFound {
		return ErrWordDoesNotExist
	}

	if err == nil {
		delete(d, key)
		return nil
	}

	return err
}
