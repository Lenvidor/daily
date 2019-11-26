package main

import (
	"container/list"
	"fmt"
)

func main() {
	MaxWindow([]int{10, 5, 2, 7, 8, 7}, 3)
}

func MaxWindow(arr []int, k int) {
	l := list.New()
	for i := 1 - k; i < len(arr)-k+1; i++ {
		j := i + k - 1
		for e := l.Back(); e != nil && arr[e.Value.(int)] <= arr[j]; e = l.Back() {
			l.Remove(e)
		}
		l.PushBack(j)

		if i >= 0 {
			fmt.Print(arr[l.Front().Value.(int)], " ")
		}

		if l.Front().Value == i {
			l.Remove(l.Front())
		}
	}
}
