package rally

const (
	// BestOfThree is the maximum number of sets in a best-of-three match.
	BestOfThree = 3
	// BestOfFive is the maximum number of sets in a best-of-five match.
	BestOfFive = 5
)

// SetsToWin returns the number of sets needed to win a best-of-three or
// best-of-five match. It returns 0 for any other match length.
func SetsToWin(bestOf int) int {
	switch bestOf {
	case BestOfThree:
		return 2
	case BestOfFive:
		return 3
	default:
		return 0
	}
}
