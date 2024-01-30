package sort

import "fmt"

func UseBubbleSort() {
	unsorted := []int{1, 3, 2, 4, 8, 6, 7, 2, 3, 0}
	fmt.Println(unsorted)
	sorted := BubbleSort(unsorted)
	fmt.Println(sorted)
}

func FastSort(array []int) []int {
    if len(array) <= 1 {
        return array
    }
    var min, max, res []int
    p := array[0]
    for _, num := range array[1:] {
        if num <= p {
            min = append(min, num)
        } else {
            max = append(max, num)
        }
    }
    res = append(FastSort(min), p)
    res = append(res, FastSort(max)...)
    return res
}
