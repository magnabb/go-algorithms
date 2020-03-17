package sort

func selectionSort(arr *[]int) {
	selection(arr)
}

func selection(arr *[]int) {
	length := len(*arr) - 1
	for i := 0; i <= length; i++ {
		minIdx, min := i, (*arr)[i]

		for j := i + 1; j <= length; j++ {
			if (*arr)[j] < min {
				minIdx, min = j, (*arr)[j]
			}
		}

		if i != minIdx {
			(*arr)[i], (*arr)[minIdx] = (*arr)[minIdx], (*arr)[i]
		}
	}
}
