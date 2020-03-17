package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSelection(t *testing.T) {
	arr := []int{10, 2, 5, 0, 6, 8, 1}
	ethalonArr := []int{0, 1, 2, 5, 6, 8, 10}

	selectionSort(&arr)

	fmt.Println(arr)

	compareInt(arr, ethalonArr, t)
}

func TestSelectionPerformance(t *testing.T) {
	arr := make([]int, SIZE)

	for i := 0; i < SIZE-1; i++ {
		arr[i] = rand.Int()
	}

	selectionSort(&arr)
	for idx, val := range arr {
		if idx != len(arr)-1 {
			if val > arr[idx+1] {
				t.Errorf("faile quick")
				t.Fail()
			}
		}
	}
}
