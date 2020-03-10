package alg

import "testing"

func TestCount(t *testing.T) {
	arr := []int{1, 4, 8, 11, 21, 23, 24, 64, 70, 84, 125, 704, 1099, 1200, 1201, 1202, 2000}

	if elms := Count(arr, 1, 10); elms != 3 {
		t.Errorf("Count failed: Count( arr, %d, %d)", 1, 10)
	} // 3   (подходят элементы 1, 4, 8)

	if elms := Count(arr, 200, 700); elms != 0 {
		t.Errorf("Count failed: Count( arr, %d, %d)", 200, 700)
	} // 0   (в этом промежутке нет искомых чисел)

	if elms := Count(arr, 20, 100); elms != 6 {
		t.Errorf("Count failed: Count( arr, %d, %d)", 20, 100)
	} // 6   (подходят элементы 21,23,24,64,70,84)

	if elms := Count(arr, 1000, 2000); elms != 5 {
		t.Errorf("Count failed: Count( arr, %d, %d)", 1000, 2000)
	} // 5   (подходят элементы 1099, 1200, 1201, 1202, 2000)

	if elms := Count(arr, 84, 84); elms != 1 {
		t.Errorf("Count failed: Count( arr, %d, %d)", 84, 84)
	} // 1   (подходит элемент 84)
}
