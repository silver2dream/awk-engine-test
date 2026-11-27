package rally

// Tiebreak returns the umpire's call for the two players' tiebreak points.
// It announces a winner once a player has at least seven points and leads by two or more.
func Tiebreak(a, b int) string {
	if (a >= 7 && b <= a-2) || (b >= 7 && a <= b-2) {
		return Winner(a, b)
	}
	return Score(a, b)
}
