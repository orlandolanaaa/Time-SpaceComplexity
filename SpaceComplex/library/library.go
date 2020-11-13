package library

import (
	"fmt"
	"math/rand"
	"time"
)

/* alokasi memori variable = 4 byte
   setiap pemanggilan fungsi (Auxelary Space) = 4 byte */

// ShoutWhatsUpDawg O(1)
func ShoutWhatsUpDawg(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Hi guys !")
	}

	// i = 4b
	// AuxSpace = 4b
	// 4+4 = 8 => O(1)
}

// SumArr O(1)
func SumArr(arr []int) int {
	var total int
	for _, num := range arr {
		total += num
	}
	return total
	// total = 4b
	// num = 4b
	// AuxSpace = 4b
	// 4+4+4 = 12 => O(1)
}

// ReverseString O(n)
func ReverseString(str string) string {
	var reversedString string = ""

	for i := 0; i < len(str); i++ {
		reversedString = string(str[i]) + reversedString
	}

	return reversedString
	// reversedString = n * 4b
	// i = 4b
	// AuxSpace = 4b
	// (4n)+4+4 = 4n+8 => O(n)
}

// KeepRandom O(n/2) => O(n)
func KeepRandom(arr []int) []int {
	var resArr []int

	for _, s := range arr {
		rand.Seed(time.Now().UnixNano())
		if rand.Intn(100) <= 50 {
			resArr = append(resArr, s)
		}
	}

	return resArr
	// resArr = (n*4) / 2
	// s = 4b
	// AuxSpace = 4b
	// (4n/2)4+4= (4n/2)+8 => O(n)
}
