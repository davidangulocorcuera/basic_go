package main

import "fmt"

// * 37 Haz un programa que ordene de mayor a menor 5 números introducidos por teclado

func main() {
	var numbers = [5]int{}
	var currentNumber int

	for i := 0; i < len(numbers); i++ {
		fmt.Print("Introduce un número ")
		fmt.Scanf("%d", &numbers[i])
	}

	for i := 0; i < len(numbers); i++ {
		currentNumber = numbers[i]
		for j := 0; j < len(numbers); j++ {
			if currentNumber > numbers[j] {
				var temp = numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
			}
		}

	}

	for i := 0; i < len(numbers); i++ {
		print(numbers[i], ",")
	}
}
