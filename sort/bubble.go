package sort

import "fmt"

func UseBubbleSort() {
	unsorted := []int{1, 3, 2, 4, 8, 6, 7, 2, 3, 0}
	fmt.Println(unsorted)
	sorted := BubbleSort(unsorted)
	fmt.Println(sorted)
}

func BubbleSort(array []int) []int {
	changed := true
	length := len(array)

	// пока есть изменения
	for changed {
		changed = false

		for i := 1; i < length; i++ {
			// если порядок не верный
			if array[i-1] > array[i] {
				// меняем местами элементы
				array[i], array[i-1] = array[i-1], array[i]

				// продолжаем сортировку
				changed = true
			}
		}
	}

	return array
}
