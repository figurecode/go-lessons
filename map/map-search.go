package main

import (
	"fmt"
	"math/rand"
)

// Ассоциативные массивы широко применяются при решении алгоритмических задач.
// Когда количество данных более нескольких десятков, поиск значения в map происходит эффективнее, чем в массиве.
// Опираясь на эту информацию, попробуйте решить следующую задачу, которую часто предлагают на собеседованиях.

// Дан массив целых чисел A и целое число k. Нужно найти и вывести индексы пары чисел, сумма которых равна k.
// Если таких чисел нет, то вернуть пустой слайс. Индексы можно вернуть в любом порядке.

func main() {
	r := rand.New(rand.NewSource(100))

	digits := r.Perm(100)

	var sum = make(map[int]int)

	var k int = 23
	var i, v int

	for i, v = range digits {

		for j := 0; j < len(digits); j++ {

			if i == j {
				continue
			}

			if (digits[j] + v) == k {
				sum[i] = j
			}
		}
	}

	fmt.Println(digits)
	fmt.Println(k)
	fmt.Println(sum)

	// Решение с одним циклом от Yandex
    func find(arr []int, k int) []int {
        // Создадим пустую map  
        m := make(map[int]int)
        // будем складывать в неё индексы массива, а в качестве ключей использовать само значение 
        for i, a := range arr {
            if j, ok := m[k-a]; ok { // если значение k-a уже есть в массиве, значит, arr[j] + arr[i] = k и мы нашли, то что нужно
                return []int{i, j}
            }
            // если искомого значения нет, то добавляем текущий индекс и значение в map
            m[a] = i
        }
        // не нашли пары подходящих чисел
        return nil
    }    
    // как можно заметить, алгоритм пройдётся по массиву всего один раз
    // если бы мы искали подходящее значение каждый раз через перебор массива, то пришлось бы сделать гораздо больше вычислений

}
