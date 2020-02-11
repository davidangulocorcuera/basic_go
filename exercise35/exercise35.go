package main

//      * 35 Haz un programa que imprima por pantalla los 100 primeros números impares
func main() {
	var counter = 0
	var currentNumber = 1
	for counter < 100{
		if currentNumber % 2 != 0 {
			println(currentNumber)
			counter++
		}
		currentNumber++
	}
}