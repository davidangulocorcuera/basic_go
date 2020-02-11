package main

import "fmt"

//      * 18 Haz un programa que recorra un array de 9 posiciones y te imprima
//     el valor de cada posición, el array debe contener los números del 1 al 9.

func main() {
	var numbers = [9]int{}
	for i := 0; i < len(numbers); i++ {
		numbers[i] = i + 1
		fmt.Print(numbers[i])
	}

}
