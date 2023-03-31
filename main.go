package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Print("Введите длину первой стороны треугольника, см: ")
	fmt.Scanf("%f", &a)
	fmt.Print("Введите длину второй стороны треугольника, см: ")
	fmt.Scanf("%f", &b)
	fmt.Print("Введите длину третьей стороны треугольника, см: ")
	fmt.Scanf("%f", &c)
	s := (a + b + c) / 2
	area := math.Sqrt(s * (s - a) * (s - b) * (s - c))
	fmt.Printf("Площадь треугольника, см2: %.2f\n", area)
}
