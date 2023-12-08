// Текст задания:
// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
package task24

import (
	"fmt"
	"math"
)

// Создаем структуру, представляющая точку в двумерном пространстве
// где x, и y координаты точки
type pointStruct struct {
	x, y int
}

// Создаем конструктор, который возвращается структуру точки - Point
func newPoint(x, y int) pointStruct {
	return pointStruct{x, y}
}

// Функция для определения расстояния
func distanceBetweenPoints(p1, p2 pointStruct) float64 {
	// Расстояние между двумя точками равно квадратному корню из суммы квадратов разностей координат по каждой оси.
	// Поэтому вычисляем разности координат по осям
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	// А затем вычисляем квадратный корень( Sqrt() ) суммы квадратов(dx*dx, dy*dy) разностей
	// Так как метод Sqrt() работает с float64 - преобразуем наши int
	return math.Sqrt(float64(dx*dx) + float64(dy*dy))
}

func Task24() {
	var point1x, point1y, point2x, point2y int
	message := `
	Need to input 4 values: fisrt point x, y and second point x, y
	For example: point1 (1, 2), point2 (4, 6)`
	fmt.Println(message)
	// Первая точка:
	fmt.Print("\nFirst point x: ")
	fmt.Scan(&point1x)
	fmt.Print("First point y: ")
	fmt.Scan(&point1y)
	// Вторая точка:
	fmt.Print("\nSecond point x: ")
	fmt.Scan(&point2x)
	fmt.Print("Second point y: ")
	fmt.Scan(&point2y)
	// Создаем две точки
	point1 := newPoint(point1x, point1y)
	point2 := newPoint(point2x, point2y)
	// Вызываем функцию рассчета расстояния между точками
	distance := distanceBetweenPoints(point1, point2)
	// Выводим результат в терминал
	fmt.Printf("\nDistance: %f\n", distance)
}
