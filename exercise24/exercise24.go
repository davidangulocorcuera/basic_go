package main

import (
	"fmt"
	"math"
)

//      * 24 Haz un programa que diga si la raiz cuadrada de n es impar, siendo n un número introducido por teclado.

func main() {
	var number float64
	fmt.Print("Introduce un número ")
	fmt.Scanf("%b", &number)

	if math.Mod(math.Sqrt(number), 2) != 0 {
		fmt.Print("la raiz cuadrada de ", number, " es impar")
	} else {
		fmt.Print("la raiz cuadrada de ", number, " es par")
	}
}