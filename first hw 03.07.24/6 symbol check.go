package main

import (
	"fmt"
)

func main() {
	var symbol string
	fmt.Print("Введите 1 символ латинской раскладки или любого знака: ")
	fmt.Scanf("%v", &symbol)
	if len(symbol) > 1 {
		fmt.Println("Я же просил !!!!! :-(")
	} else {
		if (int(symbol[0]) == 65) || (int(symbol[0]) == 97) || (int(symbol[0]) == 101) || (int(symbol[0]) == 69) || (int(symbol[0]) == 105) || (int(symbol[0]) == 73) || (int(symbol[0]) == 101) || (int(symbol[0]) == 69) || (int(symbol[0]) == 111) || (int(symbol[0]) == 79) || (int(symbol[0]) == 117) || (int(symbol[0]) == 85) {
			fmt.Println("Латинская гласная буква")
		} else {
			fmt.Println("Другое")
		}
	}
}
