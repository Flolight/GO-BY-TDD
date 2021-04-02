package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const FinalWord = "Go!"
const CountdownStart = 3

func Countdown(out io.Writer) {
	for i := CountdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, FinalWord)
}

func main() {
	Countdown(os.Stdout)
}
