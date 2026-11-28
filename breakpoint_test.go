package rally

import (
	"math"
	"testing"
)

func TestBreakPoint(t *testing.T) {
	tests := []struct {
		name             string
		server, receiver int
		want             bool
	}{
		{name: "0-40", server: 0, receiver: 3, want: true},
		{name: "15-40", server: 1, receiver: 3, want: true},
		{name: "30-40", server: 2, receiver: 3, want: true},
		{name: "receiver advantage", server: 3, receiver: 4, want: true},
		{name: "receiver advantage at maximum score", server: math.MaxInt - 1, receiver: math.MaxInt, want: true},
		{name: "40-40", server: 3, receiver: 3, want: false},
		{name: "30-30", server: 2, receiver: 2, want: false},
		{name: "0-30", server: 0, receiver: 2, want: false},
		{name: "server advantage", server: 4, receiver: 3, want: false},
		{name: "negative receiver count", server: 0, receiver: -1, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BreakPoint(tt.server, tt.receiver); got != tt.want {
				t.Errorf("BreakPoint(%d, %d) = %t, want %t", tt.server, tt.receiver, got, tt.want)
			}
		})
	}
}
