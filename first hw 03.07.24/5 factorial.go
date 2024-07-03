package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scanf("%d", &num)
	ans := 1
	for i := 1; i <= num; i++ {
		ans *= i
	}
	fmt.Printf("%d! = %d", num, ans)
}
