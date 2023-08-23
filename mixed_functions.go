package main

import (
	"fmt"
)

// I check for:
// Readability
// Understandability
// Performance
// Error handling
// Bugs & Logical flaws
// DRY-Principle -> Don't write redundant code
// SOP-Principle -> Each thing does its own thing
// KISS-Principle -> Try to keep it simple
// SOLID-Principle -> https://en.wikipedia.org/wiki/SOLID
// POLA-Principle -> Code should work as one would expect
// YAGNI-Principle -> Don't code stuff you might not need

// Removed comment because function is self-explanatory
// Removed abbreviations to make code easier to read and understand
// If-else was redundant
func isOdd(number int) bool {
	return number%2 != 0
}

// Removed comment and renamed function to make it self-explanatory
// "Array" and "Index" were not very descriptive.
// Fixed for-loop, the index was offset by 1 which caused an error.
// Adjusted the code, so it works with negative numbers too.
func getLargestValue(numberCollection []int) int {
	if len(numberCollection) == 0 {
		return 0
	}

	largestValue := numberCollection[0]
	for numberIndex := 0; numberIndex < len(numberCollection); numberIndex++ {
		if numberCollection[numberIndex] > largestValue {
			largestValue = numberCollection[numberIndex]
		}
	}

	return largestValue
}

// Removed comment because function is self-explanatory
// Removed abbreviations to make code easier to read and understand
// Fixed for-loop, the index was offset by 1 which caused an error.
func reverseString(inputString string) string {
	var reversedString string

	for characterIndex := len(inputString) - 1; characterIndex >= 0; characterIndex-- {
		reversedString += string(inputString[characterIndex])
	}

	return reversedString
}

func main() {
	number := 6
	if isOdd(number) {
		fmt.Printf("%d is an odd number.\n", number)
	} else {
		fmt.Printf("%d is not an odd number.\n", number)
	}

	fmt.Printf("The maximum number in the array is: %d\n", getLargestValue([]int{3, 7, 1, 9, 11}))

	fmt.Printf("Reversed string: %s\n", reverseString("hello"))
}
