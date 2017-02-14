package main

import "fmt"

func main() {
	fmt.Println("music.main")
}

type Scale struct {
	degrees []int
	span    int
}

func NewScale(degrees []int) *Scale {
	return &Scale{degrees: degrees, span: 12}
}

// Intervals gets the intervals between the scale degrees
func (scale *Scale) Intervals() []int {
	return degreesToIntervals(scale.degrees, scale.span)
}

// degreesToIntervals gets the intervals between the given scale degrees in a range
//   TS: degrees.map((degree, i) => (degrees[i+1] || span) - degree);
func degreesToIntervals(degrees []int, span int) []int {
	intervals := make([]int, len(degrees))
	for i, degree := range degrees {
		if i+1 < len(degrees) {
			intervals[i] = degrees[i+1] - degree
		} else {
			intervals[i] = span - degree
		}
	}
	return intervals
}
