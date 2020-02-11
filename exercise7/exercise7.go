package main

//     * 7 Haz un programa que imprima la raíz cuadrada de un número introducido por teclado

import (
	"fmt"
	"math"
	)

func main() {
	var num float64

	fmt.Print("Introduce un número ")
	fmt.Scanf("%b", &num)
	fmt.Print(math.Sqrt(num))
}
