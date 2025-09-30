package main

func main() {
}

type (
	Dictionary    map[string]string
	DictionaryErr string
)

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrWordNotDefined     = DictionaryErr("couldn't find the word were looking for")
	ErrWordAlreadyDefined = DictionaryErr("word has already defined, can't modify")
)

func (d Dictionary) Search(k string) (string, error) {
	def, ok := d[k]
	if !ok {
		return "", ErrWordNotDefined
	}
	return def, nil
}

func (d Dictionary) Add(k, v string) error {
	_, err := d.Search(k)
	switch err {
	case ErrWordNotDefined:
		d[k] = v
	case nil:
		return ErrWordAlreadyDefined
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, description string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = description
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
