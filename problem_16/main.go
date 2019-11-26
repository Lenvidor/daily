package main

import "fmt"

func main() {
	a := New(2)

	fmt.Println(a)
	fmt.Println(a.GetLast(0))
	a.Record(2)
	a.Record("Hello world")
	fmt.Println(a)
	fmt.Println(a.GetLast(0))
	a.Record("No 2")
	fmt.Println(a)
}

type CyclicList struct {
	n      int
	length int
	index  int
	Values []interface{}
}

func New(n int) *CyclicList { return new(CyclicList).Init(n) }

func (l *CyclicList) Init(n int) *CyclicList {
	l.n = n
	l.index = 0
	l.length = 0
	l.Values = make([]interface{}, n)
	return l
}

func (l *CyclicList) Record(v interface{}) {
	if l.length < l.n {
		l.Values[l.length] = v
		l.length++
	} else {
		l.Values[l.index] = v
		l.index = (l.index + 1) % l.n
	}
}

func (l *CyclicList) GetLast(i int) (interface{}, bool) {
	if i >= l.length {
		return nil, false
	}
	return l.Values[(l.index-i+l.n-1)%l.n], true
}
