package alg

func ChooseSort(arr []int) []int {
	res := make([]int, 0, len(arr))

	for range arr {
		min, minIdx := getMin(arr)
		res = append(res, min)

		arr = del(arr, minIdx)
	}

	return res
}

func del(arr []int, idx int) []int {
	arr[idx] = arr[len(arr)-1]
	arr[len(arr)-1] = 0
	return arr[:len(arr)-1]
}

func getMin(arr []int) (int, int) {
	min := arr[0]
	minIdx := 0

	for idx, val := range arr {
		if val < min {
			min = val
			minIdx = idx
		}
	}

	return min, minIdx
}
