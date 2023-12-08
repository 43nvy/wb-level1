// Текст задания:
// Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.
package task5

import (
	"fmt"
	"time"
)

// Создаем функцию(горутину) для чтения канала
func readMessage(channel <-chan string) {
	for value := range channel {
		fmt.Printf("Read: '%s' message\n", value)
	}
	fmt.Println("Exiting reader.")
}

func Task5() {
	var sec int
	fmt.Print("Seconds: ")
	fmt.Scan(&sec)
	// Создаем канал
	channel := make(chan string)

	// Запускаем горутину для чтения из канала
	go readMessage(channel)

	// Создаем таймер на N секунд
	runTime := time.Duration(sec) * time.Second
	timer := time.NewTimer(runTime)

	// Отправляем значения в бесконечном цикле
	for {
		// С помощью конструкции select case делаем проверку на закрытие канала
		select {
		case <-timer.C:
			// Закрываем канал после N секунд
			close(channel)
			fmt.Println("Channel closed.")
			// Завершаем работу программы обычным завершением без ошибок
			fmt.Println("Program exited.")
			return
		default:
			channel <- "Hello :)"
			// Пауза для наглядности
			time.Sleep(time.Millisecond * 75)
		}
	}
}
