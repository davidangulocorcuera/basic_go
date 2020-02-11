package main

//      * 36 Haz un programa que imprima por pantalla los 15 primeros números de la serie de fibonacci.
// 1,1,2,3,5,8,13,21,34 . . .

func main() {
	var counter = 1
	var num1 = 1
	var num2 = 1
	var currentValue = 1

	print(num1, ",")
	for counter <= 15 {
		print(currentValue, ",")
		currentValue = num1 + num2
		num1 = num2
		num2 = currentValue
		counter++
	}

}
