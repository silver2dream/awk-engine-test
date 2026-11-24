package rally

import "testing"

func TestCall(t *testing.T) {
	tests := []struct {
		name   string
		points int
		want   string
	}{
		{name: "zero points", points: 0, want: "love"},
		{name: "one point", points: 1, want: "15"},
		{name: "two points", points: 2, want: "30"},
		{name: "three points", points: 3, want: "40"},
		{name: "below range", points: -1, want: ""},
		{name: "above range", points: 4, want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Call(tt.points); got != tt.want {
				t.Errorf("Call(%d) = %q, want %q", tt.points, got, tt.want)
			}
		})
	}
}
