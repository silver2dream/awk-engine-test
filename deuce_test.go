package rally

import (
	"math"
	"testing"
)

func TestDeuce(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want string
	}{
		{name: "first deuce", a: 3, b: 3, want: "deuce"},
		{name: "later deuce", a: 5, b: 5, want: "deuce"},
		{name: "advantage first", a: 4, b: 3, want: "advantage first"},
		{name: "advantage second", a: 3, b: 4, want: "advantage second"},
		{name: "advantage at maximum score", a: math.MaxInt, b: math.MaxInt - 1, want: "advantage first"},
		{name: "zero-point tie", a: 0, b: 0, want: ""},
		{name: "two-point tie", a: 2, b: 2, want: ""},
		{name: "one-point first lead before deuce", a: 3, b: 2, want: ""},
		{name: "one-point second lead before deuce", a: 2, b: 3, want: ""},
		{name: "first leads by two", a: 5, b: 3, want: ""},
		{name: "second leads by more than two", a: 3, b: 6, want: ""},
		{name: "negative count", a: math.MaxInt, b: math.MinInt, want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Deuce(tt.a, tt.b); got != tt.want {
				t.Errorf("Deuce(%d, %d) = %q, want %q", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
