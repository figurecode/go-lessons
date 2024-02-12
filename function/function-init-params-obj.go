package main

import "fmt"

// @see https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html

type Item struct {
	NoOption   string
	Parameter1 string
	Parameter2 int
}

func NewItem(opts ...option) *Item {
	// Инициализируем типовыми значениями
	i := &Item{
		NoOption:   "usual",
		Parameter1: "default",
		Parameter2: 42,
	}

	// применяем опции в том порядке, в котором они были заявлены
	for _, opt := range opts {
		opt(i)
	}

	return i
}

// Здесь опции — это функции, применяемые к объекту.
// За это подход получил название funcopts.
type option func(*Item)

// Чтобы устанавливать параметры, будем использовать функции
// высшего порядка, возвращающие значениями функции `option`

func Option1(option1 string) option {
	return func(i *Item) {
		i.Parameter1 = option1
	}
}

func Option2(option2 int) option {
	return func(i *Item) {
		i.Parameter2 = option2
	}
}

// Тогда инициализация объекта конструктором будет выглядеть так:

func main() {
	itemDef := NewItem()

	fmt.Println(itemDef)

	itemOption1 := NewItem(Option1("unusual"))

	fmt.Println(itemOption1)

	itemOptionsSet := NewItem(Option1("test"), Option2(404))

	fmt.Println(itemOptionsSet)
}
