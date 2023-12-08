// Текст задания:
// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
package task8

import (
	"fmt"
)

func setBit(num int64, position uint, value uint) int64 {
	// Создаем маску с, представляющую собой последовательность из 63 нулей и последний бит единица.
	// Затем, с помощью побитового сдвига смещаем на нужную позицию.
	// Пример: число "10" в двоичной системе - "1010", мы создаем маску "0001".
	// Хотим заменить второй бит поэтому двигаем маску 2 раза - "0010" - "0100"
	mask := (int64(1) << position)
	// В зависимости от того, что нам нужно поставить "1" или "0", нужно использовать разные побитовые операции
	// Примеры ниже
	if value == 1 {
		num |= mask
	} else {
		num ^= mask
	}
	fmt.Printf("64-bit representation of a number: %064b\n", mask)
	// С помощью побитового OR( | ) ставим значение в исходном числе, в случае, если нам надо установить 0 то используем побитовое XOR ( ^ )
	// В случае с примером: 1010                                                            1010
	//                        |  "ПОБИТОВО ДЕЛИМ"                                             ^
	//                      0100                                                            0010 (В данном случае будем менять первый бит(отсчёт справа и с 0))
	//                      ----                                                            ----
	//         РЕЗУЛЬТАТ:   1110                                               РЕЗУЛЬТАТ:   1000
	// Что в десятичной системе исчислетния = "14"

	return num
}

func Task8() {
	var num int64
	var position uint
	var value uint

	// Вводим данные от пользователя
	fmt.Print("Enter a number (int64): ")
	fmt.Scan(&num)
	fmt.Printf("64-bit representation of a number '%d' = %064b\n", num, num)

	fmt.Print("Enter the bit position (counting from the right, starting with 0): ")
	fmt.Scan(&position)

	fmt.Print("Enter the value (0 or 1): ")
	fmt.Scan(&value)

	// Устанавливаем бит
	result := setBit(num, position, value)

	// Выводим результат
	fmt.Printf("Result after setting the bit: %d\n", result)
	fmt.Printf("64-bit representation of a number '%d' = %064b\n", result, result)
}
