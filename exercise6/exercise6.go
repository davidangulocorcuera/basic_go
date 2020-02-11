package main

//     * 6 Haz un programa que imprima la división de dos números introducidos por teclado con decimales.

import "fmt"

func main() {
	var num1 float64
	var num2 float64

	fmt.Print("Introduce el primer número ")
	fmt.Scanf("%b", &num1)

	fmt.Print("Introduce el segundo número ")
	fmt.Scanf("%b", &num2)

	fmt.Print(num1 / num2)
}
