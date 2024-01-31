## Линейный поиск

## Реализация

``` go
func LinSearch(array []int, val int) int{
    for i, num := range array {
        if num == val {
            return i
        }
    }
    return -1    
}
```