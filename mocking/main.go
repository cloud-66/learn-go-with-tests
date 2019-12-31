package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type DefaultSleeper struct {
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

const finalWorld = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWorld)
}
