package sort

import "fmt"

func UseInsertionSort() {
	unsorted := []int{1, 3, 2, 4, 8, 6, 7, 2, 3, 0}
	fmt.Println(unsorted)
	sorted := MergeSort(unsorted)
	fmt.Println(sorted)
}

func MergeSort(array []int) []int {
    if len(array) <= 1 {
        return array
    }
    var left, right []int
    mid := len(array) / 2
    left = MergeSort(array[:mid])
    right = MergeSort(array[mid:])
    
    return Merge(left, right)
}

func Merge(left, right []int) []int {
    res := make([]int, 0, len(left) + len(right))
    for len(left) > 0 || len(right) > 0 {
        if len(left) == 0 {
            res = append(res, right...)
        } else if len(right) == 0 {
            res = append(res, left...)
        } else if left[0] < right[0] {
            res = append(res, left[0])
            left = left[1:]
        } else {
            res = append(res, right[0])
            right = right[1:]
        }
    }
    return res
}