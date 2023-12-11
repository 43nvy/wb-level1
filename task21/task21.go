// Текст задания:
// Реализовать паттерн «адаптер» на любом примере.
//
// Паттерн "адаптер" используется для соединения двух несовместимых интерфейсов,
// предоставляя некоторый адаптер между ними
package task21

import "fmt"

// Создаем структуру dog, для которой сделаем метод
type dog struct{}

func (d *dog) woof() {
	fmt.Println("Gav-gav")
}

// Аналагоично, со структурой cat
type cat struct{}

// Немного усложним функцию, передавая дополнительный параметр
// Тем самым, демонстрируем разные методы
func (c *cat) meow(ksKs bool) {
	if ksKs {
		fmt.Println("Meow-meow")
	}
}

// Мы имеем две разные структуры, с разными методами,
// Но перед нами стоит задача удобно ими пользоваться
// Для этого создадим интерфейс, с одним общим методом
type animalAdapter interface {
	reaction()
}

// Структура-адаптер, отличная от dog, и хранящяя в себе ссылку на dog
type dogAdapter struct {
	*dog
}

// Описываем метод reaction() для dogAdapter, которая вызывает метод Woof()
func (a *dogAdapter) reaction() {
	a.woof()
}

// Создаем функцию, которая вернет адаптер
func newDogAdapter(dg *dog) animalAdapter {
	return &dogAdapter{dg}
}

// Аналогично с dogAdapter
type catAdapter struct {
	*cat
}

// Пусть, адаптер поумолчанию вызывает meow(), передавая туда true
func (a *catAdapter) reaction() {
	a.meow(true)
}

// Функция, возвращающая адаптер
func newCatAdapter(ct *cat) animalAdapter {
	return &catAdapter{ct}
}

func Task21() {
	// Создаем массив с адаптерами
	animals := [2]animalAdapter{newDogAdapter(&dog{}), newCatAdapter(&cat{})}
	// Вызываем описанный метод для каждого адаптера
	for _, animal := range animals {
		animal.reaction()
	}
}
