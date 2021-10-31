package mocking

import (
	"fmt"
	"io"
)

const finalWord = "Go!"
const constdownStart = 3

func Countdown(out io.Writer) {
	for i := constdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}

	fmt.Fprint(out, finalWord)
}
