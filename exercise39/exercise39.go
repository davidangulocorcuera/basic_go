package main

//* 39 Haz un programa que ordene caracteres del abecedario de mayor a menos, entendiendo que la a es 1 y la z 27.

import "strconv"

func main() {
	var letters = loadMap()
	var inputCharacters = [5]rune{'b', 'x', 'd', 'a','u'}

	for i := 0; i < len(inputCharacters); i++ {
		currentCharacter := letters[inputCharacters[i]]
		for j := 0; j < len(inputCharacters); j++ {
			otherValue := letters[inputCharacters[j]]
			if currentCharacter < otherValue {
				var temp = inputCharacters[i]
				inputCharacters[i] = inputCharacters[j]
				inputCharacters[j] = temp
			}
		}

	}

	for i := 0; i < len(inputCharacters); i++ {
		print(strconv.QuoteRune(inputCharacters[i])  , ",")
	}
}

func loadMap() map[rune]int {
	var letters = make(map[rune]int)
	letters['a'] = 1
	letters['b'] = 2
	letters['c'] = 3
	letters['d'] = 4
	letters['e'] = 5
	letters['f'] = 6
	letters['g'] = 7
	letters['h'] = 8
	letters['i'] = 9
	letters['j'] = 10
	letters['k'] = 11
	letters['l'] = 12
	letters['m'] = 13
	letters['n'] = 14
	letters['ñ'] = 15
	letters['o'] = 16
	letters['p'] = 17
	letters['q'] = 18
	letters['r'] = 19
	letters['s'] = 20
	letters['t'] = 21
	letters['u'] = 22
	letters['v'] = 23
	letters['w'] = 24
	letters['x'] = 25
	letters['y'] = 26
	letters['z'] = 27
	return letters
}
