package alg

import (
	"testing"
)

func TestBinarySearchExistsElements(t *testing.T) {
	arr := getSlice()

	assert(t, arr, 8, 2)
	assert(t, arr, 70, 8)
	assert(t, arr, 1202, 15)
}

func TestBinarySearchNotExistsElements(t *testing.T) {
	arr := getSlice()

	assert(t, arr, 10, 0)
	assert(t, arr, 1098, 0)
	assert(t, arr, 1999, 0)
}

func getSlice() []int {
	return []int{1, 4, 8, 11, 21, 23, 24, 64, 70, 84, 125, 704, 1099, 1200, 1201, 1202, 2000}
}

func assert(t *testing.T, arr []int, target, result int) {
	if idx, _ := BinarySearch(arr, target); idx != result {
		t.Errorf("BinarySearch(%v, %d) exitet with %d. Expected %d", arr, target, idx, result)
	}
}
