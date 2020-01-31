package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(45678)
	fmt.Printf("Your favorite number is %d ", rand.Intn(100))
	fmt.Println("Hello awesome! The time is ", time.Now())
}

func AddTwo(x int) (result int) {
	result = x + 2
	return result
}

func Swap(x, y string) (string, string) {
	return y, x
}
