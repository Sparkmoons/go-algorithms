package sort

import "fmt"

func UseChooseSort() {
	unsorted := []int{1, 3, 2, 4, 8, 6, 7, 2, 3, 0}
	fmt.Println(unsorted)
	sorted := ChooseSort(unsorted)
	fmt.Println(sorted)
}

func ChooseSort(array []int) []int {
	length := len(array)

	for i := 1; i < length; i++ {
		j := i

		for j > 0 {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
			j = j - 1
		}
	}

	return array
}