package rally

import "strings"

// MatchLine returns the set and game calls, or just the final set call when the set is won.
func MatchLine(gamesA, gamesB, pointsA, pointsB int) string {
	set := SetScore(gamesA, gamesB)
	if strings.HasPrefix(set, "games ") {
		return set + ", " + GameScore(pointsA, pointsB)
	}
	return set
}
