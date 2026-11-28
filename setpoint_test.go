package rally

import (
	"math"
	"testing"
)

func TestSetPoint(t *testing.T) {
	tests := []struct {
		name                 string
		games, opponentGames int
		want                 bool
	}{
		{name: "5-3", games: 5, opponentGames: 3, want: true},
		{name: "5-4", games: 5, opponentGames: 4, want: true},
		{name: "6-5", games: 6, opponentGames: 5, want: true},
		{name: "6-6 tiebreak", games: 6, opponentGames: 6, want: true},
		{name: "4-4", games: 4, opponentGames: 4},
		{name: "5-5", games: 5, opponentGames: 5},
		{name: "3-5", games: 3, opponentGames: 5},
		{name: "4-3 below six", games: 4, opponentGames: 3},
		{name: "5-6 trailing after five", games: 5, opponentGames: 6},
		{name: "0-0", games: 0, opponentGames: 0},
		{name: "5-0", games: 5, opponentGames: 0, want: true},
		{name: "8-7 extended set", games: 8, opponentGames: 7, want: true},
		{name: "8-8 tied above six", games: 8, opponentGames: 8},
		{name: "negative opponent games", games: 5, opponentGames: -1},
		{name: "winning game count cannot overflow", games: math.MaxInt, opponentGames: math.MaxInt - 1, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetPoint(tt.games, tt.opponentGames); got != tt.want {
				t.Errorf("SetPoint(%d, %d) = %t, want %t", tt.games, tt.opponentGames, got, tt.want)
			}
		})
	}
}
