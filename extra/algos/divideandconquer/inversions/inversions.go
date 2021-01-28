package inversions

// MergeSortWithInversions returns the sorted array (using MergeSort), and the number of inversions it finds in an []int
// Example: [1, 5, 2, 4, 3], would have the following inversions: (5, 2), (5, 4), (5, 3), (4, 3), and hence return 4.
func MergeSortWithInversions(arr []int) ([]int, int) {

	if len(arr) <= 1 {
		return arr, 0
	}

	sortedFirstHalf, leftInversions := MergeSortWithInversions(arr[:len(arr)/2])
	sortedSecondHalf, rightInversions := MergeSortWithInversions(arr[len(arr)/2:])
	merged, crossInv := merge(sortedFirstHalf, sortedSecondHalf)

	return merged, (leftInversions + rightInversions + crossInv)
}

func merge(a, b []int) (result []int, crossInv int) {
	i, j := 0, 0

	for i < len(a) && j < len(b) { // while we can still merge
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
			crossInv += len(a) - i
		}
	}

	if i < len(a) { // if we have leftover in a
		result = append(result, a...)
	}

	if j < len(b) { // if we have leftover in b
		result = append(result, b...)
	}

	return
}