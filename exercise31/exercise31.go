package main

import "fmt"

//      * 32 Haz un programa que compruebe que un texto introducido por teclado termina en "on"

func main() {
	var text string
	fmt.Print("Introduce un texto ")
	fmt.Scanf("%s", &text)

	var last = text[len(text)-2:]

	if last == "on" {
		fmt.Print("si termina en on")
	} else {
		fmt.Print("no termina en on")
	}
}
