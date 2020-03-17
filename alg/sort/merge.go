package sort

func mergeSort(arr *[]int) {
	merge(arr, 0, len(*arr)-1)
}

func merge(arr *[]int, startIdx, endIdx int) {
	if endIdx <= startIdx {
		return
	}

	pivotIdx := (endIdx + startIdx) / 2

	merge(arr, startIdx, pivotIdx)
	merge(arr, pivotIdx+1, endIdx)

	res := make([]int, 0, endIdx-startIdx+1)

	i, j := startIdx, pivotIdx+1

	for i <= pivotIdx && j <= endIdx {
		if (*arr)[i] < (*arr)[j] {
			res = append(res, (*arr)[i])
			i++
		} else {
			res = append(res, (*arr)[j])
			j++
		}
	}

	for i <= pivotIdx {
		res = append(res, (*arr)[i])
		i++
	}

	for j <= endIdx {
		res = append(res, (*arr)[j])
		j++
	}

	for k := startIdx; k <= endIdx; k++ {
		(*arr)[k] = res[k-startIdx]
	}
}
