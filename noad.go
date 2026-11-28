package rally

// NoAdDeciding reports whether either player can win a no-advantage game on the
// next point. It returns true exactly when both players have three points (40-40).
func NoAdDeciding(a, b int) bool {
	return a == 3 && b == 3
}
