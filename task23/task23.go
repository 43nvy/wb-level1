// Текст задания:
// Удалить i-ый элемент из слайса.
package task23

import "fmt"

// По сути эти две функции делают одно и то же.
// Под капотом любого слайса - массив, в этих функциях создается
// новый слайс, который указывает на один и тот же массив, просто на разные индексы
func removeWithAppend(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

// Создаем новый слайс - копию, но исключаем индекс
func removeWithCopy(slice []int, index int) []int {
	result := make([]int, len(slice)-1)
	copy(result[:index], slice[:index])
	copy(result[index:], slice[index+1:])
	return result
}

func Task23() {
	// Создаем исходный слайс и вводимый индекс
	numbers := []int{1, 3, 5, 9, 11, 8, 8, 21, 244, 19, 27}
	var index int
	fmt.Println("Have array: [1, 3, 5, 9, 11, 8, 8, 21, 244, 19, 27]")
	fmt.Print("Choose index: ")
	fmt.Scan(&index)
	// Вызываем первую функцию, удаление с помощью append
	numbers = removeWithAppend(numbers, index)
	fmt.Println("Arr after 'remove with append': ", numbers)
	// Демонстрация второй функции - удаление с помощью copy
	fmt.Print("Choose index again: ")
	fmt.Scan(&index)

	numbers = removeWithCopy(numbers, index)
	fmt.Println("And second remove - 'remove with copy': ", numbers)
}
