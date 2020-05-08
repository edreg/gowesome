package types

type (
	Error string
)

func (e Error) Error() string {
	return string(e)
}

const CouldNotOpenFile = "Couldn't open the csv file"
const WantedAnErrorButDidNot = "wanted an error but didn't go one"
const ExpectedRecordLengthTwo = "expected records with a length of two"
