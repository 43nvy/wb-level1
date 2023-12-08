// Текст задания:
// Реализовать конкурентную запись данных в map.
package task7

import (
	"fmt"
	"sync"
	"time"
)

// Создаем новую структуру, которая содержит в себе мапу и мьютекс, для ее синхронизации
type awesomeMap struct {
	data map[string]int
	mu   sync.RWMutex
}

func newAwesomeMap() *awesomeMap {
	return &awesomeMap{
		data: make(map[string]int),
	}
}

// В Golang сообществе рекомендуют обращаться к структурам с помощью геттеров и сеттеров:
func (a *awesomeMap) mapSet(key string, value int) {
	a.data[key] = value
}

func (a *awesomeMap) mapGet(key string) int {
	return a.data[key]
}

// Реализуем конкуренцию за последовательную запись в мапу
func concurrencySetter(newMap *awesomeMap, keys [8]string, values [8]int, wg *sync.WaitGroup, id int) {
	// С помощью defer сразу сообщаем WaitGroup, что горутина закончила и выводим сообщение о завершении работы
	defer wg.Done()
	defer fmt.Printf("Goroutine %d end work\n", id)
	// Пройдемся по всем ключам и значениям
	for i, k := range keys {
		// Залочим мапу и проверим, есть ли такой ключ в ней
		newMap.mu.Lock()
		read := newMap.mapGet(k)
		if read != 0 {
			// Если горутина нашла такой ключ, то отдает поводья другой горутине
			newMap.mu.Unlock()
			continue
		}
		// Для наглядности симулируем hardworking
		time.Sleep(time.Second)
		fmt.Printf("Goroutine %d set value {%d} in key {%s}\n", id, values[i], k)
		newMap.mapSet(k, values[i])
		// Разблокируем мапу для другой горутины
		newMap.mu.Unlock()
	}
}

func Task7() {
	awesomeMap := newAwesomeMap()
	// Добавляем группу ожидания, чтобы main горутина ждала остальные
	var wg sync.WaitGroup
	// Добавляем массивы ключей и значений
	keys := [8]string{"First", "Second", "Third", "Fourth", "Fifth", "Sixth", "Seventh", "Eighth"}
	values := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	// Явно запускаем два паралельных сеттера
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go concurrencySetter(awesomeMap, keys, values, &wg, i)
	}
	// Ждем когда горутины закончат работу
	wg.Wait()
	// Мапа - неупорядоченная структура данных, поэтому вывод не всегда одинаковый
	fmt.Println("Map: ", awesomeMap.data)
}
