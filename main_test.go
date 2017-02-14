package main

import (
	"reflect"
	"testing"
)

var majorDegrees = []int{0, 2, 4, 5, 7, 9, 11}
var majorIntervals = []int{2, 2, 1, 2, 2, 2, 1}
var majorScale = Scale{majorDegrees, 12}

func testCompare(t *testing.T, title string, actual, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v: actual %v is not equivalent to expected %v", title, actual, expected)
	}
}

func TestIntervals(t *testing.T) {
	testCompare(t, "scale.Intervals", majorScale.Intervals(), majorIntervals)
}

func TestDegreesToIntervals(t *testing.T) {
	testCompare(t, "degreesToIntervals", degreesToIntervals(majorDegrees, 12), majorIntervals)
}

/*func TestNewScale(t *testing.T) {
	scale := NewScale([]int{0, 2, 4, 5, 7, 9, 11})
}*/
