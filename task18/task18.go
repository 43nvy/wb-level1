// Текст задания:
// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.
package task18

import (
	"fmt"
	"sync"
)

// Создаем структуру, включающую в себя сам счетчик и мьютекс
type counter struct {
	count int
	mu    sync.Mutex
}

// Функция, для получения ссылки на счётчик
func newCounter() *counter {
	return &counter{}
}

// Функция, для увеличение счетчика, с использованием мьютекса
// Что не позволяет всем горутинам одновременно заполнить счетчик
func (c *counter) increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// Функция, реализующая параллелизм, для наглядности выводит сообщение
func concurrency(id int, c *counter, wg *sync.WaitGroup) {
	// Показываем группе ожидания, что горутина закончила работу
	defer wg.Done()
	c.increment()
	fmt.Printf("Gorutine %d add count\n", id)
}

func Task18() {
	// Инициализируем счетчик
	c := newCounter()
	// Создаем группу ожидания, чтобы мейн-горутина ждала остальных
	var wg sync.WaitGroup
	// Создаем 100 горутин
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go concurrency(i, c, &wg)
	}
	// Ждем все горутины и выводим результат счетчика
	wg.Wait()
	fmt.Println("Count: ", c.count)
}
