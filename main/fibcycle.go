package main

import (
	"fmt"

	fc "github.com/lutzky/fibcycle"
)

func analyzeFibLoop(initial fc.FibState) {
	loopLength, used := initial.FindCycle()
	fmt.Printf("%v loops in %d steps\n", initial, loopLength)
	for _, thru := range []fc.FibState{
		fc.NewFibState([]uint{0, 1}, 10),
		fc.NewFibState([]uint{0, 2}, 10),
		fc.NewFibState([]uint{1, 0}, 10),
		fc.NewFibState([]uint{2, 0}, 10),
	} {
		fmt.Printf("%v goes through %v: %t\n", initial, thru, used[thru])
	}
	fmt.Printf("%v first nonzero unused state: %v\n", initial, fc.FirstUnused(used))
}

func main() {
	for _, state := range []fc.FibState{
		fc.NewFibState([]uint{0, 0}, 10),
		fc.NewFibState([]uint{0, 1}, 10),
		fc.NewFibState([]uint{2, 0}, 10),
		fc.NewFibState([]uint{1, 0}, 10),
	} {
		analyzeFibLoop(state)
	}
}
