package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scanf("%d", &num)

	if num % 2 == 0 {
		fmt.Println("Четное")
	} else {
		fmt.Println("Нечетное")
	}
}
