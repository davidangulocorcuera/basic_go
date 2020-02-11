package main

import "fmt"

//      * 29 Haz un programa que sume 5 números introducidos por teclado y te diga si el resultado es par o impar.
func main() {
	var result = 0
	var number int

	for i := 1; i <= 5; i++ {
		fmt.Print("Introduce un número ")
		fmt.Scanf("%d", &number)
		result += number
	}

	if result%2 == 0 {
		fmt.Print(result, " es par")
	} else {
		fmt.Print(result, " es impar")
	}

}
