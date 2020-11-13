package main

import (
	. "SpaceComplex/library"
	"fmt"
)

func main() {
	fmt.Println("////////ShoutWhatsUpDawg////////")
	ShoutWhatsUpDawg(5)

	fmt.Println("////////ReverseString////////")
	var qasirReversed = ReverseString("qasir")
	fmt.Println(qasirReversed)

	fmt.Println("////////KeepRandom////////")
	var arrNum = [...]int{2, 6, 2, 4, 1}
	var keepRandArr = KeepRandom(arrNum[:])
	fmt.Println(keepRandArr)

	fmt.Println("////////SumArr////////")
	var sum = SumArr(arrNum[:])
	fmt.Println(sum)
}
