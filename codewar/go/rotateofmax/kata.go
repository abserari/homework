package main

import (
	"container/list"
	"fmt"
)

func MaxRot(n int64) int64 {
	var x list.List
	for n/10 != 0 {
		x.PushFront(n % 10)
		n /= 10
	}
	x.PushFront(n)
	for i := 0; i < x.Len()-1; i++ {
		y := x.Front()
		for j := 0; j < i; j++ {
			y = x.Front().Next()
		}
		x.PushBack(x.Remove(y))
	}
	var m int64
	for i := 0; i < x.Len(); i++ {
		m = (x.Remove(x.Front()).(int64) + m*10)
		x.PushBack(1)
	}
	return m
}

func main() {
	fmt.Println(MaxRot(56789))
}
