package main

//     * 14 Haz un programa que te diga si el numero introducido por teclado es impar.

import "fmt"

func main() {
	var number int32
	fmt.Print("Introduce un número ")
	fmt.Scanf("%d",&number)

	if number % 2 != 0 {
		fmt.Print("El número introducido es impar ")
	} else if number == 0 {
		fmt.Print("El número introducido es cero ")
	} else {
		fmt.Print("El número introducido es par ")
	}
}