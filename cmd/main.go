package main

import (
	"fmt"

	"github.com/43nvy/wb-level1/task1"
	"github.com/43nvy/wb-level1/task10"
	"github.com/43nvy/wb-level1/task11"
	"github.com/43nvy/wb-level1/task12"
	"github.com/43nvy/wb-level1/task13"
	"github.com/43nvy/wb-level1/task14"
	"github.com/43nvy/wb-level1/task15"
	"github.com/43nvy/wb-level1/task16"
	"github.com/43nvy/wb-level1/task17"
	"github.com/43nvy/wb-level1/task18"
	"github.com/43nvy/wb-level1/task19"
	"github.com/43nvy/wb-level1/task2"
	"github.com/43nvy/wb-level1/task20"
	"github.com/43nvy/wb-level1/task21"
	"github.com/43nvy/wb-level1/task22"
	"github.com/43nvy/wb-level1/task23"
	"github.com/43nvy/wb-level1/task24"
	"github.com/43nvy/wb-level1/task25"
	"github.com/43nvy/wb-level1/task26"
	"github.com/43nvy/wb-level1/task3"
	"github.com/43nvy/wb-level1/task4"
	"github.com/43nvy/wb-level1/task5"
	"github.com/43nvy/wb-level1/task6"
	"github.com/43nvy/wb-level1/task7"
	"github.com/43nvy/wb-level1/task8"
	"github.com/43nvy/wb-level1/task9"
)

func main() {
	var taskNumber int

	for {
		fmt.Print("\nInput task number [1 to 26]: ")
		fmt.Scan(&taskNumber)

		if taskNumber < 1 || taskNumber > 26 {
			fmt.Print("No, between [1 to 26]")
		}

		switch taskNumber {
		case 1:
			task1.Task1()
		case 2:
			task2.Task2()
		case 3:
			task3.Task3()
		case 4:
			task4.Task4()
		case 5:
			task5.Task5()
		case 6:
			task6.Task6()
		case 7:
			task7.Task7()
		case 8:
			task8.Task8()
		case 9:
			task9.Task9()
		case 10:
			task10.Task10()
		case 11:
			task11.Task11()
		case 12:
			task12.Task12()
		case 13:
			task13.Task13()
		case 14:
			task14.Task14()
		case 15:
			task15.Task15()
		case 16:
			task16.Task16()
		case 17:
			task17.Task17()
		case 18:
			task18.Task18()
		case 19:
			task19.Task19()
		case 20:
			task20.Task20()
		case 21:
			task21.Task21()
		case 22:
			task22.Task22()
		case 23:
			task23.Task23()
		case 24:
			task24.Task24()
		case 25:
			task25.Task25()
		case 26:
			task26.Task26()
		}
	}
}
