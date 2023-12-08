// Текст задания:
// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
//
// Материал взят с википедии, там был пример реализации на псевдокоде
// Итак, для начала надо выбрать "опорный" элемент, и разбить массив на части
// Затем разбить массив на 2(а иногда на 3) части
// После разбиения, переместить элменты, меньше опорного перед ним, а больше - после опорного
// В каждой части сделать то же самое(рекурсия)
package task16

import "fmt"

// Функция для разбиения
func partition(arr []int, low, high int) ([]int, int) {
	// Выбираем опорный элемент
	pivot := arr[high]
	i := low
	// Проходимся по каждому элменту и сравниваем его
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			// Перемещаем элементы
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	// Возвращаем опорный элемент на свое место
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	// Делаем проверку на то, отсортирован массив, или нет
	// Тоесть, если меньший элемент равен или больше большего, то массив отсортирован
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		// Рекурсивно вызываем функцию, которая проверяет левую часть
		arr = quickSort(arr, low, p-1)
		// А теперь правую часть
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func Task16() {
	arr := []int{55, 11, 5, 6, 7, 2, 1, 0, 10, 27}
	fmt.Println("So arr:", quickSortStart(arr))
}