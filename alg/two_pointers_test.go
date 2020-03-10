package alg

import "testing"

func TestTwoPointersAlg(t *testing.T) {
	arr := []int{160, 165, 170, 170, 173, 175, 176, 180, 183, 190, 191, 192, 193}

	if val := TwoPointersAlg(arr, 10); val != 6 {
		t.Errorf("TwoPointersAlg failed with params %v %d and result %d expected %d", arr, 10, val, 6)
	} // 6 (можно выбрать 170, 170, 173, 175, 176, 180)

	if val := TwoPointersAlg(arr, 0); val != 2  {
		t.Errorf("TwoPointersAlg failed with params %v %d and result %d expected %d", arr, 0, val, 2)
	} // 2 (можно выбрать 170, 170)

	if val := TwoPointersAlg(arr, 3); val != 4  {
		t.Errorf("TwoPointersAlg failed with params %v %d and result %d expected %d", arr, 3, val, 4)
	} // 4 (можно выбрать 190, 191, 192, 193)

	if val := TwoPointersAlg(arr, 20); val != 9 {
		t.Errorf("TwoPointersAlg failed with params %v %d and result %d expected %d", arr, 20, val, 9)
	} // 9 (можно выбрать всех начиная с 160 и заканчивая 180 - тогда ответ был бы 8, но если взять с 173 по 193 - то ответ будет 9)
}
