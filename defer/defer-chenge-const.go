package main

import "fmt"

var Global = 5

func changeConst() {
	defer func(checkout int) {
		Global = checkout // присваиваем Global значение аргумента

		fmt.Println(checkout)
	}(Global) // копируем значение Global в аргументы функции

	Global = 42 // Изменяем Global

	fmt.Println(Global)
	// Здесь будет вызвана отложенная функция
}

func main() {
	fmt.Println(Global)
	changeConst()
	fmt.Println(Global)
}
