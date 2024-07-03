package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string

	fmt.Print("Введите строку: ")
	fmt.Scanf("%v", &str1)
	var size int = len(str1)
	str := strings.Split(str1, "")
	for i := 0; i < size/2; i++ {
		str[size-1-i], str[i] = str[i], str[size-1-i]
	}

	for _, i := range str {
		fmt.Printf("%v ", i)
	}
}
