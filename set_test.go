package rally

import (
	"math"
	"testing"
)

func TestSetScore(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want string
	}{
		{name: "first wins at six", a: 6, b: 4, want: "first wins the set 6-4"},
		{name: "first wins at eight", a: 8, b: 6, want: "first wins the set 8-6"},
		{name: "first wins tie-break", a: 7, b: 6, want: "first wins the set 7-6"},
		{name: "second wins at six", a: 3, b: 6, want: "second wins the set 3-6"},
		{name: "second wins at eight", a: 6, b: 8, want: "second wins the set 6-8"},
		{name: "second wins tie-break", a: 6, b: 7, want: "second wins the set 6-7"},
		{name: "first leads below six", a: 5, b: 4, want: "games 5-4"},
		{name: "two-game lead below six", a: 5, b: 3, want: "games 5-3"},
		{name: "first leads by one at six", a: 6, b: 5, want: "games 6-5"},
		{name: "second leads by one at six", a: 5, b: 6, want: "games 5-6"},
		{name: "first leads by one at eight", a: 8, b: 7, want: "games 8-7"},
		{name: "tied at six", a: 6, b: 6, want: "games 6-6"},
		{name: "no games played", a: 0, b: 0, want: "games 0-0"},
		{name: "winning lead cannot overflow", a: math.MaxInt, b: math.MinInt, want: "first wins the set " + Score(math.MaxInt, math.MinInt)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetScore(tt.a, tt.b); got != tt.want {
				t.Errorf("SetScore(%d, %d) = %q, want %q", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
