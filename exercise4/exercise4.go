package main

//      * 4 Haz un programa que imprima la resta de dos números introducidos por teclado sin decimales

import "fmt"

func main() {
	var num1 int32
	var num2 int32

	fmt.Print("Introduce el primer número ")
	fmt.Scanf("%d", &num1)

	fmt.Print("Introduce el segundo número ")
	fmt.Scanf("%d", &num2)

	fmt.Print(num1 - num2)
}
