// Текст задания:
// Поменять местами два числа без создания временной переменной.
package task13

import "fmt"

func Task13() {
	var firstNumber int
	var secondNumber int

	fmt.Print("First number: ")
	fmt.Scan(&firstNumber)
	fmt.Print("Second number: ")
	fmt.Scan(&secondNumber)
	// Не баг а фича, в го-туре было)
	// Подозреваю что тут работают ссылки
	firstNumber, secondNumber = secondNumber, firstNumber
	fmt.Printf("And first: %d, second: %d\n", firstNumber, secondNumber)
	// Еще мне в голову пришли математические операции:
	// Так как переменные могут производить операции с самими собой
	// Просто создавая копию значения (не переменной)
	// Ведь в Go, все работают с копиями значений.
	firstNumber = firstNumber + secondNumber
	secondNumber = firstNumber - secondNumber
	firstNumber = firstNumber - secondNumber

	fmt.Printf("Again first: %d, second: %d\n", firstNumber, secondNumber)
}
