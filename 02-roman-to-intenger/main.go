package main

import "fmt"

func romanToIntenger(roman string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	number := 0
	n := len(roman)

	for i := 0; i < n; i++ {
		// Si el siguiente valor es mayor, resta el valor actual
		if i < n-1 && romanMap[roman[i]] < romanMap[roman[i+1]] {
			number -= romanMap[roman[i]]
		} else {
			// De lo contrario, suma el valor actual
			number += romanMap[roman[i]]
		}
	}

	return number
}

func main() {
	fmt.Println(romanToIntenger("IV")) // Output: 3
	// Explicacion: III = 3
	fmt.Println(romanToIntenger("LVIII")) // Output: 58
	// Explicacion: L = 50, V= 5, III = 3
	fmt.Println(romanToIntenger("MCMXCIV")) // Output: 1994
	// Explicacion: M = 1000, CM = 900, XC = 90, IV = 4
}
