## Сортировка слиянием

## Сложность

|Name                   | Best        | Average             | Worst               | Memory    | Stable    | Comments  |
| --------------------- | :---------: | :-----------------: | :-----------------: | :-------: | :-------: | :-------- |
| **Merge sort**        | n           | n*log(n)            | n*log(n)            | 1         | Yes       |           |



## Реализация

```go
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
```
