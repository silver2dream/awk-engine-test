package rally

// SetPoint reports whether winning the game in progress would win the set.
// Negative game counts return false.
func SetPoint(games, opponentGames int) bool {
	return (games >= 5 && opponentGames >= 0 && opponentGames < games) ||
		(games == 6 && opponentGames == 6)
}
