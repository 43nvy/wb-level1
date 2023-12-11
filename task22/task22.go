// Текст задания:
// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
//
// Тип Int64 хранит в себе значения с -9223372036854775808 по 9223372036854775807
// Поэтому, значения привышающие эти числа будут неполные, так как int64 будет отбрасывать цифры в числе
// Я использую пакет math/big, для работы с числами больше, чем int64 может вмещать
package task22

import (
	"fmt"
	"math/big"
)

func Task22() {
	// Инициализируем вводимые значения
	var a, b, powerVal int64

	message := `
	Need to input 'a', 'b' and 'power' values.
	For example: a = 17, b = 29, power = 20: aBig = 17^20 and bBig = 29^20
	`
	fmt.Println(message)
	// Берем значения пользователя и преобразуем в тип big.Int
	fmt.Print("Input 'a' value: ")
	fmt.Scan(&a)
	aBig := big.NewInt(a)

	fmt.Print("Input 'b' value: ")
	fmt.Scan(&b)
	bBig := big.NewInt(b)

	fmt.Print("Input power of values: ")
	fmt.Scan(&powerVal)
	powerBigVal := big.NewInt(powerVal)

	aBig.Exp(aBig, powerBigVal, nil)
	bBig.Exp(bBig, powerBigVal, nil)
	fmt.Printf("Big 'a' value:%s\nBig 'b' value:%s\n", aBig, bBig)
	// Все операции с числами происходят с помощью методов пакета:
	// Сложение
	addResult := new(big.Int).Add(aBig, bBig)
	fmt.Printf("Addition: %s + %s = %s\n", aBig.String(), bBig.String(), addResult.String())
	// Вычитание
	subtractionResult := new(big.Int).Sub(aBig, bBig)
	fmt.Printf("Subtruction: %s - %s = %s\n", aBig.String(), bBig.String(), subtractionResult.String())
	// Умножение
	mulResult := new(big.Int).Mul(aBig, bBig)
	fmt.Printf("Multiplication: %s * %s = %s\n", aBig.String(), bBig.String(), mulResult.String())
	// Деление
	// Так как при работе с больишми числами, у меня часто при делении было 0 целых
	// и неизвестное количество знаков после запятой - я решил поставить 100 знаков после запятой
	divisionResult := new(big.Rat).SetFrac(aBig, bBig)
	fmt.Printf("Division: %s / %s = %s\n", aBig.String(), bBig.String(), divisionResult.FloatString(100))
	// Проверял значения с помощью калькулятора в гугле...
}
