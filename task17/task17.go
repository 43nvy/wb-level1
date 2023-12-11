// Текст задания:
// Реализовать бинарный поиск встроенными методами языка.
//
// Бинарная сортировка - распространненый алгоритм сортировки
// Его суть заключается в том, что мы делим исходный массив на два
// А затем проверяем если ключ меньше значения середины, то ищем в первой половине
// Иначе - во второй. Так как это цикл, то мы будем повторять эти действия,
// пока не найдем нужное значение, или интервал поиска не станет пустым

package task17

import "fmt"

func binarySearch(arr [10]int, value int) (int, bool) {
	// Определяем конечный и начальный индексы
	high := len(arr) - 1
	low := 0
	for low <= high {
		// Определяем индекс середины массива
		mid := (low + high) / 2
		// Узнаем значение середины
		midValue := arr[mid]
		// Делаем проверку, является ли значение серединного элемента искомым
		if midValue == value {
			// Если да, то выходим из функции и возвращаем значение и статус
			return mid, true
			// Если нет, то проверяем значение среднего элемента больше или меньше искомого
		} else if midValue > value {
			// Если больше - то переопределяем конечный индекс и ищем в левой части
			high = mid - 1
		} else {
			// Если меньше - то в правой
			low = mid + 1
		}
	}
	// Если цикл закончился, а искомое не найдено - возвращаем 0 и false
	return 0, false
}

func Task17() {
	// Инициализируем массив и искомое значение
	arr := [10]int{1, 2, 4, 5, 9, 11, 19, 28, 29, 33}
	var value int
	fmt.Println("Have arr [1, 2, 4, 5, 9, 11, 19, 28, 29, 33]")
	fmt.Print("Input value for search: ")
	fmt.Scan(&value)

	key, ok := binarySearch(arr, value)
	if !ok {
		fmt.Printf("Value '%d' not found\n", value)
		return
	}

	fmt.Printf("Value: '%d', key: '%d'\n", value, key)
}
