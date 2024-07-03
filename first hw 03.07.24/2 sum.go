package main

import (
	"fmt"
)

func sum(a, b int) int { return a + b }

func main() {
	var a, b int
	fmt.Print("Enter a: ")
	fmt.Scanf("%d", &a)
	fmt.Print("Enter b: ")
	fmt.Scanf("%d", &b)
	fmt.Println(sum(a, b))
}
