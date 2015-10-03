package fibcycle

import "testing"

func TestSum(t *testing.T) {
	testCases := []struct {
		state FibState
		want  uint64
	}{
		{NewFibState([]uint{0, 1}, 10), 1},
		{NewFibState([]uint{5, 5}, 10), 10},
	}

	for _, tc := range testCases {
		got := tc.state.sum()
		if got != tc.want {
			t.Errorf("%v.sum() == %v; want %v", tc.state, got, tc.want)
		}
	}
}

func TestFibGen(t *testing.T) {
	wantedStates := [][]uint{
		{0, 1}, {1, 1}, {1, 2}, {2, 3}, {3, 5}, {5, 8}, {8, 3}, {3, 1},
	}

	c := FibGen(NewFibState([]uint{0, 1}, 10))

	for i, want := range wantedStates {
		got := <-c
		fsWant := NewFibState(want, 10)
		if got != fsWant {
			t.Errorf("FibGen(01)[%d] == %v; want %v", i, got, fsWant)
		}
	}
}

func TestFirstUnused(t *testing.T) {
	testCases := []struct {
		used [][]uint
		want []uint
	}{
		{
			[][]uint{{0, 1}, {0, 2}},
			[]uint{0, 3},
		},
		{
			[][]uint{{0, 0}, {0, 2}},
			[]uint{0, 1},
		},
		{
			[][]uint{
				{0, 0}, {0, 1}, {0, 2}, {0, 3},
				{0, 4}, {0, 5}, {0, 6}, {0, 7},
				{0, 8}, {0, 9},
			},
			[]uint{1, 0},
		},
	}

	for _, tc := range testCases {
		usedMap := map[FibState]bool{}
		for _, x := range tc.used {
			usedMap[NewFibState(x, 10)] = true
		}

		got := FirstUnused(usedMap)
		fsWant := NewFibState(tc.want, 10)
		if got != fsWant {
			t.Errorf("FirstUnused(%v) = %v; want %v", tc.used, got, fsWant)
		}
	}
}
