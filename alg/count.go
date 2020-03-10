//  Задан отсортированный массив,
// необходимо подсчитать количество чисел, которые находятся в интервале [L; R]

package alg

func Count(arr []int, l, r int) int {
	return plainCount(arr, l, r)
}

func plainCount(arr []int, l, r int) int {
	var count int

	for _, val := range arr {
		if val > r {
			break
		}

		if val >= l && val <= r {
			count++
		}
	}

	return count
}