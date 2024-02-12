package main

import (
	"fmt"
	"math/rand"
	"time"
)

func metricTime(start time.Time) {
	fmt.Println(time.Now().Sub(start))
}

func VeryLongTimeFunction() {
	defer metricTime(time.Now())

	var items []int

	r := rand.New(rand.NewSource(100))

	items = r.Perm(100)

	for i := 0; i < len(items); i++ {
		fmt.Printf("%d: %d \n", items[i], items[i]*items[i])
	}
}

func main() {
	VeryLongTimeFunction()
}
