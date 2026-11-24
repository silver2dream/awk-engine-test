package rally

// Call returns the tennis score call for a player's point count.
// It returns an empty string for counts outside zero through three.
func Call(points int) string {
	switch points {
	case 0:
		return "love"
	case 1:
		return "15"
	case 2:
		return "30"
	case 3:
		return "40"
	default:
		return ""
	}
}
