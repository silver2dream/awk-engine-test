package rally

// Deuce returns the umpire's deuce or advantage call once both players have
// at least three points. It returns an empty string for all other scores.
func Deuce(a, b int) string {
	if a < 3 || b < 3 {
		return ""
	}

	switch call := GameScore(a, b); call {
	case "deuce", "advantage first", "advantage second":
		return call
	default:
		return ""
	}
}
