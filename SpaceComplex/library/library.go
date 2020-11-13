package library

import (
	"fmt" 
	"math/rand"
)

// ShoutWhatsUpDawg O(1)
func ShoutWhatsUpDawg(n int) {
	
	for i := 0; i < n; i++ {
		fmt.Println("Hi guys !")	
	}

}

// SumArr O(1)
func SumArr (arr []int) int{
	var total int
	for _, num := range arr {
    	total+=num
	}
	return total

}

// ReverseString O(n)
func ReverseString (str string) string{
	var reversedString  string = ""
	
	for i := 0; i < len(str); i++ {
		reversedString = string(str[i]) + reversedString	
	}
	
	return reversedString
}

// KeepRandom O(n/2) => O(n)
func KeepRandom (arr []int) []int{
	var resArr []int
	
	for _, s := range arr {
		if (rand.Intn(100)<= 50) {
			resArr = append(resArr,s)
		}
	}
	
	return resArr
}
