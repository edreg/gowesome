package reader

import (
	"encoding/csv"
	"fmt"
	"github.com/edreg/awesome/quiz/problem"
	"github.com/edreg/awesome/quiz/quiz"
	"github.com/edreg/awesome/quiz/types"
	"io"
	"os"
)

func ParseFile(csvFile *os.File) (quiz.Quiz, error) {
	// Parse the file
	r := csv.NewReader(csvFile)
	records := make(map[string]string)

	problems := make([]problem.Problem, len(records))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if len(record) != 2 {
			err = types.Error(types.ExpectedRecordLengthTwo)
			return quiz.Quiz{}, err
		}
		problems = append(problems, problem.New(record))

		fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
	}

	return quiz.New(problems), nil
}

func GetFile(fileName string) (*os.File, error) {
	csvFile, err := os.Open(fileName)
	if err != nil {
		return csvFile, types.Error(types.CouldNotOpenFile)
	}
	return csvFile, err
}
