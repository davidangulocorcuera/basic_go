package main

//      * 8 Haz un programa que imprima el cuadrado de un número introducido por teclado

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Print("Introduce un número ")
	fmt.Scanf("%b", &num)

	fmt.Print(math.Pow(num,2))
}
