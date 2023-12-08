// Текст задания:
// Реализовать все возможные способы остановки выполнения горутины.
package task6

import (
	"context"
	"fmt"
	"time"
)

// Имитируем работу первой горутины
func workChan(ch chan bool) {
	defer fmt.Println("First goroutine end work")
	fmt.Println("First goroutine working 2 seconds..")
	time.Sleep(2 * time.Second)
	ch <- true
}

// Имитируем работу второй горутины
func workCtx(ctx context.Context) {
	defer fmt.Println("Second goroutine end work")
	fmt.Println("Second goroutine working..")
	time.Sleep(100 * time.Second)
}

func Task6() {
	// Используем канал для остановки горутины workChan
	ch := make(chan bool, 1)

	// Запускаем горутину workChan
	go workChan(ch)

	// Ждем горутину workChan
	res := <-ch
	// Как только придет true - закрываем канал
	if res {
		close(ch)
	}
	// Используем контекст для остановки горутины workCtx
	ctx, cancel := context.WithCancel(context.Background())
	// Запускаем горутину workCtx
	go workCtx(ctx)
	// Ждем 2 секунды и вручную завершаем горутину workCtx
	time.Sleep(2 * time.Second)
	fmt.Println("Cancel second gorutine")
	cancel()
	// Горутина workCtx не успела завершить свою работу, но мы уже ЕЕ саму завершили.
	// Даже defer не помог... жалко ее, конечно...
}
