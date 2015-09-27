package main

import (
	"fmt"

	fc "github.com/lutzky/fibcycle"
)

func analyzeFibLoop(initial fc.FibState) {
	loopLength, used := initial.FindCycle()
	fmt.Printf("%v loops in %d steps\n", initial, loopLength)
	for _, thru := range []fc.FibState{
		{0, 1}, {0, 2}, {1, 0}, {2, 0},
	} {
		fmt.Printf("%v goes through %v: %t\n", initial, thru, used[thru])
	}
	fmt.Printf("%v first nonzero unused state: %v\n", initial, fc.FirstUnused(used))
}

func main() {
	for _, state := range []fc.FibState{
		{0, 0},
		{0, 1},
		{2, 0},
		{1, 0},
	} {
		analyzeFibLoop(state)
	}
}
