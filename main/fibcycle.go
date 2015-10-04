package main

import (
	"fmt"

	fc "github.com/lutzky/fibcycle"
)

func analyzeFibLoop(initial fc.FibState, usedGlobally map[fc.FibState]bool) {
	loopLength, used := initial.FindCycle()
	fmt.Printf("%v loops in %d steps\n", initial, loopLength)
	for k, v := range used {
		usedGlobally[k] = v
	}
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
	used := map[fc.FibState]bool{}
	for state := fc.NewFibState([]uint{0, 0}, 10); !state.Equals(&fc.FibState{}); state = fc.FirstUnused(used) {
		analyzeFibLoop(state, used)
	}
}
