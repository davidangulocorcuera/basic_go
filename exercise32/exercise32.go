package main

import "fmt"

//* 33 Haz un programa que compruebe que un texto introducido por teclado empieza por "in"

func main() {
	var text string
	fmt.Print("Introduce un texto ")
	fmt.Scanf("%s", &text)

	var first = text[0:2]

	if first == "in" {
		fmt.Print("si empieza en in")
	} else {
		fmt.Print("no empieza en in")
	}
}
