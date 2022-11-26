package variables

import "fmt"

func RunStrings() {
	var source = "Some days are ☀️, and other days are 🌧."
	var sunny = source[0:20]
	var rainy = source[26 : len(source)-1]

	fmt.Println(len(sunny))
	fmt.Println(sunny)
	fmt.Println(len(rainy))
	fmt.Println(rainy)
}
