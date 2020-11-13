package library

import (
	"fmt"
)

// AddUpTo func O(n)
func AddUpTo(num int) int {
	var total int = 0

	for i := 1; i <= num; i++ {
		total += i

	}

	return total
}

// AddUpToOpt func Big O(1)
func AddUpToOpt(num int) int {
	return num * (num + 1) / 2
}

// CountUpAndDown func Big O(n+m) => O(n)
func CountUpAndDown(n int) {
	fmt.Println("Going up!")
	for i := 1; i < n; i++ {
		fmt.Println(i)
	}
	fmt.Println("At the top!\nGoing down...")
	for j := n - 1; j >= 0; j-- {
		fmt.Println(j)
	}
	fmt.Println("Back down. Bye!")
}

// PrintAllPairs func Big O(n+m) => O(n)
func PrintAllPairs(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Println(i, j)
		}
	}
}

// NumberOfHalves func Big O(n)
func NumberOfHalves(n int) int {
	var count int = 0
	for n > 1 {
		n /= 2
		count++
	}
	return count
}

// TotalNumberOfHalves func Big O(n+m) => O(n)
func TotalNumberOfHalves(n int) int {
	var total = 0
	for i := 0; i < n; i++ {
		total += NumberOfHalves(n)
	}
	return total
}

// PrintAtLeast5 func Big O(n)
func PrintAtLeast5(n int) {
	for i := 0; i < Max(5, n); i++ {
		fmt.Println(i)
	}
}

// PrintAtMost5 func Big O(1) karena konstan min di 5
func PrintAtMost5(n int) {
	for i := 0; i < Min(5, n); i++ {
		fmt.Println(i)
	}
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
