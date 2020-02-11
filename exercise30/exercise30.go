package main

import "fmt"

//     * 30 haz un programa que introduzcas un número por teclado y te diga si es primo o no.

func main() {
	var num int
	var divisores = 0

	fmt.Print("Introduce un número ")
	fmt.Scanf("%d", &num)

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			divisores++
		}
	}

	if divisores == 2 {
		println("es primo")
	} else {
		println("no es primo")
	}

}
