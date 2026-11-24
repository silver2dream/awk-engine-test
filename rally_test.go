package rally

import "testing"

func TestPlayers(t *testing.T) {
	if Players != 2 {
		t.Errorf("Players = %d, want 2", Players)
	}
}

func TestScore(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want string
	}{
		{name: "single-digit points", a: 1, b: 2, want: "1-2"},
		{name: "double-digit and zero points", a: 10, b: 0, want: "10-0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Score(tt.a, tt.b); got != tt.want {
				t.Errorf("Score(%d, %d) = %q, want %q", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestWinner(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want string
	}{
		{name: "first wins", a: 4, b: 2, want: "first wins 4-2"},
		{name: "second wins", a: 1, b: 4, want: "second wins 1-4"},
		{name: "tie", a: 3, b: 3, want: "no winner 3-3"},
		{name: "zero-point tie", a: 0, b: 0, want: "no winner 0-0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Winner(tt.a, tt.b); got != tt.want {
				t.Errorf("Winner(%d, %d) = %q, want %q", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
