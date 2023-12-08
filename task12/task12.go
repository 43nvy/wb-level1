// Текст задания:
// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
package task12

import "fmt"

func Task12() {
	// Создаем слайс строк из задания
	stringsArr := []string{"cat", "cat", "dog", "cat", "tree"}
	// В виде множества будет выступать неупорядоченная структура - мапа
	// Так как ключи в мапе уникальные - мы избавляемся от дубликатов
	// И получаем множество с уникальными значениями
	stringsMap := make(map[string]bool)
	// Перебираем слайс и добавляем ключи в мапу
	for _, value := range stringsArr {
		stringsMap[value] = true
	}

	fmt.Println("Intersection: ", stringsMap)
}
