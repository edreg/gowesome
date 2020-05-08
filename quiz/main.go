package main

import (
	"github.com/edreg/awesome/quiz/reader"
	"github.com/edreg/awesome/quiz/types"
	"log"
	"os"
)

func main() {
	file, err := reader.GetFile("data/problems.csv")
	if err != nil {
		log.Fatalln(types.CouldNotOpenFile, err)
	}

	quiz, parseErr := reader.ParseFile(file)
	if parseErr != nil {
		log.Fatalln(types.CouldNotOpenFile, err)
	}

	quiz.Run(os.Stdout, os.Stdin)
}
