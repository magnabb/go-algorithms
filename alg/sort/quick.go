package sort

func quick(arr *[]int) []int {
	//return quickArrayCopying(*arr)
	quickPointers(arr, 0, len(*arr)-1)
	return *arr
}

func quickArrayCopying(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]

	var less, greater []int

	for idx, val := range arr {
		if idx == len(arr)/2 {
			continue
		}

		if val < pivot {
			less = append(less, val)
		} else if val >= pivot {
			greater = append(greater, val)
		}
	}

	return append(
		append(quickArrayCopying(less), pivot),
		quickArrayCopying(greater)...,
	)
}

func quickPointers(arr *[]int, startIdx, endIdx int) {
	l := startIdx
	r := endIdx
	pivot := (*arr)[(endIdx+startIdx)/2]

	for l <= r {
		for (*arr)[l] < pivot {
			l++
		}

		for (*arr)[r] > pivot {
			r--
		}

		if l <= r {
			t := (*arr)[r]
			(*arr)[r] = (*arr)[l]
			(*arr)[l] = t

			l++
			r--
		}
	}

	if l < endIdx {
		quickPointers(arr, l, endIdx)
	}
	if r > startIdx {
		quickPointers(arr, startIdx, r)
	}
}
