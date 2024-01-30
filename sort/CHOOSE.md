# Сортировка выбором

## Сложность

| Name                  | Best            | Average             | Worst               | Memory    | Stable    | Comments  |
| --------------------- | :-------------: | :-----------------: | :-----------------: | :-------: | :-------: | :-------- |
| **Choose sort**       | n               | n<sup>2</sup>       | n<sup>2</sup>       | 1         | Yes       |           |

## Реализация

```go
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
```