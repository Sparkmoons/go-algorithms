package sort

import "fmt"

func UseInsertionSort() {
	unsorted := []int{1, 3, 2, 4, 8, 6, 7, 2, 3, 0}
	fmt.Println(unsorted)
	sorted := InsertionSort(unsorted)
	fmt.Println(sorted)

	
	//var2
	unsorted2 := []int{1, 3, 2, 4, 8, 6, 7, 2, 3, 0}
	fmt.Println(unsorted2)
	i := 1
	for i > len(unsorted2) {
		j := i
		for j >= 1 && n[j] < n[j-1] {
			n[j], n[j-1] = n[j-1], n[j]
			j--
		}
		i++
	}
	fmt.Println(unsorted2)
}

func InsertionSort(array []int) []int {
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
