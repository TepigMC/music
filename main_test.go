package main

import (
	"reflect"
	"testing"
)

/*func TestFirst(t *testing.T) {
	countme := 1
	for i := 1; i < 10; i++ {
		countme *= i
	}
	if countme != 123454 {
		t.Error("failed", countme)
	}
}*/

func TestIntervals(t *testing.T) {
	scale := Scale{[]int{0, 2, 4, 5, 7, 9, 11}, 12}
	expected := []int{2, 2, 1, 2, 2, 2, 1}
	intervals := scale.Intervals()
	if !reflect.DeepEqual(intervals, expected) {
		t.Error("intervals", intervals, "is not equivalent to expected", expected)
	} else {
		t.Log("scale.Intervals passes")
	}
}

func TestDegreesToIntervals(t *testing.T) {
	degrees := []int{0, 2, 4, 5, 7, 9, 11}
	expected := []int{2, 2, 1, 2, 2, 2, 1}
	intervals := degreesToIntervals(degrees, 12)
	if !reflect.DeepEqual(intervals, expected) {
		t.Error("intervals", intervals, "is not equivalent to expected", expected)
	} else {
		t.Log("degreesToIntervals passes")
	}
}

/*func TestNewScale(t *testing.T) {
	scale := NewScale([]int{0, 2, 4, 5, 7, 9, 11})
}*/
