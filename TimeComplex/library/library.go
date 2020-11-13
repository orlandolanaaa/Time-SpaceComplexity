package library

import (
	"fmt"
)

// AddUpTo func O(1)+n*O(1)+O(1) = O(n)
func AddUpTo(num int) int {
	var total int = 0
	//O(1)

	for i := 1; i <= num; i++ {
		total += i
		//n*O(1)
	}
	return total
	//O(1)
}

// AddUpToOpt func Big O(1)
func AddUpToOpt(num int) int {
	return num * (num + 1) / 2 // O(1)
}

// CountUpAndDown func Big O(n+m) => O(n)
func CountUpAndDown(n int) {
	fmt.Println("Going up!")
	//O(1)
	for i := 1; i < n; i++ {
		fmt.Println(i)
		//n*O(1)
	}
	fmt.Println("At the top!\nGoing down...")
	for j := n - 1; j >= 0; j-- {
		fmt.Println(j)
		//m*O(1)
	}
	fmt.Println("Back down. Bye!")
	//O(1)
}

// PrintAllPairs func Big O(n*n) = O(n2)
func PrintAllPairs(n int) {
	for i := 0; i < n; i++ {
		//n*O(1)
		for j := 0; j < n; j++ {
			fmt.Println(i, j)
			//n*O(1)
		}
	}
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
