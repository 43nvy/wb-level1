// Текст задания:
// Реализовать пересечение двух неупорядоченных множеств.
package task11

import "fmt"

func intersection(set1, set2 []int) []int {
	// Создаем срез длиной 0 (нельзя сделать make для среза, не указав длину), для записи результата
	result := make([]int, 0)

	// Создаем множество для быстрого поиска в set2
	// Тут дело в том, что мапа это хеш-таблица, это все знают
	// но прикол в поиске - у мапы константная (1) сложность на поиск, так как ключи это хеши и они уникальные
	// конечно, существует коллизия, расскажу про это, уже на созвоне) Если спросите)
	// а у массива(слайс под капотом это массив) сложность зависит от количества элементов (n), потому что мы перебираем по ЗНАЧЕНИЮ
	// в случае перебора по ключу(индексу), конечно, тоже константная (1)
	// *тут речь о средних сложностях алгоритмов*
	set2Map := make(map[int]bool)
	for _, element := range set2 {
		set2Map[element] = true
	}

	// Проверяем каждый элемент из set1
	for _, element := range set1 {
		// Если элемент присутствует в set2 - добавляем его в результат
		if set2Map[element] {
			result = append(result, element)
		}
	}

	return result
}

func Task11() {
	// Пример двух неупорядоченных множеств в виде неупорядоченной структуры данных - слайс
	set1 := []int{1, 2, 3, 4}
	set2 := []int{3, 4, 5, 6}

	// Вызываем функцию, которая смотрит на пересечения
	result := intersection(set1, set2)

	// Выводим результат в терминал
	fmt.Println("Intersection of many:", result)
}
