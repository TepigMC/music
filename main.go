package main

import "fmt"

func main() {
	fmt.Println("hello Go")
	i := 1
	fmt.Println(i)
	scale := NewScale([]int{0, 2, 4, 5, 7, 9, 11})
	fmt.Println(scale)
}

type Scale struct {
	degrees []int
	span    int
}

func NewScale(degrees []int) *Scale {
	return &Scale{degrees: degrees, span: 12}
}

/*func (scale *Scale) Intervals() []int {
	// TS: return this._notes.map((note, i) => (this._notes[i + 1] || this._range) - note);
	intervals := make([]int, scale.degrees.length)
	for i, degree := range scale.degrees {
		intervals = append(intervals, (scale.degrees[i+1] > scale.span ? scale.degrees[i+1] : scale.span) - degree)
	}
	return intervals
}*/
