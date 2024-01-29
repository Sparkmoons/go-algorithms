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

	for i := 0; i < length; i++ {
  min := i

  for j := i+1; j < length; j++ {
			if array[j] < array[min] {
				min = j
			}
		}
  array[i], array[min] = array[min], array[i]
	}

	return array
}