package sort

func MergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	return merge(MergeSort(arr[:len(arr)/2]), MergeSort(arr[len(arr)/2:]))
}

func merge(a, b []int) (result []int) {

	for len(a) > 0 && len(b) > 0 {
		if a[0] < b[0] {
			result = append(result, a[0])
			a = a[1:]
		} else {
			result = append(result, b[0])
			b = b[1:]
		}
	}

	if len(a) > 0 {
		result = append(result, a...)
	}

	if len(b) > 0 {
		result = append(result, b...)
	}

	return result
}

func IsSortedInts(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}