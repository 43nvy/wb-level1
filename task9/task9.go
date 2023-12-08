// Текст задания:
// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.
package task9

import (
	"fmt"
	"sync"
	"time"
)

// Функция для горутины, которая пишет из слайса в канал
func arrWriter(arr []int, inputChan chan int, wg *sync.WaitGroup) {
	// Выводим в консоль что горутина закончила работу, затем закрываем канал и сообщаем всей группе ожидания, что одна из горутин завершила работу.
	defer wg.Done()
	defer close(inputChan)
	defer fmt.Println("Gorutine 'arrWriter' end work")

	for _, num := range arr {
		// Обычный перебор слайса, со слипом, для демонстрации
		time.Sleep(75 * time.Millisecond)
		inputChan <- num
	}
}

// Функция для горутины, которая читает из канала, удваивает значение и передает в другой канал
func arrDoubler(inputChan chan int, outputChan chan int, wg *sync.WaitGroup) {
	// Аналогичные, предыдущей горутине, действия, только здесь закрываем канал который передает удвоенные значения
	defer wg.Done()
	defer close(outputChan)
	defer fmt.Println("Gorutine 'arrDoubler' end work")

	for num := range inputChan {
		// Получаем из канала первоначальные значения, затем удваиваем и отправляем в другой канал
		time.Sleep(75 * time.Millisecond)
		double := num * 2
		outputChan <- double
	}
}

func arrReader(outputChan chan int, wg *sync.WaitGroup) {
	// Выводим сообщение в консоль и сообщяем группе что горутина завершила работу
	defer wg.Done()
	defer fmt.Println("Gorutine 'arrReader' end work")

	for double := range outputChan {
		// Просто читаем канал и выводим удвоенные значения в терминал
		fmt.Printf("Gorutine 'arrReader' take %d\n", double)
	}
}

func Task9() {
	// Создаем два канала
	inputChan := make(chan int)
	outputChan := make(chan int)
	// Инициализируем группу ожидания
	var wg sync.WaitGroup
	// Вот наш слайс значений
	numbers := []int{18, 12, 33, 40, 95, 66, 27, 18, 89, 10}
	// Сообщаем группе ожидания, что у нас появилось 3 горутины и запускаем их
	wg.Add(1)
	go arrWriter(numbers, inputChan, &wg)
	wg.Add(1)
	go arrDoubler(inputChan, outputChan, &wg)
	wg.Add(1)
	go arrReader(outputChan, &wg)
	// Ждем, когда все горутины отметятся, что закончили
	wg.Wait()
}
