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
