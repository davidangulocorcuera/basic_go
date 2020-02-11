package main

import "fmt"

//      * 27 Haz un programa que obtenga el último caracter de un texto y lo imprima.

func main() {
	var text string
	fmt.Print("Introduce un texto ")
	fmt.Scanf("%s", &text)

	fmt.Print("El ultimo caracter de ", text, " es ", text[len(text)-1:])
}
