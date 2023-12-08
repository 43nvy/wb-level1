// Текст задания:
// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.
//
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false
package task26

import (
	"fmt"
	"strings"
)

// Почти-что повторяющеесе задание, так как в задании 12 мы уже работали с уникальными множествами.
// Так что, принцип не особо отличается - создаем мапу, которая имеет уникальные ключи в виде рун
// И проверяем каждый символ на повторное вхождение
func strVerification(str string) {
	// Приводим строку к нижнему регистру
	newStr := strings.ToLower(str)

	// Создаем мапу, с ключами rune и значениями bool
	charSet := make(map[rune]bool)
	// Проходим по каждому символу в строке и смотрим, встречался ли он уже
	for _, char := range newStr {
		// Если символ встречался, строка не имеет уникальных символов выводим сообщение и выходим из функции
		if charSet[char] {
			fmt.Printf("String '%s' - not unique\n", str)
			return
		}
		// Иначе, отмечаем символ как встреченный впервые
		charSet[char] = true
	}
	// Если в цикле мы не вышли из функции, значит все символы уникальные.
	fmt.Printf("String '%s' - is unique\n", str)
}

func Task26() {
	// Инициализируем переменну для ввода пользователем и переменную для приветственного сообщения.
	// И служебный слайс со строками из задания
	var str, message string
	exampleSlice := []string{"abcd", "abCdefAaf", "aabcd"}

	message = `
	Need to input string, for unique verification
	For example: strings - 'abcd', 'abCdefAaf', 'aabcd'
	Program will return:`
	fmt.Println(message)
	fmt.Println()
	for _, i := range exampleSlice {
		strVerification(i)
	}

	fmt.Println("\nInput your string: ")
	fmt.Scan(&str)
	// Вызываем функцию
	strVerification(str)
}
