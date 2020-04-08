package main

import (
	. "fmt"
	"io"
	"os"
	"strconv"
	"time"
)

const write = "write"
const sleep = "sleep"

type Sleeper interface {
	Sleep()
}

type CountdownOperationsSpy struct {
	Calls []string
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func main() {
	sleeper := &ConfigurableSleeper{
		duration: 1 * time.Second,
		sleep:    time.Sleep,
	}
	Countdown(os.Stdout, "3", sleeper)
}

func Countdown(writer io.Writer, number string, sleeper Sleeper) {
	start, err := strconv.ParseInt(number, 10, 64)

	if err == nil {
		for i := start; i > 0; i-- {
			sleeper.Sleep()
			Fprintf(writer, strconv.Itoa(int(i))+"\n")
		}

		sleeper.Sleep()
		Fprintf(writer, "Go!")
	}
}
