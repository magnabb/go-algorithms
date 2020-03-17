package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestQuick(t *testing.T) {
	arr := []int{10, 2, 5, 0, 6, 8, 1}
	ethalonArr := []int{0, 1, 2, 5, 6, 8, 10}

	quick(&arr)

	fmt.Println(arr)

	compareInt(arr, ethalonArr, t)
}

func TestQuickPerformance(t *testing.T) {
	arr := make([]int, SIZE)

	for i := 0; i < SIZE-1; i++ {
		arr[i] = rand.Int()
	}

	res := quick(&arr)
	for idx, val := range res {
		if idx != len(res)-1 {
			if val > res[idx+1] {
				t.Errorf("faile quick")
				t.Fail()
			}
		}
	}
}
