package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search(Comparable{word: "test"})
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search(Comparable{word: "falseWord"})

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		errAdd := dictionary.Add(Entry{key: "test", value: "this is just a test"})

		if errAdd != nil {
			assertError(t, errAdd, ErrWordExists)
		}

		_, errSearch := dictionary.Search(Comparable{word: "test"})

		if errSearch != nil {
			t.Fatal("should have found an added word:", errSearch)
		}

		assertDefinition(t, dictionary, Entry{key: "test", value: "this is just a test"})
	})

	t.Run("existing word", func(t *testing.T) {
		entry := Entry{key: "test", value: "this is just a test"}
		dictionary := Dictionary{entry.key: entry.value}
		err := dictionary.Add(Entry{key: "test", value: "this is just a new test"})

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, entry)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		entry := Entry{key: "test", value: "this is just a test"}
		updateEntry := Entry{key: "test", value: "this is just a new test"}
		dictionary := Dictionary{entry.key: entry.value}

		err := dictionary.Update(updateEntry)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, updateEntry)
	})

	t.Run("new word", func(t *testing.T) {
		entry := Entry{key: "test", value: "this is just a test"}
		dictionary := Dictionary{}

		err := dictionary.Update(entry)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		entry := Entry{key: "test", value: "this is just a test"}
		dictionary := Dictionary{entry.key: entry.value}

		err := dictionary.Delete(entry.key)
		assertError(t, err, nil)
		//assertDefinition(t, dictionary, updateEntry)
	})

	t.Run("new word", func(t *testing.T) {
		entry := Entry{key: "test", value: "this is just a test"}
		dictionary := Dictionary{}

		err := dictionary.Delete(entry.key)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, entry Entry) {
	t.Helper()

	got, err := dictionary.Search(Comparable{word: entry.key})
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if entry.value != got {
		t.Errorf("got %q want %q", got, entry.value)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}

	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error.")
	}
}
