package sort

import "fmt"

func UseBinSearch(){
    n := []int{1, 6, 4, 7, 10, 3, 5}
    fmt.Println(BinSearch(n, 6))
}

func BinSearch(array []int, val int) []int{
    if len(array) == 0 {
        return -1
    }
    first, right := 0, len(array) - 1
    for first <= last {
        mid := ((last - first) / 2) + first
        if array[mid] == val {
            return mis
        } else if array[mid] > val {
            last = mid - 1
        } else if array[mid] < val {
            first = first + 1
        }
    }
    return -1
}
