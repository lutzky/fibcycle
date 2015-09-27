package fibcycle

import "testing"

func TestFibGen(t *testing.T) {
	wantedStates := []FibState{
		{0, 1}, {1, 1}, {1, 2}, {2, 3}, {3, 5}, {5, 8}, {8, 3},
	}

	c := FibGen(FibState{0, 1})

	for i, want := range wantedStates {
		got := <-c
		if got != want {
			t.Errorf("FibGen({0 1})[%d] == %d; want %d", i, got, want)
		}
	}
}

func TestFirstUnused(t *testing.T) {
	testCases := []struct {
		used []FibState
		want FibState
	}{
		{
			[]FibState{FibState{1, 0}, FibState{2, 0}},
			FibState{3, 0},
		},
		{
			[]FibState{FibState{0, 0}, FibState{2, 0}},
			FibState{1, 0},
		},
		{
			[]FibState{
				FibState{0, 0}, FibState{1, 0}, FibState{2, 0}, FibState{3, 0},
				FibState{4, 0}, FibState{5, 0}, FibState{6, 0}, FibState{7, 0},
				FibState{8, 0}, FibState{9, 0},
			},
			FibState{0, 1},
		},
	}

	for _, tc := range testCases {
		usedMap := map[FibState]bool{}
		for _, x := range tc.used {
			usedMap[x] = true
		}

		got := FirstUnused(usedMap)
		if got != tc.want {
			t.Errorf("FirstUnused(%v) = %v; want %v", tc.used, got, tc.want)
		}
	}
}
