package main

import "fmt"

type Rectangle struct {
	height, width float64
}

func (a Rectangle) area() float64 {
	return a.height * a.width
}

func main() {
	var a1, b1 float64
	fmt.Print("Введите высоту и ширину через пробел: ")
	fmt.Scanf("%v %v", &a1, &b1)
	a := Rectangle{a1, b1}
	fmt.Printf("Площадь прямоугольника: %v", a.area())
}
