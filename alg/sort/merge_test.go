package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMerge(t *testing.T) {
	arr := []int{10, 2, 5, 0, 6, 8, 1}
	ethalonArr := []int{0, 1, 2, 5, 6, 8, 10}

	mergeSort(&arr)

	fmt.Println(arr)

	compareInt(arr, ethalonArr, t)
}

func TestMergePerformance(t *testing.T) {
	arr := make([]int, SIZE)

	for i := 0; i < SIZE-1; i++ {
		arr[i] = rand.Int()
	}

	mergeSort(&arr)
	for idx, val := range arr {
		if idx != len(arr)-1 {
			if val > arr[idx+1] {
				t.Errorf("faile quick")
				t.Fail()
			}
		}
	}
}
