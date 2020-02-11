package main

import "fmt"

//     * 26 Haz un programa que obtenga el primer caracter de un texto y lo imprima.

func main() {
	var text string
	fmt.Print("Introduce un texto ")
	fmt.Scanf("%s", &text)

	fmt.Print("El primer caracter de ", text, " es ", text[0:1])
}
