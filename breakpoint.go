package rally

// BreakPoint reports whether winning the next point would win the ongoing game
// for the receiver, given the server's and receiver's current point counts.
func BreakPoint(server, receiver int) bool {
	return receiver >= 3 && receiver > server
}
