package main

import "fmt"

//      * 28 Haz un programa que tenga dos opciones,
//     en la primera realice a + b y en la segunda a - b, el suario debe elegir la opcion que quiera

func main() {

	var option int
	var a int
	var b int

	fmt.Print("Introduce el valor de a ")
	fmt.Scanf("%d", &a)

	fmt.Print("Introduce el valor de b ")
	fmt.Scanf("%d", &b)

	fmt.Print("Introduce una opción 1 para sumar o 2 para restar ")
	fmt.Scanf("%d", &option)

	switch option {
	case 1:
		{
			fmt.Print(a + b)
		}
	case 2:
		{
			fmt.Print(a - b)
		}
	default:
		fmt.Print("No has introducido una opción correcta")
	}
}
