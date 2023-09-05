package search

// Linear
//
//	searches through the array/slice
//	returns index if found, else returns -1
func Linear(arr []int, n int) int {
	length := len(arr)
	for i := 0; i < length; i++ {
		if arr[i] == n {
			return i
		}
	}

	return -1
}
