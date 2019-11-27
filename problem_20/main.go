package main

import "fmt"

type node struct {
	value int
	next  *node
}

func main() {
	a10 := node{10, nil}
	a8 := node{8, &a10}
	a7 := node{7, &a8}
	a3 := node{3, &a7}
	a1 := node{1, &a8}
	a99 := node{99, &a1}
	printNodes(&a99)
	printNodes(&a3)
	fmt.Println(getIntersectionNode(&a3, &a99))
	printNodes(&a99)
	printNodes(&a3)
}

func printNodes(root *node) {
	defer fmt.Println()
	for {
		fmt.Print(root.value, " ")
		if root = root.next; root == nil {
			break
		}
	}
}

func getIntersectionNode(headA, headB *node) *node {
	lenA := getLen(headA)
	lenB := getLen(headB)

	maxLen, minLen := lenA, lenB
	if lenA < lenB {
		maxLen, minLen = lenB, lenA
	}
	maxHead, minHead := headA, headB
	if lenA < lenB {
		maxHead, minHead = headB, headA
	}

	for maxLen > minLen {
		maxHead, maxLen = maxHead.next, maxLen-1
	}
	for maxHead != minHead {
		maxHead, minHead = maxHead.next, minHead.next
	}

	return maxHead
}

func getLen(head *node) (result int) {
	if head == nil {
		return
	}
	for result = 1; head.next != nil; head, result = head.next, result+1 {
	}
	return
}
