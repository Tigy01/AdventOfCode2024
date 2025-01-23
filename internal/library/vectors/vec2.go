package vectors

import "math"

type Vec2 struct {
	X int
	Y int
}
type Vec2Float struct {
	X float64
	Y float64
}

// DEPRICATED: prefer to use vec1 == vec2
func (self Vec2) Eql(other Vec2) bool {
	return self.X == other.X && self.Y == other.Y
}

// Sums the X and Y of the vectors
func (self Vec2) Add(other Vec2) Vec2 {
	return Vec2{self.X + other.X, self.Y + other.Y}
}

// Difference of the X and Y of the vectors
func (self Vec2) Sub(other Vec2) Vec2 {
	return Vec2{self.X - other.X, self.Y - other.Y}
}

// Checks if vector is within a range given: low.X <= self.X < high.X and low.Y <= self.Y < high.Y
func (self Vec2) IsInRange(low Vec2, high Vec2) bool {
	if low.X <= self.X && self.X < high.X {
		if low.Y <= self.Y && self.Y < high.Y {
			return true
		}
	}
	return false
}

func (self Vec2) Magnitude() float64 {
	x := float64(self.X)
	y := float64(self.Y)
	return math.Sqrt(x*x + y*y)
}

func (self Vec2Float) Magnitude() float64 {
	x := self.X
	y := self.Y
	return math.Sqrt(x*x + y*y)
}

func (self Vec2) Normalize() (normalized Vec2Float) {
	if self.X == 0 {
		normalized.X = 0
        normalized.Y = float64(self.Y) / math.Abs(float64(self.Y))
		return
	}

	if self.Y == 0 {
		normalized.Y = 0
		normalized.X = float64(self.X) / math.Abs(float64(self.X))
		return
	}

	slope := math.Abs(float64(self.Y) / float64(self.X))

	normalized.X = math.Sqrt(1.0 / ((slope * slope) + 1))
	if self.X < 0 {
		normalized.X *= -1
	}
	normalized.Y = math.Sqrt(1 - normalized.X*normalized.X)
	if self.Y < 0 {
		normalized.Y *= -1
	}
	return
}
