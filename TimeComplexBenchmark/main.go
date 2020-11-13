package main

import "fmt"

// Value struct
type Value struct {
	num int
}

// AddUpTo func1 O(N)
func (v Value) AddUpTo() int {
	var total int = 0

	for i := 1; i <= v.num; i++ {
		total += i
	}
	return total
}

// AddUpToOpt func2 Big O(1)
func (v Value) AddUpToOpt() int {
	return v.num * (v.num + 1) / 2
}

func main() {
	var Val Value
	Val.num = 9223372010
	fmt.Println(Val.AddUpTo())
	fmt.Println(Val.AddUpToOpt())
}
