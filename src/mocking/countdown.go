package mocking

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const constdownStart = 3

func Countdown(out io.Writer) {
	for i := constdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Fprint(out, finalWord)
}
