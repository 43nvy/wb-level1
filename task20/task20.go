// Текст задания:
// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».
package task20

import (
	"fmt"
	"strings"
)

func Task20() {
	// Исходная строка со словами
	str := "snow dog sun"
	fmt.Printf("Words string: %s\n", str)
	// Создаем срез из слов, с помощью Split() из пакета strings, по разделителю - " "
	words := strings.Split(str, " ")
	// Меняем местами в цикле
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	// С помощью Join() обьеденяем срез в строку с разделителем: " "
	fmt.Printf("Reversed string: %s\n", strings.Join(words, " "))
}
