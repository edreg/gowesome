package hello

import "fmt"

func main() {
	fmt.Println(Hello("", "en"))
}

func Hello(name string, language string) string {

	to := "world"
	if len(name) > 0 {
		to = name
	}

	hello := getHelloByLanguage(language)

	return hello + " " + to
}

func getHelloByLanguage(language string) string {
	hello := "hello"
	switch language {
	case "de":
		hello = "Hallo"
	case "es":
		hello = "Ola"
	default:
		hello = "hello"
	}
	return hello
}
