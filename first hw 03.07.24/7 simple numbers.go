package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Введите натуральное число: ")
	fmt.Scanf("%d", &num)
	sieve := make([]bool, num+1, num+1)

	for i := 2; i <= num; i++ {
		if !sieve[i] {
			for j := i; j*i <= num; j++ {
				sieve[j*i] = true
			}
		}
	}
	for i := 2; i <= num; i++ {
		if !sieve[i] {
			fmt.Printf("%d ", i)
		}
	}

}
