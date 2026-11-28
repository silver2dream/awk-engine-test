package rally

// Bagel reports whether either player won the set six games to love.
func Bagel(a, b int) bool {
	return (a == 6 && b == 0) || (a == 0 && b == 6)
}
