package rally

import (
	"math"
	"testing"
)

func TestTiebreak(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want string
	}{
		{name: "no points played", a: 0, b: 0, want: "0-0"},
		{name: "first leads early", a: 3, b: 2, want: "3-2"},
		{name: "two-point lead below seven", a: 6, b: 4, want: "6-4"},
		{name: "tied at six", a: 6, b: 6, want: "6-6"},
		{name: "first leads by one at seven", a: 7, b: 6, want: "7-6"},
		{name: "second leads by one above seven", a: 8, b: 9, want: "8-9"},
		{name: "first wins at seven", a: 7, b: 5, want: "first wins 7-5"},
		{name: "second wins above seven", a: 6, b: 8, want: "second wins 6-8"},
		{name: "winning lead cannot overflow", a: math.MaxInt, b: math.MinInt, want: "first wins " + Score(math.MaxInt, math.MinInt)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tiebreak(tt.a, tt.b); got != tt.want {
				t.Errorf("Tiebreak(%d, %d) = %q, want %q", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
