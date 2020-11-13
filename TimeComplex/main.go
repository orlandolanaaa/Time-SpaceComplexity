package main

import (
	. "TimeComplex/library"
	"fmt"
)

func main() {
	fmt.Println("////////AddUpTo////////")
	var num = AddUpTo(5)
	fmt.Println(num)

	fmt.Println("////////AddUpToOpt////////")
	num = AddUpToOpt(5)
	fmt.Println(num)

	fmt.Println("////////CountUpAndDown////////")
	CountUpAndDown(5)

	fmt.Println("////////PrintAllPairs////////")
	PrintAllPairs(5)

	fmt.Println("////////NumberOfHalves////////")
	var count = NumberOfHalves(10)
	fmt.Println(count)

	fmt.Println("////////TotalNumberOfHalves////////")
	var countTotal = TotalNumberOfHalves(5)
	fmt.Println(countTotal)
}
