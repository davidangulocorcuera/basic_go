package main

import (
	"fmt"
	"math"
)

//      * 25 Haz un programa que realice esta operación a * b - c (a-d)^2, introduciendo todos los números por teclado.

func main() {
	var a, b, c, d float64

	fmt.Print("Introduce el valor de a ")
	fmt.Scanf("%b", &a)

	fmt.Print("Introduce el valor de b ")
	fmt.Scanf("%b", &b)

	fmt.Print("Introduce el valor de c ")
	fmt.Scanf("%b", &c)

	fmt.Print("Introduce el valor de d ")
	fmt.Scanf("%b", &d)

	result := a*b - c*math.Pow(a-d, 2)

	fmt.Print(" El resultado de a * b - c (a-d)^2 es igual a ", result)
}
