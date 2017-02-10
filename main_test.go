package main

import (
	"testing"
)

func TestFirst(t *testing.T) {
	countme := 1
	for i := 1; i < 10; i++ {
		countme *= i
	}
	if countme != 123454 {
		t.Error("failed", countme)
	}
}
