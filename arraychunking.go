package Arraychunking

// --- Directions
// Given an array and chunk size, divide the array into many subarrays
// where each subarray is of length size
// --- Examples
// chunk([1, 2, 3, 4], 2) --> [[ 1, 2], [3, 4]]
// chunk([1, 2, 3, 4, 5], 2) --> [[ 1, 2], [3, 4], [5]]
// chunk([1, 2, 3, 4, 5, 6, 7, 8], 3) --> [[ 1, 2, 3], [4, 5, 6], [7, 8]]
// chunk([1, 2, 3, 4, 5], 4) --> [[ 1, 2, 3, 4], [5]]
// chunk([1, 2, 3, 4, 5], 10) --> [[ 1, 2, 3, 4, 5]]

func Arraychunking(arr []int, size int) (result [][]int) {
	i := 0
	var chunked [][]int
	for i < len(arr) {
		if i+size > len(arr) {
			chunked = append(chunked, arr[i:len(arr)])
		} else {
			chunked = append(chunked, arr[i:i+size])
		}
		i = i + size
	}
	return chunked
}
