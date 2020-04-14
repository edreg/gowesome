package main

import (
	"github.com/edreg/awesome/basics/clock"
	"os"
	"time"
)

//go run main.go > clock.svg
func main() {
	clock.SVGWriter(os.Stdout, time.Now())
}
