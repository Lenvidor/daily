package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := New()
	m := make(map[int]int)
	for i := 0; i < 5000; i++ {
		r.Init(false)
		r.Push(0)
		r.Push(1)
		r.Push(2)
		r.Push(3)
		r.Push(4)
		r.Push(5)
		if result, ok := r.Get().(int); ok {
			m[result]++
		} else {
			panic("something is wrong")
		}
	}
	fmt.Println(m)
}

type RandomFromStream struct {
	length        int
	value         interface{}
	randGenerator *rand.Rand
}

func New() *RandomFromStream {
	return new(RandomFromStream).Init(true)
}

func (r *RandomFromStream) Init(initGenerator bool) *RandomFromStream {
	r.length = 0
	r.value = nil
	if initGenerator {
		fmt.Println("generator is init")
		seed := rand.NewSource(time.Now().UnixNano())
		r.randGenerator = rand.New(seed)
	}
	return r
}

func (r *RandomFromStream) Push(value interface{}) {
	sv := r.randGenerator.Intn(r.length + 1)
	if sv >= r.length {
		r.value = value
	}
	r.length++
}

func (r *RandomFromStream) Get() interface{} {
	return r.value
}
