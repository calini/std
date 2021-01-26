package main

import (
	"fmt"
	"log"
)

func MergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	return merge(MergeSort(arr[:len(arr)/2]), MergeSort(arr[len(arr)/2:]))
}

func merge(a, b []int) (result []int) {
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	if i < len(a) {
		for ; i < len(a); i++ {
			result = append(result, a[i])
		}
	}

	if j < len(b) {
		for ; j < len(b); j++ {
			result = append(result, b[j])
		}
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

func main() {
	test := []int{0, 8, 3, 9, -3, 4, 1, -1, 2, 6}

	test = MergeSort(test)

	if IsSortedInts(test) {
		fmt.Println("Sorted!")
	} else {
		log.Fatalf("ERROR SORTING: %v\n", test)
	}
}