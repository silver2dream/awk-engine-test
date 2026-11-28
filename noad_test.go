package rally

import "testing"

func TestNoAdDeciding(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want bool
	}{
		{name: "40-40", a: 3, b: 3, want: true},
		{name: "30-30", a: 2, b: 2, want: false},
		{name: "40-30", a: 3, b: 2, want: false},
		{name: "30-40", a: 2, b: 3, want: false},
		{name: "0-0", a: 0, b: 0, want: false},
		{name: "tie past 40-40", a: 4, b: 4, want: false},
		{name: "negative point count", a: -1, b: 3, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoAdDeciding(tt.a, tt.b); got != tt.want {
				t.Errorf("NoAdDeciding(%d, %d) = %t, want %t", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
