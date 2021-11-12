package main

import (
	"fmt"
)

func sorting(items []int32) []int32 {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {

			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
	return items
}

func checkSlice(items []int32, val int32) bool {
	for _, item := range items {
		if item == val {
			return true
		}
	}
	return false
}

type numberDetail struct {
	Number int32
	Total int32
}

func main(){
	var items, once, twice, result []int32
	var number1, number2, number3, number4, number5, number6 int32
	fmt.Print("Input Number 1 : ")
	fmt.Scanln(&number1)
	items = append(items, number1)
	fmt.Print("Input Number 2 : ")
	fmt.Scanln(&number2)
	items = append(items, number2)
	fmt.Print("Input Number 3 : ")
	fmt.Scanln(&number3)
	items = append(items, number3)
	fmt.Print("Input Number 4 : ")
	fmt.Scanln(&number4)
	items = append(items, number4)
	fmt.Print("Input Number 5 : ")
	fmt.Scanln(&number5)
	items = append(items, number5)
	fmt.Print("Input Number 6 : ")
	fmt.Scanln(&number6)
	items = append(items, number6)

	var n = len(items)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if items[i] == items[j] {
				twice = append(twice, items[i])
			}
		}
		if statement := checkSlice(twice, items[i]) ; !statement {
			once = append(once, items[i])
		}
	}

	onceOrdered := sorting(once)
	twiceOrdered := sorting(twice)

	result = append(result, onceOrdered...)

	for _, second := range twiceOrdered {
		for i := 0 ; i < 2 ; i++ {
			result = append(result, second)
		}
	}

	fmt.Println(items)

	fmt.Println(once)
	fmt.Println(twice)
	fmt.Println(result)
}
