package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello Go")
	i := 1
	fmt.Println(i)
}

type Scale struct {
	notes []int
	span  int
}

func newScale(notes []int) {

}
