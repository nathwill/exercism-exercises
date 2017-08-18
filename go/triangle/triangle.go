// Package triangle contains methods for handling triangles
package triangle

import (
	"math"
	"sort"
)

const (
	testVersion = 3 // match the testVersion for our test package
	NaT         = iota // not a triangle
	Equ         = iota // equilateral
	Iso         = iota // isosceles
	Sca         = iota // scalene
	Deg         = iota // degenerate
)

// Kind reflects the triangle type enumerated by the exported constants
type Kind int

// KindFromSides detects the kind of triangle from the lengths of its sides
func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}
	sort.Sort(sort.Float64Slice(sides))

	switch {
	case math.IsNaN(sides[0]) || sides[0] <= 0 || math.IsInf(sides[2], 0) || sides[0]+sides[1] < sides[2]:
		return Kind(NaT)
	case sides[0] == sides[1] && sides[1] == sides[2]:
		return Kind(Equ)
	case sides[0] == sides[1] || sides[1] == sides[2] || sides[0] == sides[2]:
		return Kind(Iso)
	default:
		return Kind(Sca)
	}
}
