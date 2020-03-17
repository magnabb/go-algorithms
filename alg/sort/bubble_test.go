package sort

import (
	"math/rand"
	"testing"
)

const SIZE = 1_000_000

func TestBubble(t *testing.T) {
	arr := []int{10, 2, 5, 0, 6, 8, 1}
	ethalonArr := []int{0, 1, 2, 5, 6, 8, 10}

	bubble(&arr)

	compareInt(arr, ethalonArr, t)
}

func compareInt(arr []int, ethalonArr []int, t *testing.T) {
	for idx, val := range arr {
		if val != ethalonArr[idx] {
			t.Errorf("failed")
			t.Fail()
		}
	}
}

func TestBubblePerfomance(t *testing.T) {
	arr := make([]int, SIZE)

	for i := 0; i < SIZE-1; i++ {
		arr[i] = rand.Int()
	}

	bubble(&arr)
	for idx, val := range arr {
		if idx != len(arr)-1 {
			if val > arr[idx+1] {
				t.Errorf("faile buble")
				t.Fail()
			}
		}
	}
}
