package rally

// Server returns "first" for odd-numbered games and "second" for even-numbered games.
// It returns an empty string for game numbers below 1.
func Server(game int) string {
	switch {
	case game < 1:
		return ""
	case game%Players == 0:
		return "second"
	default:
		return "first"
	}
}
