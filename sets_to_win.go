package rally

// SetsToWin returns the number of sets needed to win a best-of-three or
// best-of-five match. It returns 0 for any other match length.
func SetsToWin(bestOf int) int {
	switch bestOf {
	case 3:
		return 2
	case 5:
		return 3
	default:
		return 0
	}
}
