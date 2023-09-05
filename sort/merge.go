package sort

// Merge
//
//	sorts array using merge sort algorithm
func Merge(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2
	left := Merge(arr[:mid])
	right := Merge(arr[mid:])

	return mergeArrays(left, right)
}

// merges two arrays, based on smallest values
func mergeArrays(arr1, arr2 []int) []int {
	tempArr := make([]int, 0)
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			tempArr = append(tempArr, arr1[i])
			i++
		} else {
			tempArr = append(tempArr, arr2[j])
			j++
		}
	}

	for ; i < len(arr1); i++ {
		tempArr = append(tempArr, arr1[i])
	}

	for ; j < len(arr2); j++ {
		tempArr = append(tempArr, arr2[j])
	}

	return tempArr
}
