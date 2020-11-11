package sort

func Ints(arr []int) {

	return
}

func IsSortedInts(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}