package search

// Binary
//
//	does binary search on sorted list
//	returns index if found, else returns -1
func Binary[T int | float32 | float64](arr []T, n T) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == n {
			return mid
		}

		if n > arr[mid] {
			left = mid + 1
		}

		if n < arr[mid] {
			right = right - 1
		}
	}

	return -1
}
