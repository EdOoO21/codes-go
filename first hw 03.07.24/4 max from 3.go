package main

import (
	"fmt"
)

func main() {
	var a, b, c, maxnum int
	fmt.Print("Введите 3 числа через пробел: ")
	fmt.Scanf("%d %d %d", &a, &b, &c)

	maxnum = max(a, b)
	maxnum = max(maxnum, c)
	fmt.Printf("Максимальное из них: %d", maxnum)
}
