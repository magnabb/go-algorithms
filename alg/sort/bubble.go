package sort

func bubble(arr *[]int) {
	var sorted bool

	for !sorted {
		sorted = true
		for idx, val := range *arr {
			if idx != len(*arr)-1 {
				if val > (*arr)[idx+1] {
					sorted = false

					t := (*arr)[idx]
					(*arr)[idx] = (*arr)[idx+1]
					(*arr)[idx+1] = t
				}
			}
		}
	}
}
