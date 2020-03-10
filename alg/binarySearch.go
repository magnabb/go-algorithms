// Реализуйте бинарный поиск в массиве.

package alg

import "fmt"

func BinarySearch(arr []int, target int) (int, error) {
	left := 0
	right := len(arr) - 1

	for left <= right {
		midIdx := (left + right) / 2

		if arr[midIdx] == target {
			return midIdx, nil
		}

		if arr[midIdx] < target {
			left = midIdx + 1
		} else {
			right = midIdx - 1
		}
	}

	return 0, fmt.Errorf("element not found")
}
