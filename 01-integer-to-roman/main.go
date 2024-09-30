package main

import "fmt"

func main() {
	fmt.Println(intengerToRoman(3749))
	fmt.Println(intengerToRoman(58))
	fmt.Println(intengerToRoman(1994))
}

func intengerToRoman(num int) string {
	var result string
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := range values {
		for num >= values[i] {
			num -= values[i]
			result += romans[i]
		}
	}

	return result
}
