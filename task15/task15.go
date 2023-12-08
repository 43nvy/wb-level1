// Текст задания:
// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
//
// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

//	func main() {
//	  someFunc()
//	}
package task15

import "fmt"

var justString string

// Так как мы не знаем что это за функция, но она просто создает большую строку, исходя из названия
func createHugeString(size int) string {
	// Строка будет содержать "str" и умножится на себя 2^10 раз
	var newString string
	for i := 0; i <= size; i++ {
		newString += "str"
	}
	return newString
}

func someFunc() {
	v := createHugeString(1 << 10)
	// Я пологаю что функция createHugeString не всегда создаст строку
	// достаточной длины, для извлечения среза [:100], поэтому сделаем проверку
	if len(v) > 101 {
		justString = v[:100]
		fmt.Printf("And string: %s\n\n", justString)
	} else {
		fmt.Println("Can't make slice [:100] from 'v' variable")
	}
	// А так же, возможно, задание подразумивает освобождение памяти,
	// ведь под капотом у строки - массив из рун
	// поэтому, сделав строку пустой, мы освободим базовый массив
	v = ""
	justString = ""
	fmt.Printf("Strings 'v' and 'justString': '%s' and '%s'\n", v, justString)
}

func Task15() {
	someFunc()
}
