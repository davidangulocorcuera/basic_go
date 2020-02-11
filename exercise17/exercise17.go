package main

import "fmt"

//      * 17 Haz un programa que recorra un array de 5 posiciones y te imprima "hola" por cada iteracción.

func main() {
	var words = [5]string{}
	for i:= 0; i < len(words); i++  {
		fmt.Println("Hola")
	}
}
