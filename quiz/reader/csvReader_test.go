package reader_test

import (
	"github.com/edreg/awesome/quiz/reader"
	"github.com/edreg/awesome/quiz/types"
	"testing"
)

func TestGetFile(t *testing.T) {
	t.Run("get error if file does not exist", func(t *testing.T) {
		_, err := reader.GetFile("someFileName.ext")

		if err == nil {
			t.Fatal(types.WantedAnErrorButDidNot)
		}

		if err.Error() != types.CouldNotOpenFile {
			t.Fatalf("got %s want %s", err.Error(), types.CouldNotOpenFile)
		}
	})
}

func TestReader(t *testing.T) {
	t.Run("get error for error containing file", func(t *testing.T) {
		file, _ := reader.GetFile("../data/problems_errors.csv")
		_, err := reader.ParseFile(file)

		if err == nil {
			t.Fatal(types.WantedAnErrorButDidNot)
		}

		if err.Error() != types.ExpectedRecordLengthTwo {
			t.Fatalf("got %s want %s", err.Error(), types.ExpectedRecordLengthTwo)
		}
	})
}
