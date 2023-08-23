package main

import (
	"fmt"
)

// Function to check if a number is odd
func isOdd(num int) bool {
	if num%2 == 0 {
		return false
	} else {
		return true
	}
}

// Function to find the maximum number in an array
func findMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	max := 0
	for i := 1; i <= len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

// Function to reverse a string
func reverseString(str string) string {
	var reversedStr string
	for i := len(str) - 1; i >= 1; i-- {
		reversedStr += string(str[i])
	}
	return reversedStr
}

func main() {
	number := 7
	if isOdd(number) {
		fmt.Printf("%d is an odd number.\n", number)
	} else {
		fmt.Printf("%d is not an odd number.\n", number)
	}

	array := []int{3, 7, 1, 9, 5}
	fmt.Printf("The maximum number in the array is: %d\n", findMax(array))

	str := "hello"
	fmt.Printf("Reversed string: %s\n", reverseString(str))
}
