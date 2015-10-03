// Package fibcycle provides tools for analzying cycles in modular fibonacci
// sequences.
package fibcycle

import "fmt"

// FibState represents a Fibonacci sequence state.
type FibState struct {
	state  uint64
	base   uint64
	length uint
}

func (fs FibState) sum() uint64 {
	sum := uint64(0)
	for i := uint(0); i < fs.length; i++ {
		sum += (fs.state % fs.base)
		fs.state /= fs.base
	}
	return sum
}

// NewFibState creates a Fibonacci state. For "last decimal digit of the classic
// Fibonacci series", use NewFibState([0,1],10).
func NewFibState(state []uint, base uint64) FibState {
	result := FibState{
		base:   base,
		length: uint(len(state)),
	}
	for _, x := range state {
		result.state *= base
		result.state += uint64(x)
	}
	return result
}

func pow(x uint64, y uint) uint64 {
	if y == 0 {
		return 1
	}
	result := uint64(1)
	for i := uint(0); i < y; i++ {
		result *= x
	}
	return result
}

// Next gets the next state in the fibonacci sequence
func (fs FibState) Next() FibState {
	sum := fs.sum()
	fs.state %= pow(fs.base, fs.length-1)
	fs.state *= fs.base
	fs.state += sum % fs.base
	return fs
}

func (fs FibState) String() string {
	// TODO(lutzky): This only makes sense if fs.base == 10
	return fmt.Sprint(fs.state)
}

// Increment sets fs to the next possible FibState, with no relation to the
// Fibonacci rule.
func (fs *FibState) Increment() {
	// The fs.state representation makes this particularly simple.
	fs.state++
}

// FibGen returns a channel which will return an infinite sequence of FibStates,
// starting at initial, each the Next() of the previous.
func FibGen(initial FibState) chan FibState {
	c := make(chan FibState)
	go func(c chan FibState) {
		state := initial
		c <- state

		for true {
			nextState := state.Next()
			c <- nextState
			state = nextState
		}
	}(c)
	return c
}

// Equals returns true iff fs == rhs
func (fs *FibState) Equals(rhs *FibState) bool {
	return fs.state == rhs.state && fs.base == rhs.base && fs.length == rhs.length
}

// FindCycle returns the cycle length of fs, and the set (map->bool) of states
// it goes through.
func (fs FibState) FindCycle() (uint, map[FibState]bool) {
	c := FibGen(fs)
	<-c
	found := map[FibState]bool{}
	for i := uint(1); ; i++ {
		state := <-c
		found[state] = true
		if state.Equals(&fs) {
			return i, found
		}
	}
}

// FirstUnused returns the first unused state in used, ignoring the zero-state.
// If no unused states are found, the zero-state is returned.
func FirstUnused(used map[FibState]bool) FibState {
	state := NewFibState([]uint{0, 0}, 10)

	end := NewFibState([]uint{9, 9}, 10)

	for !state.Equals(&end) {
		state.Increment()
		if !used[state] {
			return state
		}
	}
	return FibState{}
}
