// Package FibCycle provides tools for analzying cycles in modular fibonacci
// sequences.

package fibcycle

// FibState represents a fibonacci sequence state.
type FibState struct {
	A, B int
}

// Next gets the next state in the fibonacci sequence
func (fs FibState) Next() FibState {
	return FibState{fs.B, (fs.A + fs.B) % 10}
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
func (fs FibState) Equals(rhs FibState) bool {
	return fs.A == rhs.A && fs.B == rhs.B
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
		if state.Equals(fs) {
			return i, found
		}
	}
}

// FirstUnused returns the first unused state in used, ignoring {0 0}
func FirstUnused(used map[FibState]bool) FibState {
	state := FibState{1, 0}

	for !state.Equals(FibState{9, 9}) {
		if !used[state] {
			return state
		}
		state.A++
		if state.A == 10 {
			state.B++
			state.A = 0
		}
	}
	return FibState{-1, -1}
}
