package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides determines if a triangle is equilateral, isosceles, or scalene
func KindFromSides(a, b, c float64) Kind {
	if (a <= 0 || b <= 0 || c <= 0) || (a+b < c || a+c < b || b+c < a) {
		return NaT
	}

	switch {
	case a == b && a == c && b == c:
		return Equ
	case (a == b) || (b == c) || (a == c):
		return Iso
	case (a != b) && (b != c) && (a != c):
		return Sca
	}

	return NaT
}
