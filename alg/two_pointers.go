// Задача на метод 2 указателей:
//В шеренгу, по росту, выстроились программисты, тестировщики, дизайнеры и СТО,
//необходимо выбрать наибольшее количество людей, но так, чтобы рост 2 любых людей в выборке,
//не отличался больше чем K.

package alg

func TwoPointersAlg(arr []int, maxChange int) int {
	var count, lastJ int

	lastJ = 0 // hotfix. could be idx
	for idx, val := range arr {

		for j := lastJ + 1; j <= len(arr)-1; j++ {

			lastJ = j

			if (arr[j] - val) <= maxChange {
				if (j-idx)+1 > count {
					count = (j - idx) + 1
				}
			}

			if (arr[j] - val) > maxChange {
				break
			}
		}
	}

	return count
}
