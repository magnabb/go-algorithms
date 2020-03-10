package alg

import (
	"testing"
)

func TestChooseSort(t *testing.T) {
	ethanolArr := []int{2,3,5,6,10}

	arr := []int{5,3,6,2,10}

	sortedArr := ChooseSort(arr)

	for idx, val := range sortedArr {
		if ethanolArr[idx] != val {
			t.Errorf("failed")
		}
	}
}