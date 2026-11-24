package rally

import (
	"math"
	"testing"
)

func TestGameScore(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want string
	}{
		{name: "first leads early", a: 1, b: 0, want: "15-love"},
		{name: "second leads early", a: 2, b: 3, want: "30-40"},
		{name: "zero-point tie", a: 0, b: 0, want: "love-all"},
		{name: "two-point tie", a: 2, b: 2, want: "30-all"},
		{name: "first deuce", a: 3, b: 3, want: "deuce"},
		{name: "later deuce", a: 5, b: 5, want: "deuce"},
		{name: "advantage first", a: 4, b: 3, want: "advantage first"},
		{name: "advantage second after another deuce", a: 5, b: 6, want: "advantage second"},
		{name: "first wins at four", a: 4, b: 2, want: "game first"},
		{name: "first wins after deuce", a: 6, b: 4, want: "game first"},
		{name: "first wins without losing a point", a: 4, b: 0, want: "game first"},
		{name: "second wins after deuce", a: 3, b: 5, want: "game second"},
		{name: "negative count uses Call fallback", a: -1, b: 0, want: "-love"},
		{name: "large winning lead does not overflow", a: math.MaxInt, b: math.MinInt, want: "game first"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GameScore(tt.a, tt.b); got != tt.want {
				t.Errorf("GameScore(%d, %d) = %q, want %q", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
