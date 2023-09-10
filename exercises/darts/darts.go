package darts

import "math"

// Score returns the earned points in a single toss of a Darts game.
func Score(x, y float64) int {
	switch d := math.Hypot(x, y); {
	case d <= 1.0: // inner circle a  of 1
		return 10
	case d <= 5.0: // middle circle a radius of 5 units
		return 5
	case d <= 10.0: //outer circle has a radius of 10 units
		return 1
	default:
		return 0
	}
}
