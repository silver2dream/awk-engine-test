package rally

import "strconv"

// Players is how many players a rally has.
const Players = 2

// Score returns the two players' points joined by a dash.
func Score(a, b int) string {
	return strconv.Itoa(a) + "-" + strconv.Itoa(b)
}
