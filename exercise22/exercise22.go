package main

import (
	"fmt"
	"math"
)

//     * 22 Haz un programa que diga si el cuadrado de n es impar, siendo n un número introducido por teclado.

func main() {
	var number float64
	fmt.Print("Introduce un número ")
	fmt.Scanf("%b", &number)

	if math.Mod(math.Pow(number, 2), 2) != 0 {
		fmt.Print("el cuadrado de ", number, " es impar")
	} else {
		fmt.Print("el cuadrado de ", number, " es par")
	}
}
