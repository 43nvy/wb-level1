// Текст задания:
// Реализовать собственную функцию sleep.
//
// Ничего умнее чем создать бесконечный цикл я не придумал...
package task25

import (
	"fmt"
	"time"
)

// Создаем кастомный макет отображения времени, для функции time.Format()
const layoutHHMMSS = "15:04:05"

func mySleep(seconds int) {
	// Смотрим текущее локальное время
	currentTime := time.Now().Local()
	// Добавляем к текущему времени секунды ожидания
	needTime := currentTime.Add(time.Second * time.Duration(seconds))
	// Запускаем бесконечный цикл
	for {
		// На каждой итерации смотрим сколько сейчас времени и сравниваем с нужным
		nowTime := time.Now()
		if nowTime.After(needTime) {
			return
		} else {
			continue
		}
	}
}

func Task25() {
	// Инициализируем переменную и запрашиваем ввод у пользователя
	var seconds int
	fmt.Print("Input seconds for sleep: ")
	fmt.Scan(&seconds)
	fmt.Println("Now time is: ", time.Now().Local().Format(layoutHHMMSS))
	fmt.Printf("Program sleep %d seconds\n", seconds)
	// Спим
	mySleep(seconds)
	// Выводим сообщение о пробуждении
	fmt.Println("Time to up: ", time.Now().Format(layoutHHMMSS))
	for i := 0; i < seconds*2; i++ {
		fmt.Println("Wake up, Neo")
	}
}
