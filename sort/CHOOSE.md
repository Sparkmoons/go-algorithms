:)
| Name                  | Best            | Average             | Worst               | Memory    | Stable    | Comments  |
| --------------------- | :-------------: | :-----------------: | :-----------------: | :-------: | :-------: | :-------- |
| **Bubble sort**       | n               | n<sup>2</sup>       | n<sup>2</sup>       | 1         | Yes       |           |

## Реализация

```go
func BubbleSort(array []int) []int {
	changed := true
	length := len(array)

	// Массив в алгоритме считается отсортированным. При первой замене доказывается обратное и запускается еще одна итерация.
	// Цикл останавливается, когда все пары элементов в массиве пропускаются без замен
	for changed {
		changed = false

		for i := 1; i < length; i++ {
			// если порядок не верный
			if array[i-1] > array[i] {
				// меняем местами элементы
				array[i], array[i-1] = array[i-1], array[i]

				// продолжаем сортировку
				changed = true
			}
		}
	}

	return array
}
```