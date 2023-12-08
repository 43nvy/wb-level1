// Текст задания:
// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.
package task14

import (
	"fmt"
	"reflect"
)

func getType(variable interface{}) string {
	// Используем reflect.TypeOf для получения информации о типе переменной
	// Рефлексия это как раз и есть изменение своего кода, во время выполнения, тоесть в рантайме
	varType := reflect.TypeOf(variable)

	return varType.String()
}

func Task14() {
	// Пример переменных различных типов
	var integerVar int = 42
	var stringVar string = "Hello, Golang!"
	var boolVar bool = true
	var channelVar chan int = make(chan int)

	// Пример переменной типа interface{}
	var anyVar interface{} = "Variable"

	// Используем функцию, для определения типов переменных в райнтайме
	intType := getType(integerVar)
	stringType := getType(stringVar)
	boolType := getType(boolVar)
	channelType := getType(channelVar)
	anyType := getType(anyVar)

	// Вывод типов переменных в терминал
	fmt.Printf("Type of integerVar: %s\n", intType)
	fmt.Printf("Type of stringVar: %s\n", stringType)
	fmt.Printf("Type of boolVar: %s\n", boolType)
	fmt.Printf("Type of channelVar: %s\n", channelType)
	fmt.Printf("Type of anyVar: %s\n", anyType)
}
