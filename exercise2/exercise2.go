package main

//2 Haz un programa que imprima un número introducido por teclado

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Print("Introduce un número ")
	var number, _ = reader.ReadString('\n')
	fmt.Print(number)
}

