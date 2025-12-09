package main

import "cmp"

func binarySearch[T cmp.Ordered](arr []T, item T) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2
		guess := arr[mid]

		switch {
		case guess == item:
			return mid
		case guess < item:
			low = mid + 1
		default:
			high = mid - 1
		}
	}

	return -1
}
