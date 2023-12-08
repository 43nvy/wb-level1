// Текст задания:
// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
//
// Этот пакет представляет composition structures, это реализация наследования в Golang - встраивание структур.
package task1

import "fmt"

// Структура "Human" описывает человека
type Human struct {
	name    string
	age     int
	actions Action
}

// Структура "Action" описывает некоторые действия
type Action struct{}

func (a *Action) Speak() {
	fmt.Println("Speaking something")
}

func (a *Action) Walk() {
	fmt.Println("Walking somewhere")
}

// В Go не рекомендуется обращаться напрямую к полям структуры, для этого используются
// Getters and Setters(геттеры и сеттеры)
func (h *Human) SetName(setName string) {
	h.name = setName
}

func (h *Human) SetAge(setAge int) {
	h.age = setAge
}

func (h *Human) GetName() string {
	return h.name
}

func (h *Human) GetAge() int {
	return h.age
}

func Task1() {
	// Создаем объект Human
	person := Human{}
	// И устанавливаем значения полей
	person.SetName("Andrey")
	person.SetAge(22)
	fmt.Println(person.GetName(), person.GetAge())
	// Вызываем методы Speak и Walk как если бы они были частью Human
	person.actions.Speak()
	person.actions.Walk()
}
