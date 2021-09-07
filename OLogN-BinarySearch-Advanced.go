package main

import "fmt"

func binarySearch(needle int, haystack []int, searchFront bool) int {

	// initialize the low and high point values, and default result value
	low := needle
	high := len(haystack) - 1
	result := -1

	// iterate through the array, from low to high
	for low <= high {

		// find the middle index
		mid := (low + high) / 2

		// when 'needle' is found at the middle of the array...
		if haystack[mid] == needle {

			// ...assign the middle value to the result placeholder (do not move the ladder, stay put for now)
			result = mid

			// determine which size of the sorted array to calculate
			if searchFront {
				high = mid - 1 // search from the front, moved down the ladder
			} else {
				low = mid + 1 // search from the end, move up the ladder
			}

		} else if needle < haystack[mid] { // ...when zero is found at the 'front' of the array
			// recalculate the 'high', moved down the ladder
			high = mid - 1
		} else { // ...otherwise, zero is found at the 'back' of the array
			// recalculate the 'low', move up the ladder
			low = mid + 1
		}
	}

	// return the calculated result (or -1 when nothing is found)
	return result
}

func main() {
	items := []int{-100, -10, -6, -3, -2, -1, -1, 0, 0, 0, 0, 3, 5, 6, 10, 40}
	firstIdx := binarySearch(0, items, true)
	lastIdx := binarySearch(0, items, false)

	fmt.Println((lastIdx - firstIdx) + 1) // should be '4', since there are 4 zeros in the array
}
