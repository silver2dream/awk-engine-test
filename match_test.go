package rally

import "testing"

func TestMatchLine(t *testing.T) {
	tests := []struct {
		name             string
		gamesA, gamesB   int
		pointsA, pointsB int
		want             string
	}{
		{name: "deuce late in the set", gamesA: 5, gamesB: 4, pointsA: 3, pointsB: 3, want: "games 5-4, deuce"},
		{name: "start of the set", want: "games 0-0, love-all"},
		{name: "one-game lead at six", gamesA: 6, gamesB: 5, pointsA: 4, pointsB: 3, want: "games 6-5, advantage first"},
		{name: "negative point count uses game call fallback", pointsA: -1, want: "games 0-0, -love"},
		{name: "first wins the set", gamesA: 6, gamesB: 3, pointsA: 2, pointsB: 1, want: "first wins the set 6-3"},
		{name: "second wins the set in a tie-break", gamesA: 6, gamesB: 7, pointsA: 3, pointsB: 3, want: "second wins the set 6-7"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchLine(tt.gamesA, tt.gamesB, tt.pointsA, tt.pointsB); got != tt.want {
				t.Errorf("MatchLine(%d, %d, %d, %d) = %q, want %q", tt.gamesA, tt.gamesB, tt.pointsA, tt.pointsB, got, tt.want)
			}
		})
	}
}
