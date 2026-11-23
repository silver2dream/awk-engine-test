package rally

import "testing"

func TestPlayers(t *testing.T) {
	if Players != 2 {
		t.Errorf("Players = %d, want 2", Players)
	}
}
