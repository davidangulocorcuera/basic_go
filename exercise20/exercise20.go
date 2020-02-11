package main

import (
	"fmt"
	"math"
)

//     * 20 Haz un programa que calcule el arcotangente de n , siendo n un número introducido por teclado.

func main() {
	var number float64
	fmt.Print("Introduce un número ")
	fmt.Scanf("%b", &number)

	fmt.Print(math.Atan(number))
}
