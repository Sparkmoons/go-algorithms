package sort

import "fmt"

func UseLinSearch(){
    n := []int{1, 6, 4, 7, 10, 3, 5}
    fmt.Println(LinSearch(n, 6))
}

func LinSearch(array []int, val int) int{
    for i, num := range array {
        if num == val {
            return i
        }
    }
    return -1    
}
