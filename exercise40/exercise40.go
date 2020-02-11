package main

//  * 40 Haz un programa que imprima por pantalla los 1000 primeros números primos

func main() {
	var counter = 0
	var currentNumber = 1
	var divisores = 0

	for counter < 1000 {
		for i := 1; i <= currentNumber; i++ {
			if currentNumber%i == 0 {
				divisores++
			}
		}
		if divisores == 2 {
			counter++
			println(currentNumber)
		}
		divisores = 0
		currentNumber++
	}

}
