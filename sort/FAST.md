## Быстрая сортировка

Name                  | Best        | Average             | Worst               | Memory    | Stable    | Comments  |
| --------------------- | :---------: | :-----------------: | :-----------------: | :-------: | :-------: | :-------- |
| **Insertion sort**    | n           | n*log(n)       | n*log(n)       | 1         | Yes       |           |


## Реализация

```go
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
```
