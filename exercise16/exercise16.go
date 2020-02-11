package main

//* 16 Haz un programa que te diga de dos números introducidos por teclado cuál es menor

import "fmt"

func main() {
	var num1 int32
	var num2 int32

	fmt.Print("Introduce el primer número ")
	fmt.Scanf("%d", &num1)

	fmt.Print("Introduce el segundo número ")
	fmt.Scanf("%d", &num2)

	if num1 > num2 {
		fmt.Print(num1 ," es mayor")
	} else if num1 == num2{
		fmt.Print(num2 ," y ", num2, " son iguales")
	} else {
		fmt.Print(num2 ," es mayor")
	}
}