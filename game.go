package rally

// GameScore returns the umpire's call for the two players' game points.
func GameScore(a, b int) string {
	switch {
	case a == b && a >= 3:
		return "deuce"
	case a == b:
		return Call(a) + "-all"
	case a <= 3 && b <= 3:
		return Call(a) + "-" + Call(b)
	case a > b:
		if a-1 == b {
			return "advantage first"
		}
		return "game first"
	default:
		if b-1 == a {
			return "advantage second"
		}
		return "game second"
	}
}
