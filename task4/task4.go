// Текст задания:
// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
package task4

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func workerWork(id int, msgChan chan string, callChan chan os.Signal, wg *sync.WaitGroup) {
	// После завершения функции (воркера) убираем из счетчика 1
	defer wg.Done()
	// Читаем сообщения
	for msg := range msgChan {
		fmt.Printf("worker %d read message: %s\n", id, msg)
	}
}

func Task4() {
	var workers int
	fmt.Print("Need workers: ")
	fmt.Scan(&workers)
	// Создаем канал для сообщений
	channel := make(chan string)
	// Создаем канал для сигнала завершения работы
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	// Создаем группу ожидания для воркеров (счетчик)
	var wg sync.WaitGroup

	// Создаем N воркеров
	for i := 1; i <= workers; i++ {
		// Добавляем их в группу
		wg.Add(1)
		// И запускаем воркера
		go workerWork(i, channel, c, &wg)
	}

	// Бесконечно отправляем строки в канал
	for {
		// Используем конструкцию select case в бесконечном цикле
		select {
		case <-c:
			// Ждем сигнал завершения из канала
			// Закрываем канал сообщений, чтобы воркеры завершили свою работу
			close(channel)
			// Ждем завершения всех воркеров
			wg.Wait()
			fmt.Println("Program terminated.")
			// Завершаем работу программы обычным завершением без ошибок (0)
			return
		default:
			// Отправляем сообщение в канал
			channel <- "Hello worker"
			// Пауза для лучшего восприятия
			time.Sleep(time.Millisecond * 75)
		}
	}
}
