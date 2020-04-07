package iteration

func RepeatString(word string, times int) string {
	result := ""

	for i := 0; i < times; i++ {
		result += word
	}

	return result
}
