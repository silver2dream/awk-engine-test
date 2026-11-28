package rally

import (
	"math"
	"testing"
)

func TestSetsToWin(t *testing.T) {
	tests := []struct {
		name   string
		bestOf int
		want   int
	}{
		{name: "best of three", bestOf: BestOfThree, want: 2},
		{name: "best of five", bestOf: BestOfFive, want: 3},
		{name: "negative length", bestOf: -3, want: 0},
		{name: "zero length", bestOf: 0, want: 0},
		{name: "best of one", bestOf: 1, want: 0},
		{name: "even length", bestOf: 4, want: 0},
		{name: "best of seven", bestOf: 7, want: 0},
		{name: "maximum length", bestOf: math.MaxInt, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetsToWin(tt.bestOf); got != tt.want {
				t.Errorf("SetsToWin(%d) = %d, want %d", tt.bestOf, got, tt.want)
			}
		})
	}
}
