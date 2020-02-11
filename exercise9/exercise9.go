package main
//     * 9 Haz un programa que imprima el resultado de a * b + c, siendo a b y c números introducidos por teclado

import "fmt"

func main() {
	var num1 float64
	var num2 float64
	var num3 float64

	fmt.Print("Introduce el primer número ")
	fmt.Scanf("%b", &num1)

	fmt.Print("Introduce el segundo número ")
	fmt.Scanf("%b", &num2)

	fmt.Print("Introduce el tercer número ")
	fmt.Scanf("%b", &num3)

	fmt.Print(num1 * num2 + num3)

}
