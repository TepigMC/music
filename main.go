package main

import "fmt"

func main() {
	fmt.Println("hello Go")
}

type Scale struct {
	degrees []int
	span    int
}

func NewScale(degrees []int) *Scale {
	return &Scale{degrees: degrees, span: 12}
}

func (scale *Scale) Intervals() []int {
	return degreesToIntervals(scale.degrees, scale.span)
}

// If degrees are not ascending, it adds the span until it is
func degreesToIntervals(degrees []int, span int) []int {
	// TS: return this._degrees.map((degree, i) => (this._degrees[i + 1] || this._span) - degree);
	intervals := make([]int, len(degrees))
	for i, degree := range degrees {
		var nextDegree int
		if i+1 < len(degrees) {
			nextDegree = degrees[i+1]
		}
		for nextDegree < degree {
			nextDegree += span
		}
		intervals[i] = nextDegree - degree
	}
	return intervals
}
