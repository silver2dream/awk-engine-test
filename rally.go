package rally

import "strconv"

// Players is how many players a rally has.
const Players = 2

// Score returns the two players' points joined by a dash.
func Score(a, b int) string {
	return strconv.Itoa(a) + "-" + strconv.Itoa(b)
}

// Winner returns the winning player or "no winner" for a tie, followed by the score.
func Winner(a, b int) string {
	score := Score(a, b)
	switch {
	case a > b:
		return "first wins " + score
	case b > a:
		return "second wins " + score
	default:
		return "no winner " + score
	}
}
