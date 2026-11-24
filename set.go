package rally

// SetScore returns the set call for the two players' game counts.
func SetScore(a, b int) string {
	score := Score(a, b)
	switch {
	case a >= 6 && (b <= a-2 || a == 7 && b == 6):
		return "first wins the set " + score
	case b >= 6 && (a <= b-2 || b == 7 && a == 6):
		return "second wins the set " + score
	default:
		return "games " + score
	}
}
