package main

import (
	"fmt"
)

func main() {
	m := [][]int{{3, 3, 2}, {3, 3, 1}, {3, 1, 2}, {3, 1, 3}}
	m = [][]int{{2}}
	fmt.Println(FindCheapestWay(m))
}

func FindCheapestWay(costMatrix [][]int) int {
	type element struct {
		value int
		index int
	}

	var twoCheapest [2]element
	var prevTwoCheapest [2]element

	for _, column := range costMatrix {
		const maxInt = int(^uint(0) >> 1)
		if len(column) > 0 {
			twoCheapest = [2]element{{maxInt, -1}, {maxInt, -1}}
		} else {
			twoCheapest = [2]element{{0, -1}, {0, -1}}
		}
		for index, value := range column {
			if index == prevTwoCheapest[0].index {
				value += prevTwoCheapest[1].value
			} else {
				value += prevTwoCheapest[0].value
			}

			if twoCheapest[0].value > value {
				twoCheapest[0], twoCheapest[1] = element{value, index}, twoCheapest[0]
			} else if twoCheapest[1].value > value {
				twoCheapest[1] = element{value, index}
			}
		}
		prevTwoCheapest = twoCheapest
	}

	return twoCheapest[0].value
}
