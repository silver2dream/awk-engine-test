package rally

import "testing"

func TestBagel(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want bool
	}{
		{name: "first wins six to love", a: 6, b: 0, want: true},
		{name: "second wins six to love", a: 0, b: 6, want: true},
		{name: "opponent wins a game", a: 6, b: 1, want: false},
		{name: "seven-five set", a: 7, b: 5, want: false},
		{name: "no games played", a: 0, b: 0, want: false},
		{name: "one game short", a: 5, b: 0, want: false},
		{name: "one game over", a: 7, b: 0, want: false},
		{name: "negative game count", a: 6, b: -1, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bagel(tt.a, tt.b); got != tt.want {
				t.Errorf("Bagel(%d, %d) = %t, want %t", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
