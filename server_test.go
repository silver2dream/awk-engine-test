package rally

import "testing"

func TestServer(t *testing.T) {
	tests := []struct {
		name string
		game int
		want string
	}{
		{name: "first game", game: 1, want: "first"},
		{name: "second game", game: 2, want: "second"},
		{name: "third game", game: 3, want: "first"},
		{name: "zero", game: 0, want: ""},
		{name: "negative", game: -1, want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Server(tt.game); got != tt.want {
				t.Errorf("Server(%d) = %q, want %q", tt.game, got, tt.want)
			}
		})
	}
}
